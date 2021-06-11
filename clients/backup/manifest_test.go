package backup_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients/backup"
	"github.com/stretchr/testify/require"
)

func TestConvertBucketManifest(t *testing.T) {
	t.Parallel()

	now := time.Now()

	manifest := api.BucketMetadataManifest{
		OrganizationID:         "123",
		OrganizationName:       "org",
		BucketID:               "456",
		BucketName:             "bucket",
		DefaultRetentionPolicy: "foo",
		RetentionPolicies: []api.RetentionPolicyManifest{
			{
				Name:               "foo",
				ReplicaN:           1,
				Duration:           100,
				ShardGroupDuration: 10,
				ShardGroups: []api.ShardGroupManifest{
					{
						Id:          1,
						StartTime:   now,
						EndTime:     now,
						TruncatedAt: &now,
						Shards: []api.ShardManifest{
							{
								Id:          10,
								ShardOwners: []api.ShardOwner{{NodeID: 1}},
							},
							{
								Id:          20,
								ShardOwners: []api.ShardOwner{{NodeID: 2}, {NodeID: 3}},
							},
						},
					},
					{
						Id:        2,
						StartTime: now,
						EndTime:   now,
						DeletedAt: &now,
						Shards: []api.ShardManifest{
							{
								Id: 30,
							},
						},
					},
				},
				Subscriptions: []api.SubscriptionManifest{},
			},
			{
				Name:               "bar",
				ReplicaN:           3,
				Duration:           9999,
				ShardGroupDuration: 1,
				ShardGroups: []api.ShardGroupManifest{
					{
						Id:        3,
						StartTime: now,
						EndTime:   now,
						Shards:    []api.ShardManifest{},
					},
				},
				Subscriptions: []api.SubscriptionManifest{
					{
						Name:         "test",
						Mode:         "on",
						Destinations: []string{"here", "there", "everywhere"},
					},
					{
						Name:         "test2",
						Mode:         "off",
						Destinations: []string{},
					},
				},
			},
		},
	}

	fakeGetShard := func(id int64) (*backup.ManifestFileEntry, error) {
		if id == 20 {
			return nil, nil
		}
		return &backup.ManifestFileEntry{
			FileName:    fmt.Sprintf("%d.gz", id),
			Size:        id * 100,
			Compression: backup.GzipCompression,
		}, nil
	}

	converted, err := backup.ConvertBucketManifest(manifest, fakeGetShard)
	require.NoError(t, err)

	expected := backup.ManifestBucketEntry{
		OrganizationID:         "123",
		OrganizationName:       "org",
		BucketID:               "456",
		BucketName:             "bucket",
		DefaultRetentionPolicy: "foo",
		RetentionPolicies: []backup.ManifestRetentionPolicy{
			{
				Name:               "foo",
				ReplicaN:           1,
				Duration:           100,
				ShardGroupDuration: 10,
				ShardGroups: []backup.ManifestShardGroup{
					{
						ID:          1,
						StartTime:   now,
						EndTime:     now,
						TruncatedAt: &now,
						Shards: []backup.ManifestShardEntry{
							{
								ID:          10,
								ShardOwners: []backup.ShardOwner{{NodeID: 1}},
								ManifestFileEntry: backup.ManifestFileEntry{
									FileName:    "10.gz",
									Size:        1000,
									Compression: backup.GzipCompression,
								},
							},
						},
					},
					{
						ID:        2,
						StartTime: now,
						EndTime:   now,
						DeletedAt: &now,
						Shards: []backup.ManifestShardEntry{
							{
								ID:          30,
								ShardOwners: []backup.ShardOwner{},
								ManifestFileEntry: backup.ManifestFileEntry{
									FileName:    "30.gz",
									Size:        3000,
									Compression: backup.GzipCompression,
								},
							},
						},
					},
				},
				Subscriptions: []backup.ManifestSubscription{},
			},
			{
				Name:               "bar",
				ReplicaN:           3,
				Duration:           9999,
				ShardGroupDuration: 1,
				ShardGroups: []backup.ManifestShardGroup{
					{
						ID:        3,
						StartTime: now,
						EndTime:   now,
						Shards:    []backup.ManifestShardEntry{},
					},
				},
				Subscriptions: []backup.ManifestSubscription{
					{
						Name:         "test",
						Mode:         "on",
						Destinations: []string{"here", "there", "everywhere"},
					},
					{
						Name:         "test2",
						Mode:         "off",
						Destinations: []string{},
					},
				},
			},
		},
	}

	require.Equal(t, expected, converted)
}
