package backup_restore

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc --go_out=. meta.proto

// NOTE: An unfortunate naming collision below. Bolt calls its databases "buckets".
// These are the names that were used in the metadata DB for 2.0.x versions of influxdb.
var (
	bucketsBoltBucket       = []byte("bucketsv1")
	organizationsBoltBucket = []byte("organizationsv1")
	v1MetadataBoltBucket    = []byte("v1_tsm1_metadata")
	v1MetadataBoltKey       = []byte("meta.db")
)

// influxdbBucketSchema models the JSON structure used by InfluxDB 2.0.x to serialize
// bucket metadata in the embedded KV store.
type influxdbBucketSchema struct {
	ID                 string        `json:"id"`
	OrgID              string        `json:"orgID"`
	Type               int           `json:"type"`
	Name               string        `json:"name"`
	Description        *string       `json:"description,omitempty"`
	RetentionPeriod    time.Duration `json:"retentionPeriod"`
	ShardGroupDuration time.Duration `json:"ShardGroupDuration"`
}

// influxdbOrganizationSchema models the JSON structure used by InfluxDB 2.0.x to serialize
// organization metadata in the embedded KV store.
type influxdbOrganizationSchema struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExtractBucketMetadata reads a boltdb backed up from InfluxDB 2.0.x, converting a subset of the
// metadata it contains into a set of 2.1.x bucket manifests.
func ExtractBucketMetadata(boltPath string) ([]api.BucketMetadataManifest, error) {
	db, err := bbolt.Open(boltPath, 0666, &bbolt.Options{ReadOnly: true, Timeout: 1 * time.Second})
	if err != nil {
		// Hack to give a slightly nicer error message for a known failure mode when bolt calls
		// mmap on a file system that doesn't support the MAP_SHARED option.
		//
		// See: https://github.com/boltdb/bolt/issues/272
		// See: https://stackoverflow.com/a/18421071
		if err.Error() == "invalid argument" {
			return nil, fmt.Errorf("unable to open boltdb: mmap of %q may not support the MAP_SHARED option", boltPath)
		}

		return nil, fmt.Errorf("unable to open boltdb: %w", err)
	}
	defer db.Close()

	// Read raw metadata needed to construct a manifest.
	var buckets []influxdbBucketSchema
	orgNamesById := map[string]string{}
	dbInfoByBucketId := map[string]DatabaseInfo{}

	if err := db.View(func(tx *bbolt.Tx) error {
		bucketDB := tx.Bucket(bucketsBoltBucket)
		if bucketDB == nil {
			return errors.New("bucket metadata not found in local KV store")
		}

		if err := bucketDB.ForEach(func(k, v []byte) error {
			var b influxdbBucketSchema
			if err := json.Unmarshal(v, &b); err != nil {
				return err
			}
			if b.Type != 1 { // 1 == "system"
				buckets = append(buckets, b)
			}
			return nil
		}); err != nil {
			return fmt.Errorf("failed to read bucket metadata from local KV store: %w", err)
		}

		orgsDB := tx.Bucket(organizationsBoltBucket)
		if orgsDB == nil {
			return errors.New("organization metadata not found in local KV store")
		}

		if err := orgsDB.ForEach(func(k, v []byte) error {
			var o influxdbOrganizationSchema
			if err := json.Unmarshal(v, &o); err != nil {
				return err
			}
			orgNamesById[o.ID] = o.Name
			return nil
		}); err != nil {
			return fmt.Errorf("failed to read organization metadata from local KV store: %w", err)
		}

		v1DB := tx.Bucket(v1MetadataBoltBucket)
		if v1DB == nil {
			return errors.New("v1 database info not found in local KV store")
		}
		fullMeta := v1DB.Get(v1MetadataBoltKey)
		if fullMeta == nil {
			return errors.New("v1 database info not found in local KV store")
		}

		var pb Data
		if err := proto.Unmarshal(fullMeta, &pb); err != nil {
			return fmt.Errorf("failed to unmarshal v1 database info: %w", err)
		}
		for _, rawDBI := range pb.GetDatabases() {
			dbInfoByBucketId[*rawDBI.Name] = *rawDBI
		}

		return nil
	}); err != nil {
		return nil, err
	}

	manifests := make([]api.BucketMetadataManifest, len(buckets))
	for i, b := range buckets {
		orgName, ok := orgNamesById[b.OrgID]
		if !ok {
			return nil, fmt.Errorf("local KV store in inconsistent state: no organization found with ID %q", b.OrgID)
		}
		dbi, ok := dbInfoByBucketId[b.ID]
		if !ok {
			return nil, fmt.Errorf("local KV store in inconsistent state: no V1 database info found for bucket %q", b.Name)
		}
		manifests[i] = combineMetadata(b, orgName, dbi)
	}

	return manifests, nil
}

func combineMetadata(bucket influxdbBucketSchema, orgName string, dbi DatabaseInfo) api.BucketMetadataManifest {
	m := api.BucketMetadataManifest{
		OrganizationID:         bucket.OrgID,
		OrganizationName:       orgName,
		BucketID:               bucket.ID,
		BucketName:             bucket.Name,
		DefaultRetentionPolicy: *dbi.DefaultRetentionPolicy,
		RetentionPolicies:      make([]api.RetentionPolicyManifest, len(dbi.RetentionPolicies)),
	}
	if bucket.Description != nil && *bucket.Description != "" {
		m.Description = bucket.Description
	}
	for i, rp := range dbi.RetentionPolicies {
		m.RetentionPolicies[i] = convertRPI(*rp)
	}
	return m
}

func convertRPI(rpi RetentionPolicyInfo) api.RetentionPolicyManifest {
	m := api.RetentionPolicyManifest{
		Name:               *rpi.Name,
		ReplicaN:           int32(*rpi.ReplicaN),
		Duration:           *rpi.Duration,
		ShardGroupDuration: *rpi.ShardGroupDuration,
		ShardGroups:        make([]api.ShardGroupManifest, len(rpi.ShardGroups)),
		Subscriptions:      make([]api.SubscriptionManifest, len(rpi.Subscriptions)),
	}
	for i, sg := range rpi.ShardGroups {
		m.ShardGroups[i] = convertSGI(*sg)
	}
	for i, s := range rpi.Subscriptions {
		m.Subscriptions[i] = api.SubscriptionManifest{
			Name:         *s.Name,
			Mode:         *s.Mode,
			Destinations: s.Destinations,
		}
	}
	return m
}

func convertSGI(sgi ShardGroupInfo) api.ShardGroupManifest {
	var deleted, truncated *time.Time
	if sgi.DeletedAt != nil {
		d := time.Unix(0, *sgi.DeletedAt).UTC()
		deleted = &d
	}
	if sgi.TruncatedAt != nil {
		t := time.Unix(0, *sgi.TruncatedAt).UTC()
		truncated = &t
	}

	m := api.ShardGroupManifest{
		Id:          int64(*sgi.ID),
		StartTime:   time.Unix(0, *sgi.StartTime).UTC(),
		EndTime:     time.Unix(0, *sgi.EndTime).UTC(),
		DeletedAt:   deleted,
		TruncatedAt: truncated,
		Shards:      make([]api.ShardManifest, len(sgi.Shards)),
	}
	for i, s := range sgi.Shards {
		m.Shards[i] = convertShard(*s)
	}
	return m
}

func convertShard(shard ShardInfo) api.ShardManifest {
	m := api.ShardManifest{
		Id:          int64(*shard.ID),
		ShardOwners: make([]api.ShardOwner, len(shard.Owners)),
	}
	for i, o := range shard.Owners {
		m.ShardOwners[i] = api.ShardOwner{NodeID: int64(*o.NodeID)}
	}
	return m
}
