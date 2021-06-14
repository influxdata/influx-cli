package backup_restore

import (
	"fmt"
	"time"
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

const ManifestExtension = "manifest"

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
