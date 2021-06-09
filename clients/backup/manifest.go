package backup

import (
	"fmt"
	"os"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
)

type FileCompression int

const (
	NoCompression FileCompression = iota
	GzipCompression
)

func (c *FileCompression) Set(v string) error {
	switch v {
	case "none":
		*c = NoCompression
	case "gzip":
		*c = GzipCompression
	default:
		return fmt.Errorf("unsupported format: %q", v)
	}
	return nil
}

func (c FileCompression) String() string {
	switch c {
	case NoCompression:
		return "none"
	case GzipCompression:
		return "gzip"
	default:
		panic("Impossible!")
	}
}

type Manifest struct {
	KV      ManifestFileEntry     `json:"kv"`
	SQL     ManifestFileEntry     `json:"sql"`
	Buckets []ManifestBucketEntry `json:"buckets"`
}

type ManifestFileEntry struct {
	FileName    string          `json:"fileName"`
	Size        int64           `json:"size"`
	Compression FileCompression `json:"compression"`
}

func ConvertBucketManifest(manifest api.BucketMetadataManifest, getShard func(shardId int64) (os.FileInfo, error)) (ManifestBucketEntry, error) {
	m := ManifestBucketEntry{
		OrganizationID:         manifest.OrganizationID,
		OrganizationName:       manifest.OrganizationName,
		BucketID:               manifest.BucketID,
		BucketName:             manifest.BucketName,
		DefaultRetentionPolicy: manifest.DefaultRetentionPolicy,
		RetentionPolicies:      make([]ManifestRetentionPolicy, len(manifest.RetentionPolicies)),
	}

	for i, rp := range manifest.RetentionPolicies {
		var err error
		m.RetentionPolicies[i], err = ConvertRetentionPolicy(rp, getShard)
		if err != nil {
			return ManifestBucketEntry{}, err
		}
	}

	return m, nil
}

func ConvertRetentionPolicy(manifest api.RetentionPolicyManifest, getShard func(shardId int64) (os.FileInfo, error)) (ManifestRetentionPolicy, error) {
	m := ManifestRetentionPolicy{
		Name:               manifest.Name,
		ReplicaN:           manifest.ReplicaN,
		Duration:           manifest.Duration,
		ShardGroupDuration: manifest.ShardGroupDuration,
		ShardGroups:        make([]ManifestShardGroup, len(manifest.ShardGroups)),
		Subscriptions:      make([]ManifestSubscription, len(manifest.Subscriptions)),
	}

	for i, sg := range manifest.ShardGroups {
		var err error
		m.ShardGroups[i], err = ConvertShardGroup(sg, getShard)
		if err != nil {
			return ManifestRetentionPolicy{}, err
		}
	}

	for i, s := range manifest.Subscriptions {
		m.Subscriptions[i] = ManifestSubscription{
			Name:         s.Name,
			Mode:         s.Mode,
			Destinations: s.Destinations,
		}
	}

	return m, nil
}

func ConvertShardGroup(manifest api.ShardGroupManifest, getShard func(shardId int64) (os.FileInfo, error)) (ManifestShardGroup, error) {
	m := ManifestShardGroup{
		ID:          manifest.Id,
		StartTime:   manifest.StartTime,
		EndTime:     manifest.EndTime,
		DeletedAt:   manifest.DeletedAt,
		TruncatedAt: manifest.TruncatedAt,
		Shards:      make([]ManifestShardEntry, len(manifest.Shards)),
	}

	for i, sh := range manifest.Shards {
		var err error
		m.Shards[i], err = ConvertShard(sh, getShard)
		if err != nil {
			return ManifestShardGroup{}, err
		}
	}

	return m, nil
}

func ConvertShard(manifest api.ShardManifest, getShard func(shardId int64) (os.FileInfo, error)) (ManifestShardEntry, error) {
	shardFileInfo, err := getShard(manifest.Id)
	if err != nil {
		return ManifestShardEntry{}, fmt.Errorf("failed to download snapshot of shard %d: %w", manifest.Id, err)
	}

	m := ManifestShardEntry{
		ID:          manifest.Id,
		ShardOwners: make([]ShardOwner, len(manifest.ShardOwners)),
		ManifestFileEntry: ManifestFileEntry{
			FileName: shardFileInfo.Name(),
			Size:     shardFileInfo.Size(),
		},
	}

	for i, o := range manifest.ShardOwners {
		m.ShardOwners[i] = ShardOwner{
			NodeID: o.NodeID,
		}
	}

	return m, nil
}

type ManifestBucketEntry struct {
	OrganizationID         string                    `json:"organizationID"`
	OrganizationName       string                    `json:"organizationName"`
	BucketID               string                    `json:"bucketID"`
	BucketName             string                    `json:"bucketName"`
	DefaultRetentionPolicy string                    `json:"defaultRetentionPolicy"`
	RetentionPolicies      []ManifestRetentionPolicy `json:"retentionPolicies"`
}

type ManifestRetentionPolicy struct {
	Name               string                 `json:"name"`
	ReplicaN           int32                  `json:"replicaN"`
	Duration           int64                  `json:"duration"`
	ShardGroupDuration int64                  `json:"shardGroupDuration"`
	ShardGroups        []ManifestShardGroup   `json:"shardGroups"`
	Subscriptions      []ManifestSubscription `json:"subscriptions"`
}

type ManifestShardGroup struct {
	ID          int64                `json:"id"`
	StartTime   time.Time            `json:"startTime"`
	EndTime     time.Time            `json:"endTime"`
	DeletedAt   *time.Time           `json:"deletedAt,omitempty"`
	TruncatedAt *time.Time           `json:"truncatedAt,omitempty"`
	Shards      []ManifestShardEntry `json:"shards"`
}

type ManifestShardEntry struct {
	ID          int64        `json:"id"`
	ShardOwners []ShardOwner `json:"shardOwners"`
	ManifestFileEntry
}

type ShardOwner struct {
	NodeID int64 `json:"nodeID"`
}

type ManifestSubscription struct {
	Name         string   `json:"name"`
	Mode         string   `json:"mode"`
	Destinations []string `json:"destinations"`
}
