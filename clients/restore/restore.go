package restore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"github.com/influxdata/influx-cli/v2/pkg/gzip"
)

type ApiConfig interface {
	GetConfig() *api.Configuration
}

type Client struct {
	clients.CLI
	api.HealthApi
	api.RestoreApi
	api.BucketsApi
	api.OrganizationsApi
	ApiConfig

	manifest br.Manifest
}

type Params struct {
	// Path to local backup data created using `influx backup`
	Path string

	// Original ID/name of the organization to restore.
	// If not set, all orgs will be restored.
	OrgID string
	Org   string

	// New name to use for the restored organization.
	// If not set, the org will be restored using its backed-up name.
	NewOrgName string

	// Original ID/name of the bucket to restore.
	// If not set, all buckets within the org filter will be restored.
	BucketID string
	Bucket   string

	// New name to use for the restored bucket.
	// If not set, the bucket will be restored using its backed-up name.
	NewBucketName string

	// If true, replace all data on the server with the local backup.
	// Otherwise only restore the requested org/bucket, leaving other data untouched.
	Full bool
}

func (p *Params) matches(bkt br.ManifestBucketEntry) bool {
	if p.OrgID != "" && bkt.OrganizationID != p.OrgID {
		return false
	}
	if p.Org != "" && bkt.OrganizationName != p.Org {
		return false
	}
	if p.BucketID != "" && bkt.BucketID != p.BucketID {
		return false
	}
	if p.Bucket != "" && bkt.BucketName != p.Bucket {
		return false
	}
	return true
}

func (c *Client) Restore(ctx context.Context, params *Params) error {
	if err := c.loadManifests(params.Path); err != nil {
		return err
	}

	// The APIs we use to restore data depends on the server's version.
	legacyServer, err := br.ServerIsLegacy(ctx, c.HealthApi)
	if err != nil {
		return err
	}

	if params.Full {
		return c.fullRestore(ctx, params.Path, legacyServer)
	}
	return c.partialRestore(ctx, params, legacyServer)
}

// loadManifests finds and merges all backup manifests stored in a given directory,
// keeping the latest top-level metadata and latest metadata per-bucket.
func (c *Client) loadManifests(path string) error {
	// Read all manifest files from path, sort in ascending time.
	manifests, err := filepath.Glob(filepath.Join(path, fmt.Sprintf("*.%s", br.ManifestExtension)))
	if err != nil {
		return fmt.Errorf("failed to find backup manifests at %q: %w", path, err)
	} else if len(manifests) == 0 {
		return fmt.Errorf("no backup manifests found at %q", path)
	}
	sort.Strings(manifests)

	bucketManifests := map[string]br.ManifestBucketEntry{}
	for _, manifestFile := range manifests {
		// Skip file if it is a directory.
		if fi, err := os.Stat(manifestFile); err != nil {
			return fmt.Errorf("failed to inspect local manifest at %q: %w", manifestFile, err)
		} else if fi.IsDir() {
			continue
		}

		manifest, err := readManifest(manifestFile)
		if err != nil {
			return err
		}

		// Keep the latest KV and SQL overall.
		c.manifest.KV = manifest.KV
		c.manifest.SQL = manifest.SQL

		// Keep the latest manifest per-bucket.
		for _, bkt := range manifest.Buckets {
			// NOTE: Deduplicate here by keeping only the latest entry for each `<org-name>/<bucket-name>` pair.
			// This prevents "bucket already exists" errors during the restore when the backup manifests contain
			// entries for multiple buckets with the same name (which can happen when a bucket is deleted & re-created).
			bucketManifests[fmt.Sprintf("%s/%s", bkt.OrganizationName, bkt.BucketName)] = bkt
		}
	}

	c.manifest.Buckets = make([]br.ManifestBucketEntry, 0, len(bucketManifests))
	for _, bkt := range bucketManifests {
		c.manifest.Buckets = append(c.manifest.Buckets, bkt)
	}

	return nil
}

// fullRestore completely replaces all metadata and data on the server with the contents of a local backup.
func (c Client) fullRestore(ctx context.Context, path string, legacy bool) error {
	if legacy && c.manifest.SQL != nil {
		return fmt.Errorf("cannot fully restore data from %s: target server's version too old to restore SQL metadata", path)
	}

	// Make sure we can read both local metadata snapshots before
	readKv := readFileGzipped
	if legacy {
		// The legacy API didn't support gzipped uploads.
		readKv = readFileGunzipped
	}
	kvBytes, err := readKv(path, c.manifest.KV)
	if err != nil {
		return fmt.Errorf("failed to open local KV backup at %q: %w", filepath.Join(path, c.manifest.KV.FileName), err)
	}
	defer kvBytes.Close()

	var sqlBytes io.ReadCloser
	if c.manifest.SQL != nil {
		sqlBytes, err = readFileGzipped(path, *c.manifest.SQL)
		if err != nil {
			return fmt.Errorf("failed to open local SQL backup at %q: %w", filepath.Join(path, c.manifest.SQL.FileName), err)
		}
		defer sqlBytes.Close()
	}

	// Upload metadata snapshots to the server.
	log.Println("INFO: Restoring KV snapshot")
	kvReq := c.PostRestoreKV(ctx).ContentType("application/octet-stream").Body(kvBytes)
	if !legacy {
		kvReq = kvReq.ContentEncoding("gzip")
	}

	// Deal with new token
	newOperatorToken, err := kvReq.Execute()
	if err != nil {
		return fmt.Errorf("failed to restore KV snapshot: %w", err)
	}
	if newOperatorToken.Token != nil {
		newAuthorization := fmt.Sprintf("Token %s", *newOperatorToken.Token)
		const authorizationHeader = "Authorization"
		if newAuthorization != c.ApiConfig.GetConfig().DefaultHeader[authorizationHeader] {
			log.Println("WARN: Restoring KV snapshot overwrote the operator token, ensure following commands use the correct token")
			c.ApiConfig.GetConfig().DefaultHeader[authorizationHeader] = newAuthorization
		}
	}

	// TODO: Should we have some way of wiping out any existing SQL on the server-side in the case when there is no backup?
	if c.manifest.SQL != nil {
		// NOTE: No logic here to upload non-gzipped data because legacy=true doesn't support SQL restores.
		log.Println("INFO: Restoring SQL snapshot")
		if err := c.PostRestoreSQL(ctx).
			ContentEncoding("gzip").
			ContentType("application/octet-stream").
			Body(sqlBytes).
			Execute(); err != nil {
			return fmt.Errorf("failed to restore SQL snapshot: %w", err)
		}
	}

	// Drill down through bucket manifests to reach shard info, and upload it.
	for _, b := range c.manifest.Buckets {
		for _, rp := range b.RetentionPolicies {
			for _, sg := range rp.ShardGroups {
				for _, s := range sg.Shards {
					if err := c.restoreShard(ctx, path, s, legacy); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

// partialRestore creates a bucket (or buckets) on the target server, and seeds it with data
// from a local backup.
func (c Client) partialRestore(ctx context.Context, params *Params, legacy bool) (err error) {
	orgIds := map[string]string{}

	for _, bkt := range c.manifest.Buckets {
		// Skip internal buckets.
		if strings.HasPrefix(bkt.BucketName, "_") {
			continue
		}
		if !params.matches(bkt) {
			continue
		}

		orgName := bkt.OrganizationName
		// Before this method is called, we ensure that new-org-name is only set if
		// a filter on org-name or org-id is set. If that check passes and execution
		// reaches this code, we can assume that all buckets matching the filter come
		// from the same org, so we can swap in the new org name unconditionally.
		if params.NewOrgName != "" {
			orgName = params.NewOrgName
		}

		if _, ok := orgIds[orgName]; !ok {
			orgIds[orgName], err = c.restoreOrg(ctx, orgName)
			if err != nil {
				return
			}
		}
		bkt.OrganizationName = orgName
		bkt.OrganizationID = orgIds[orgName]

		// By the same reasoning as above, if new-bucket-name is non-empty we know
		// filters must have been set to ensure we only match 1 bucket, so we can
		// swap the name without additional checks.
		if params.NewBucketName != "" {
			bkt.BucketName = params.NewBucketName
		}

		restoreBucket := c.restoreBucket
		if legacy {
			restoreBucket = c.restoreBucketLegacy
		}
		shardIdMap, err := restoreBucket(ctx, bkt)
		if err != nil {
			return fmt.Errorf("failed to restore bucket %q: %w", bkt.BucketName, err)
		}

		for _, rp := range bkt.RetentionPolicies {
			for _, sg := range rp.ShardGroups {
				for _, sh := range sg.Shards {
					newID, ok := shardIdMap[sh.ID]
					if !ok {
						log.Printf("WARN: Server didn't map ID for shard %d in bucket %q, skipping\n", sh.ID, bkt.BucketName)
						continue
					}
					sh.ID = newID
					if err := c.restoreShard(ctx, params.Path, sh, legacy); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

// restoreBucket creates a new bucket and pre-generates a set of shards within that bucket, returning
// a mapping between the shard IDs stored in a local backup and the new shard IDs generated on the server.
func (c Client) restoreBucket(ctx context.Context, bkt br.ManifestBucketEntry) (map[int64]int64, error) {
	log.Printf("INFO: Restoring bucket %q as %q\n", bkt.BucketID, bkt.BucketName)
	bucketMapping, err := c.PostRestoreBucketMetadata(ctx).
		BucketMetadataManifest(ConvertBucketManifest(bkt)).
		Execute()
	if err != nil {
		return nil, err
	}

	shardIdMap := make(map[int64]int64, len(bucketMapping.ShardMappings))
	for _, mapping := range bucketMapping.ShardMappings {
		shardIdMap[mapping.OldId] = mapping.NewId
	}
	return shardIdMap, nil
}

// restoreBucketLegacy creates a new bucket and pre-generates a set of shards within that bucket, returning
// a mapping between the shard IDs stored in a local backup and the new shard IDs generated on the server.
//
// The server-side logic to do all this was introduced in v2.1.0. To support using newer CLI versions against
// v2.0.x of the server, we replicate the logic here via multiple API calls.
func (c Client) restoreBucketLegacy(ctx context.Context, bkt br.ManifestBucketEntry) (map[int64]int64, error) {
	log.Printf("INFO: Restoring bucket %q as %q using legacy APIs\n", bkt.BucketID, bkt.BucketName)
	// Legacy APIs require creating the bucket as a separate call.
	rps := make([]api.RetentionRule, len(bkt.RetentionPolicies))
	for i, rp := range bkt.RetentionPolicies {
		rps[i] = *api.NewRetentionRuleWithDefaults()
		rps[i].EverySeconds = int64(time.Duration(rp.Duration).Seconds())
		sgd := int64(time.Duration(rp.ShardGroupDuration).Seconds())
		rps[i].ShardGroupDurationSeconds = &sgd
	}

	bucketReq := *api.NewPostBucketRequest(bkt.OrganizationID, bkt.BucketName, rps)
	bucketReq.Description = bkt.Description

	newBkt, err := c.PostBuckets(ctx).PostBucketRequest(bucketReq).Execute()
	if err != nil {
		return nil, fmt.Errorf("couldn't create bucket %q: %w", bkt.BucketName, err)
	}

	dbi := bucketToDBI(bkt)
	dbiBytes, err := proto.Marshal(&dbi)
	if err != nil {
		return nil, err
	}

	shardMapJSON, err := c.PostRestoreBucketID(ctx, *newBkt.Id).Body(dbiBytes).Execute()
	if err != nil {
		return nil, fmt.Errorf("couldn't restore database info for %q: %w", bkt.BucketName, err)
	}

	var shardMap map[int64]int64
	if err := json.Unmarshal([]byte(shardMapJSON), &shardMap); err != nil {
		return nil, fmt.Errorf("couldn't parse result of restoring database info for %q: %w", bkt.BucketName, err)
	}

	return shardMap, nil
}

// restoreOrg gets the ID for the org with the given name, creating the org if it doesn't already exist.
func (c Client) restoreOrg(ctx context.Context, name string) (string, error) {
	// NOTE: Our orgs API returns a 404 instead of an empty list when filtering by a specific name.
	orgs, err := c.GetOrgs(ctx).Org(name).Execute()
	if err != nil {
		if apiErr, ok := err.(api.ApiError); !ok || apiErr.ErrorCode() != api.ERRORCODE_NOT_FOUND {
			return "", fmt.Errorf("failed to check existence of organization %q: %w", name, err)
		}
	}

	// If we've gotten this far and err != nil, it means err was a 404.
	if err != nil || len(orgs.GetOrgs()) == 0 {
		// Create any missing orgs.
		newOrg, err := c.PostOrgs(ctx).PostOrganizationRequest(api.PostOrganizationRequest{Name: name}).Execute()
		if err != nil {
			return "", fmt.Errorf("failed to create organization %q: %w", name, err)
		}
		return *newOrg.Id, nil
	}

	return *orgs.GetOrgs()[0].Id, nil
}

func bucketToDBI(b br.ManifestBucketEntry) br.DatabaseInfo {
	dbi := br.DatabaseInfo{
		Name:                   &b.BucketID,
		DefaultRetentionPolicy: &b.DefaultRetentionPolicy,
		RetentionPolicies:      make([]*br.RetentionPolicyInfo, len(b.RetentionPolicies)),
		ContinuousQueries:      nil,
	}
	for i, rp := range b.RetentionPolicies {
		converted := retentionPolicyToRPI(rp)
		dbi.RetentionPolicies[i] = &converted
	}
	return dbi
}

func retentionPolicyToRPI(rp br.ManifestRetentionPolicy) br.RetentionPolicyInfo {
	replicaN := uint32(rp.ReplicaN)
	rpi := br.RetentionPolicyInfo{
		Name:               &rp.Name,
		Duration:           &rp.Duration,
		ShardGroupDuration: &rp.ShardGroupDuration,
		ReplicaN:           &replicaN,
		ShardGroups:        make([]*br.ShardGroupInfo, len(rp.ShardGroups)),
		Subscriptions:      make([]*br.SubscriptionInfo, len(rp.Subscriptions)),
	}
	for i, sg := range rp.ShardGroups {
		converted := shardGroupToSGI(sg)
		rpi.ShardGroups[i] = &converted
	}
	for i, s := range rp.Subscriptions {
		converted := br.SubscriptionInfo{
			Name:         &s.Name,
			Mode:         &s.Mode,
			Destinations: s.Destinations,
		}
		rpi.Subscriptions[i] = &converted
	}
	return rpi
}

func shardGroupToSGI(sg br.ManifestShardGroup) br.ShardGroupInfo {
	id := uint64(sg.ID)
	start := sg.StartTime.UnixNano()
	end := sg.EndTime.UnixNano()
	var deleted, truncated int64
	if sg.DeletedAt != nil {
		deleted = sg.DeletedAt.UnixNano()
	}
	if sg.TruncatedAt != nil {
		truncated = sg.TruncatedAt.UnixNano()
	}
	sgi := br.ShardGroupInfo{
		ID:          &id,
		StartTime:   &start,
		EndTime:     &end,
		DeletedAt:   &deleted,
		Shards:      make([]*br.ShardInfo, len(sg.Shards)),
		TruncatedAt: &truncated,
	}
	for i, s := range sg.Shards {
		converted := shardToSI(s)
		sgi.Shards[i] = &converted
	}
	return sgi
}

func shardToSI(shard br.ManifestShardEntry) br.ShardInfo {
	id := uint64(shard.ID)
	si := br.ShardInfo{
		ID:     &id,
		Owners: make([]*br.ShardOwner, len(shard.ShardOwners)),
	}
	for i, o := range shard.ShardOwners {
		oid := uint64(o.NodeID)
		converted := br.ShardOwner{NodeID: &oid}
		si.Owners[i] = &converted
	}
	return si
}

// readFileGzipped opens a local file and returns a reader of its contents,
// compressed with gzip.
func readFileGzipped(path string, file br.ManifestFileEntry) (io.ReadCloser, error) {
	fullPath := filepath.Join(path, file.FileName)
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	if file.Compression == br.GzipCompression {
		return f, nil
	}
	return gzip.NewGzipPipe(f), nil
}

// readFileGunzipped opens a local file and returns a reader of its contents,
// gunzipping it if it is compressed.
func readFileGunzipped(path string, file br.ManifestFileEntry) (io.ReadCloser, error) {
	fullPath := filepath.Join(path, file.FileName)
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	if file.Compression == br.NoCompression {
		return f, nil
	}
	reader, err := gzip.NewGunzipReadCloser(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return reader, nil
}

// restoreShard overwrites the contents of a single shard on the server using TSM stored in a local backup.
func (c Client) restoreShard(ctx context.Context, path string, m br.ManifestShardEntry, legacy bool) error {
	read := readFileGzipped
	if legacy {
		// The legacy API didn't support gzipped uploads.
		read = readFileGunzipped
	}
	// Make sure we can read the local snapshot.
	tsmBytes, err := read(path, m.ManifestFileEntry)
	if err != nil {
		return fmt.Errorf("failed to open local TSM snapshot at %q: %w", filepath.Join(path, m.FileName), err)
	}
	defer tsmBytes.Close()

	req := c.PostRestoreShardId(ctx, fmt.Sprintf("%d", m.ID)).
		ContentType("application/octet-stream").
		Body(tsmBytes)
	if !legacy {
		req = req.ContentEncoding("gzip")
	}
	log.Printf("INFO: Restoring TSM snapshot for shard %d\n", m.ID)
	if err := req.Execute(); err != nil {
		return fmt.Errorf("failed to restore TSM snapshot for shard %d: %w", m.ID, err)
	}
	return nil
}
