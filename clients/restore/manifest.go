package restore

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/influxdata/influx-cli/v2/api"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
)

// versionSwitch models the subset of fields needed to distinguish different versions of the CLI's backup manifest.
type versionSwitch struct {
	Version int `json:"manifestVersion,omitempty"`
}

// readManifest parses the manifest file at the given path, converting it to the latest version of our manifest
// if needed.
func readManifest(path string) (manifest br.Manifest, err error) {
	var w struct {
		versionSwitch
		*br.Manifest
		*legacyManifest
	}
	buf, err := os.ReadFile(path)
	if err != nil {
		return br.Manifest{}, fmt.Errorf("failed to read local manifest at %q: %w", path, err)
	}

	if err := json.Unmarshal(buf, &w.versionSwitch); err != nil {
		return br.Manifest{}, fmt.Errorf("failed to check version of local manifest at %q: %w", path, err)
	}
	switch w.versionSwitch.Version {
	case br.ManifestVersion:
		err = json.Unmarshal(buf, &manifest)
	case 0: // InfluxDB 2.0.x manifests didn't have a version field.
		var lm legacyManifest
		if err := json.Unmarshal(buf, &lm); err != nil {
			return br.Manifest{}, fmt.Errorf("failed to parse legacy manifest at %q: %w", path, err)
		}
		manifest, err = convertManifest(path, lm)
	default:
		return br.Manifest{}, fmt.Errorf("unsupported version %d found in manifest at %q", w.versionSwitch.Version, path)
	}

	return
}

// convertManifest converts a manifest from the 2.0.x CLI into the latest manifest schema.
// NOTE: 2.0.x manifests didn't contain all the info needed by 2.1.x+, so this process requires opening & inspecting
// the bolt file referenced by the legacy manifest.
func convertManifest(path string, lm legacyManifest) (br.Manifest, error) {
	// Extract bucket metadata from the local KV snapshot.
	boltPath := filepath.Join(filepath.Dir(path), lm.KV.FileName)
	metadata, err := br.ExtractBucketMetadata(boltPath)
	if err != nil {
		return br.Manifest{}, err
	}
	shardManifestsById := make(map[int64]br.ManifestFileEntry, len(lm.Shards))
	for _, s := range lm.Shards {
		shardManifestsById[s.ShardID] = br.ManifestFileEntry{
			FileName:    s.FileName,
			Size:        s.Size,
			Compression: br.GzipCompression,
		}
	}

	m := br.Manifest{
		Version: br.ManifestVersion,
		KV: br.ManifestFileEntry{
			FileName:    lm.KV.FileName,
			Size:        lm.KV.Size,
			Compression: br.NoCompression,
		},
		Buckets: make([]br.ManifestBucketEntry, len(metadata)),
	}
	for i, bkt := range metadata {
		m.Buckets[i], err = br.ConvertBucketManifest(bkt, func(shardId int64) (*br.ManifestFileEntry, error) {
			shardManifest, ok := shardManifestsById[shardId]
			if !ok {
				return nil, nil
			}
			return &shardManifest, nil
		})
		if err != nil {
			return br.Manifest{}, fmt.Errorf("failed to parse entry for bucket %q in legacy manifest at %q: %w", bkt.BucketID, path, err)
		}
	}

	return m, nil
}

// legacyManifest models the subset of data stored in 2.0.x CLI backup manifests that is needed for conversion
// into the latest manifest format.
type legacyManifest struct {
	KV     legacyKV      `json:"kv"`
	Shards []legacyShard `json:"files"`
}

type legacyKV struct {
	FileName string `json:"fileName"`
	Size     int64  `json:"size"`
}

type legacyShard struct {
	ShardID  int64  `json:"shardID"`
	FileName string `json:"fileName"`
	Size     int64  `json:"size"`
}

// ConvertBucketManifest converts a manifest parsed from local disk into a model compatible with the server-side API.
func ConvertBucketManifest(manifest br.ManifestBucketEntry) api.BucketMetadataManifest {
	m := api.BucketMetadataManifest{
		OrganizationID:         manifest.OrganizationID,
		OrganizationName:       manifest.OrganizationName,
		BucketID:               manifest.BucketID,
		BucketName:             manifest.BucketName,
		Description:            manifest.Description,
		DefaultRetentionPolicy: manifest.DefaultRetentionPolicy,
		RetentionPolicies:      make([]api.RetentionPolicyManifest, len(manifest.RetentionPolicies)),
	}

	for i, rp := range manifest.RetentionPolicies {
		m.RetentionPolicies[i] = ConvertRetentionPolicy(rp)
	}

	return m
}

func ConvertRetentionPolicy(manifest br.ManifestRetentionPolicy) api.RetentionPolicyManifest {
	m := api.RetentionPolicyManifest{
		Name:               manifest.Name,
		ReplicaN:           manifest.ReplicaN,
		Duration:           manifest.Duration,
		ShardGroupDuration: manifest.ShardGroupDuration,
		ShardGroups:        make([]api.ShardGroupManifest, len(manifest.ShardGroups)),
		Subscriptions:      make([]api.SubscriptionManifest, len(manifest.Subscriptions)),
	}

	for i, sg := range manifest.ShardGroups {
		m.ShardGroups[i] = ConvertShardGroup(sg)
	}

	for i, s := range manifest.Subscriptions {
		m.Subscriptions[i] = api.SubscriptionManifest{
			Name:         s.Name,
			Mode:         s.Mode,
			Destinations: s.Destinations,
		}
	}

	return m
}

func ConvertShardGroup(manifest br.ManifestShardGroup) api.ShardGroupManifest {
	m := api.ShardGroupManifest{
		Id:          manifest.ID,
		StartTime:   manifest.StartTime,
		EndTime:     manifest.EndTime,
		DeletedAt:   manifest.DeletedAt,
		TruncatedAt: manifest.TruncatedAt,
		Shards:      make([]api.ShardManifest, len(manifest.Shards)),
	}

	for i, sh := range manifest.Shards {
		m.Shards[i] = ConvertShard(sh)
	}

	return m
}

func ConvertShard(manifest br.ManifestShardEntry) api.ShardManifest {
	m := api.ShardManifest{
		Id:          manifest.ID,
		ShardOwners: make([]api.ShardOwner, len(manifest.ShardOwners)),
	}

	for i, so := range manifest.ShardOwners {
		m.ShardOwners[i] = api.ShardOwner{NodeID: so.NodeID}
	}

	return m
}
