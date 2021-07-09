package backup

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/influxdata/influx-cli/v2/api"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"go.etcd.io/bbolt"
)

//go:generate protoc --gogo_out=. internal/meta.proto

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

// influxdbV1DatabaseInfo models the protobuf structure used by InfluxDB 2.0.x to serialize
// database info in the embedded KV store.
type influxdbV1DatabaseInfo struct {
	Name                   string
	DefaultRetentionPolicy string
	RetentionPolicies      []influxdbV1RetentionPolicyInfo
}

// influxdbV1RetentionPolicyInfo models the protobuf structure used by InfluxDB 2.0.x to serialize
// retention-policy info in the embedded KV store.
type influxdbV1RetentionPolicyInfo struct {
	Name               string
	ReplicaN           int32
	Duration           int64
	ShardGroupDuration int64
	ShardGroups        []influxdbV1ShardGroupInfo
	Subscriptions      []influxdbV1SubscriptionInfo
}

// influxdbV1ShardGroupInfo models the protobuf structure used by InfluxDB 2.0.x to serialize
// shard-group info in the embedded KV store.
type influxdbV1ShardGroupInfo struct {
	ID          int64
	StartTime   time.Time
	EndTime     time.Time
	DeletedAt   time.Time
	Shards      []influxdbV1ShardInfo
	TruncatedAt time.Time
}

// influxdbV1ShardInfo models the protobuf structure used by InfluxDB 2.0.x to serialize
// shard info in the embedded KV store.
type influxdbV1ShardInfo struct {
	ID     int64
	Owners []influxdbV1ShardOwnerInfo
}

// inflxudbV1ShardOwnerInfo models the protobuf structure used by InfluxDB 2.0.x to serialize
// shard-owner info in the embedded KV store.
type influxdbV1ShardOwnerInfo struct {
	NodeID int64
}

// influxdbV1SubscriptionInfo models the protobuf structure used by InfluxDB 2.0.x to serialize
// subscription info in the embedded KV store.
type influxdbV1SubscriptionInfo struct {
	Name         string
	Mode         string
	Destinations []string
}

// extractBucketManifest reads a boltdb backed up from InfluxDB 2.0.x, converting a subset of the
// metadata it contains into a set of 2.1.x bucket manifests.
func extractBucketManifest(boltPath string) ([]api.BucketMetadataManifest, error) {
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
	dbInfoByBucketId := map[string]influxdbV1DatabaseInfo{}

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

		var pb br.Data
		if err := proto.Unmarshal(fullMeta, &pb); err != nil {
			return fmt.Errorf("failed to unmarshal v1 database info: %w", err)
		}
		for _, rawDBI := range pb.GetDatabases() {
			if rawDBI == nil {
				continue
			}
			unmarshalled := unmarshalRawDBI(*rawDBI)
			dbInfoByBucketId[unmarshalled.Name] = unmarshalled
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

func unmarshalRawDBI(rawDBI br.DatabaseInfo) influxdbV1DatabaseInfo {
	dbi := influxdbV1DatabaseInfo{
		Name:                   rawDBI.GetName(),
		DefaultRetentionPolicy: rawDBI.GetDefaultRetentionPolicy(),
		RetentionPolicies:      make([]influxdbV1RetentionPolicyInfo, 0, len(rawDBI.GetRetentionPolicies())),
	}
	for _, rp := range rawDBI.GetRetentionPolicies() {
		if rp == nil {
			continue
		}
		dbi.RetentionPolicies = append(dbi.RetentionPolicies, unmarshalRawRPI(*rp))
	}
	return dbi
}

func unmarshalRawRPI(rawRPI br.RetentionPolicyInfo) influxdbV1RetentionPolicyInfo {
	rpi := influxdbV1RetentionPolicyInfo{
		Name:               rawRPI.GetName(),
		ReplicaN:           int32(rawRPI.GetReplicaN()),
		Duration:           rawRPI.GetDuration(),
		ShardGroupDuration: rawRPI.GetShardGroupDuration(),
		ShardGroups:        make([]influxdbV1ShardGroupInfo, 0, len(rawRPI.GetShardGroups())),
		Subscriptions:      make([]influxdbV1SubscriptionInfo, 0, len(rawRPI.GetSubscriptions())),
	}
	for _, sg := range rawRPI.GetShardGroups() {
		if sg == nil {
			continue
		}
		rpi.ShardGroups = append(rpi.ShardGroups, unmarshalRawSGI(*sg))
	}
	for _, s := range rawRPI.GetSubscriptions() {
		if s == nil {
			continue
		}
		rpi.Subscriptions = append(rpi.Subscriptions, influxdbV1SubscriptionInfo{
			Name:         s.GetName(),
			Mode:         s.GetMode(),
			Destinations: s.GetDestinations(),
		})
	}
	return rpi
}

func unmarshalRawSGI(rawSGI br.ShardGroupInfo) influxdbV1ShardGroupInfo {
	sgi := influxdbV1ShardGroupInfo{
		ID:          int64(rawSGI.GetID()),
		StartTime:   time.Unix(0, rawSGI.GetStartTime()).UTC(),
		EndTime:     time.Unix(0, rawSGI.GetEndTime()).UTC(),
		DeletedAt:   time.Unix(0, rawSGI.GetDeletedAt()).UTC(),
		Shards:      make([]influxdbV1ShardInfo, 0, len(rawSGI.GetShards())),
		TruncatedAt: time.Unix(0, rawSGI.GetTruncatedAt()).UTC(),
	}
	for _, s := range rawSGI.GetShards() {
		if s == nil {
			continue
		}
		sgi.Shards = append(sgi.Shards, unmarshalRawShard(*s))
	}
	return sgi
}

func unmarshalRawShard(rawShard br.ShardInfo) influxdbV1ShardInfo {
	si := influxdbV1ShardInfo{
		ID: int64(rawShard.GetID()),
	}
	// If deprecated "OwnerIDs" exists then convert it to "Owners" format.
	//lint:ignore SA1019 we need to check for the presence of the deprecated field so we can convert it
	oldStyleOwnerIds := rawShard.GetOwnerIDs()
	if len(oldStyleOwnerIds) > 0 {
		si.Owners = make([]influxdbV1ShardOwnerInfo, len(oldStyleOwnerIds))
		for i, oid := range oldStyleOwnerIds {
			si.Owners[i] = influxdbV1ShardOwnerInfo{NodeID: int64(oid)}
		}
	} else {
		si.Owners = make([]influxdbV1ShardOwnerInfo, 0, len(rawShard.GetOwners()))
		for _, o := range rawShard.GetOwners() {
			if o == nil {
				continue
			}
			si.Owners = append(si.Owners, influxdbV1ShardOwnerInfo{NodeID: int64(o.GetNodeID())})
		}
	}
	return si
}

func combineMetadata(bucket influxdbBucketSchema, orgName string, dbi influxdbV1DatabaseInfo) api.BucketMetadataManifest {
	m := api.BucketMetadataManifest{
		OrganizationID:         bucket.OrgID,
		OrganizationName:       orgName,
		BucketID:               bucket.ID,
		BucketName:             bucket.Name,
		DefaultRetentionPolicy: dbi.DefaultRetentionPolicy,
		RetentionPolicies:      make([]api.RetentionPolicyManifest, len(dbi.RetentionPolicies)),
	}
	if bucket.Description != nil && *bucket.Description != "" {
		m.Description = bucket.Description
	}
	for i, rp := range dbi.RetentionPolicies {
		m.RetentionPolicies[i] = convertRPI(rp)
	}
	return m
}

func convertRPI(rpi influxdbV1RetentionPolicyInfo) api.RetentionPolicyManifest {
	m := api.RetentionPolicyManifest{
		Name:               rpi.Name,
		ReplicaN:           rpi.ReplicaN,
		Duration:           rpi.Duration,
		ShardGroupDuration: rpi.ShardGroupDuration,
		ShardGroups:        make([]api.ShardGroupManifest, len(rpi.ShardGroups)),
		Subscriptions:      make([]api.SubscriptionManifest, len(rpi.Subscriptions)),
	}
	for i, sg := range rpi.ShardGroups {
		m.ShardGroups[i] = convertSGI(sg)
	}
	for i, s := range rpi.Subscriptions {
		m.Subscriptions[i] = api.SubscriptionManifest{
			Name:         s.Name,
			Mode:         s.Mode,
			Destinations: s.Destinations,
		}
	}
	return m
}

func convertSGI(sgi influxdbV1ShardGroupInfo) api.ShardGroupManifest {
	m := api.ShardGroupManifest{
		Id:        sgi.ID,
		StartTime: sgi.StartTime,
		EndTime:   sgi.EndTime,
		Shards:    make([]api.ShardManifest, len(sgi.Shards)),
	}
	if sgi.DeletedAt.Unix() != 0 {
		m.DeletedAt = &sgi.DeletedAt
	}
	if sgi.TruncatedAt.Unix() != 0 {
		m.TruncatedAt = &sgi.TruncatedAt
	}
	for i, s := range sgi.Shards {
		m.Shards[i] = convertShard(s)
	}
	return m
}

func convertShard(shard influxdbV1ShardInfo) api.ShardManifest {
	m := api.ShardManifest{
		Id:          shard.ID,
		ShardOwners: make([]api.ShardOwner, len(shard.Owners)),
	}
	for i, o := range shard.Owners {
		m.ShardOwners[i] = api.ShardOwner{NodeID: o.NodeID}
	}
	return m
}
