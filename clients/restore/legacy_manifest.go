package restore

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
)

type versionSwitch struct {
	Version int `json:"manifestVersion,omitempty"`
}

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
