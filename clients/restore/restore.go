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

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"github.com/influxdata/influx-cli/v2/pkg/gzip"
)

type Client struct {
	clients.CLI
	api.RestoreApi

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

func (c *Client) Restore(ctx context.Context, params *Params) error {
	if err := c.loadManifests(params.Path); err != nil {
		return err
	}
	if params.Full {
		return c.fullRestore(ctx, params.Path)
	}
	return c.partialRestore(ctx, params)
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
	sort.Sort(sort.StringSlice(manifests))

	bucketManifests := map[string]br.ManifestBucketEntry{}
	for _, manifestFile := range manifests {
		// Skip file if it is a directory.
		if fi, err := os.Stat(manifestFile); err != nil {
			return fmt.Errorf("failed to inspect local manifest at %q: %w", manifestFile, err)
		} else if fi.IsDir() {
			continue
		}

		var manifest br.Manifest
		if buf, err := os.ReadFile(manifestFile); err != nil {
			return fmt.Errorf("failed to read local manifest at %q: %w", manifestFile, err)
		} else if err := json.Unmarshal(buf, &manifest); err != nil {
			return fmt.Errorf("failed to parse manifest at %q: %w", manifestFile, err)
		}

		// Keep the latest KV and SQL overall.
		c.manifest.KV = manifest.KV
		c.manifest.SQL = manifest.SQL

		// Keep the latest manifest per-bucket.
		for _, bkt := range manifest.Buckets {
			bucketManifests[bkt.BucketID] = bkt
		}
	}

	c.manifest.Buckets = make([]br.ManifestBucketEntry, 0, len(bucketManifests))
	for _, bkt := range bucketManifests {
		c.manifest.Buckets = append(c.manifest.Buckets, bkt)
	}

	return nil
}

func (c Client) fullRestore(ctx context.Context, path string) error {
	// Make sure we can read both local metadata snapshots before
	kvBytes, err := c.readFileGzipped(path, c.manifest.KV)
	if err != nil {
		return fmt.Errorf("failed to open local KV backup at %q: %w", filepath.Join(path, c.manifest.KV.FileName), err)
	}
	defer kvBytes.Close()
	sqlBytes, err := c.readFileGzipped(path, c.manifest.SQL)
	if err != nil {
		return fmt.Errorf("failed to open local SQL backup at %q: %w", filepath.Join(path, c.manifest.SQL.FileName), err)
	}
	defer sqlBytes.Close()

	// Upload metadata snapshots to the server.
	log.Println("INFO: Restoring KV snapshot")
	if err := c.PostRestoreKV(ctx).ContentEncoding("gzip").Body(kvBytes).Execute(); err != nil {
		return fmt.Errorf("failed to restore KV snapshot: %w", err)
	}
	log.Println("INFO: Restoring SQL snapshot")
	if err := c.PostRestoreSQL(ctx).ContentEncoding("gzip").Body(sqlBytes).Execute(); err != nil {
		return fmt.Errorf("failed to restore SQL snapshot: %w", err)
	}

	// Drill down through bucket manifests to reach shard info, and upload it.
	for _, b := range c.manifest.Buckets {
		for _, rp := range b.RetentionPolicies {
			for _, sg := range rp.ShardGroups {
				for _, s := range sg.Shards {
					if err := c.restoreShard(ctx, path, s); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (c Client) partialRestore(ctx context.Context, params *Params) error {
	panic("TODO")
}

// readFileGzipped opens a local file and returns a reader of its contents,
// compressed with gzip.
func (c Client) readFileGzipped(path string, file br.ManifestFileEntry) (io.ReadCloser, error) {
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

func (c Client) restoreShard(ctx context.Context, path string, m br.ManifestShardEntry) error {
	// Make sure we can read the local snapshot.
	tsmBytes, err := c.readFileGzipped(path, m.ManifestFileEntry)
	if err != nil {
		return fmt.Errorf("failed to open local TSM snapshot at %q: %w", filepath.Join(path, m.FileName), err)
	}
	defer tsmBytes.Close()

	log.Printf("INFO: Restoring TSM snapshot for shard %d\n", m.ID)
	if err := c.PostRestoreShardId(ctx, fmt.Sprintf("%d", m.ID)).ContentEncoding("gzip").Body(tsmBytes).Execute(); err != nil {
		return fmt.Errorf("failed to restore TSM snapshot for shard %d: %w", m.ID, err)
	}
	return nil
}
