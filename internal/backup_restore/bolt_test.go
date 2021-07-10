package backup_restore_test

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"github.com/stretchr/testify/require"
)

const testFile = "testdata/test.bolt.gz"

func TestExtractManifest(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "windows" {
		t.Skip("skipping test on Windows: https://github.com/etcd-io/bbolt/issues/252")
	}

	// Extract our example input into a format the bbolt client can use.
	boltIn, err := os.Open(testFile)
	require.NoError(t, err)
	defer boltIn.Close()
	gzipIn, err := gzip.NewReader(boltIn)
	require.NoError(t, err)
	defer gzipIn.Close()

	tmp, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)

	tmpBoltPath := filepath.Join(tmp, "test.bolt")
	tmpBolt, err := os.Create(tmpBoltPath)
	require.NoError(t, err)

	_, err = io.Copy(tmpBolt, gzipIn)
	require.NoError(t, tmpBolt.Close())
	require.NoError(t, err)

	extracted, err := backup_restore.ExtractBucketMetadata(tmpBoltPath)
	require.NoError(t, err)

	expected := []api.BucketMetadataManifest{
		{
			OrganizationID:         "80c29010030b3d83",
			OrganizationName:       "test2",
			BucketID:               "5379ef2d1655b0ab",
			BucketName:             "test3",
			DefaultRetentionPolicy: "autogen",
			RetentionPolicies: []api.RetentionPolicyManifest{
				{
					Name:               "autogen",
					ReplicaN:           1,
					Duration:           0,
					ShardGroupDuration: 604800000000000,
					ShardGroups:        []api.ShardGroupManifest{},
					Subscriptions:      []api.SubscriptionManifest{},
				},
			},
		},
		{
			OrganizationID:         "375477729f9d7262",
			OrganizationName:       "test",
			BucketID:               "cce01ef3783e3678",
			BucketName:             "test2",
			DefaultRetentionPolicy: "autogen",
			RetentionPolicies: []api.RetentionPolicyManifest{
				{
					Name:               "autogen",
					ReplicaN:           1,
					Duration:           0,
					ShardGroupDuration: 3600000000000,
					ShardGroups:        []api.ShardGroupManifest{},
					Subscriptions:      []api.SubscriptionManifest{},
				},
			},
		},
		{
			OrganizationID:         "375477729f9d7262",
			OrganizationName:       "test",
			BucketID:               "d66c5360b5aa91b4",
			BucketName:             "test",
			DefaultRetentionPolicy: "autogen",
			RetentionPolicies: []api.RetentionPolicyManifest{
				{
					Name:               "autogen",
					ReplicaN:           1,
					Duration:           259200000000000,
					ShardGroupDuration: 86400000000000,
					ShardGroups:        []api.ShardGroupManifest{},
					Subscriptions:      []api.SubscriptionManifest{},
				},
			},
		},
	}

	require.Equal(t, len(expected), len(extracted))
	require.Equal(t, expected, extracted)
	for _, e := range expected {
		require.Contains(t, extracted, e)
	}
}
