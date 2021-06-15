package restore

import (
	"github.com/influxdata/influx-cli/v2/api"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
)

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
