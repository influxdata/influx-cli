package backup

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
)

func ConvertBucketManifest(manifest api.BucketMetadataManifest, getShard func(shardId int64) (*br.ManifestFileEntry, error)) (br.ManifestBucketEntry, error) {
	m := br.ManifestBucketEntry{
		OrganizationID:         manifest.OrganizationID,
		OrganizationName:       manifest.OrganizationName,
		BucketID:               manifest.BucketID,
		BucketName:             manifest.BucketName,
		DefaultRetentionPolicy: manifest.DefaultRetentionPolicy,
		RetentionPolicies:      make([]br.ManifestRetentionPolicy, len(manifest.RetentionPolicies)),
	}

	for i, rp := range manifest.RetentionPolicies {
		var err error
		m.RetentionPolicies[i], err = ConvertRetentionPolicy(rp, getShard)
		if err != nil {
			return br.ManifestBucketEntry{}, err
		}
	}

	return m, nil
}

func ConvertRetentionPolicy(manifest api.RetentionPolicyManifest, getShard func(shardId int64) (*br.ManifestFileEntry, error)) (br.ManifestRetentionPolicy, error) {
	m := br.ManifestRetentionPolicy{
		Name:               manifest.Name,
		ReplicaN:           manifest.ReplicaN,
		Duration:           manifest.Duration,
		ShardGroupDuration: manifest.ShardGroupDuration,
		ShardGroups:        make([]br.ManifestShardGroup, len(manifest.ShardGroups)),
		Subscriptions:      make([]br.ManifestSubscription, len(manifest.Subscriptions)),
	}

	for i, sg := range manifest.ShardGroups {
		var err error
		m.ShardGroups[i], err = ConvertShardGroup(sg, getShard)
		if err != nil {
			return br.ManifestRetentionPolicy{}, err
		}
	}

	for i, s := range manifest.Subscriptions {
		m.Subscriptions[i] = br.ManifestSubscription{
			Name:         s.Name,
			Mode:         s.Mode,
			Destinations: s.Destinations,
		}
	}

	return m, nil
}

func ConvertShardGroup(manifest api.ShardGroupManifest, getShard func(shardId int64) (*br.ManifestFileEntry, error)) (br.ManifestShardGroup, error) {
	m := br.ManifestShardGroup{
		ID:          manifest.Id,
		StartTime:   manifest.StartTime,
		EndTime:     manifest.EndTime,
		DeletedAt:   manifest.DeletedAt,
		TruncatedAt: manifest.TruncatedAt,
		Shards:      make([]br.ManifestShardEntry, 0, len(manifest.Shards)),
	}

	for _, sh := range manifest.Shards {
		maybeShard, err := ConvertShard(sh, getShard)
		if err != nil {
			return br.ManifestShardGroup{}, err
		}
		// Shard deleted mid-backup.
		if maybeShard == nil {
			continue
		}
		m.Shards = append(m.Shards, *maybeShard)
	}

	return m, nil
}

func ConvertShard(manifest api.ShardManifest, getShard func(shardId int64) (*br.ManifestFileEntry, error)) (*br.ManifestShardEntry, error) {
	shardFileInfo, err := getShard(manifest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to download snapshot of shard %d: %w", manifest.Id, err)
	}
	if shardFileInfo == nil {
		return nil, nil
	}

	m := br.ManifestShardEntry{
		ID:                manifest.Id,
		ShardOwners:       make([]br.ShardOwner, len(manifest.ShardOwners)),
		ManifestFileEntry: *shardFileInfo,
	}

	for i, o := range manifest.ShardOwners {
		m.ShardOwners[i] = br.ShardOwner{
			NodeID: o.NodeID,
		}
	}

	return &m, nil
}
