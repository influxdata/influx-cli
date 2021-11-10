package bucket_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/bucket"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBucketsUpdate(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		params                     bucket.BucketsUpdateParams
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPattern      string
	}{
		{
			name: "name",
			params: bucket.BucketsUpdateParams{
				BucketParams: clients.BucketParams{
					BucketID:   "123",
					BucketName: "cold-storage",
				},
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PatchBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiPatchBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().PatchBucketsIDExecute(tmock.MatchedBy(func(in api.ApiPatchBucketsIDRequest) bool {
					body := in.GetPatchBucketRequest()
					return assert.Equal(t, "123", in.GetBucketID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, "cold-storage", body.GetName()) &&
						assert.Nil(t, body.Description) &&
						assert.Empty(t, body.GetRetentionRules())
				})).Return(api.Bucket{
					Id:    api.PtrString("123"),
					Name:  "cold-storage",
					OrgID: api.PtrString("456"),
				}, nil)
			},
			expectedStdoutPattern: `123\s+cold-storage\s+infinite\s+n/a\s+456`,
		},
		{
			name: "description",
			params: bucket.BucketsUpdateParams{
				BucketParams: clients.BucketParams{BucketID: "123"},
				Description:  "a very useful description",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PatchBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiPatchBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().PatchBucketsIDExecute(tmock.MatchedBy(func(in api.ApiPatchBucketsIDRequest) bool {
					body := in.GetPatchBucketRequest()
					return assert.Equal(t, "123", in.GetBucketID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, "a very useful description", body.GetDescription()) &&
						assert.Nil(t, body.Name) &&
						assert.Empty(t, body.GetRetentionRules())
				})).Return(api.Bucket{
					Id:          api.PtrString("123"),
					Name:        "my-bucket",
					Description: api.PtrString("a very useful description"),
					OrgID:       api.PtrString("456"),
				}, nil)
			},
			expectedStdoutPattern: `123\s+my-bucket\s+infinite\s+n/a\s+456`,
		},
		{
			name: "retention",
			params: bucket.BucketsUpdateParams{
				BucketParams: clients.BucketParams{BucketID: "123"},
				Retention:    "3w",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PatchBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiPatchBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().PatchBucketsIDExecute(tmock.MatchedBy(func(in api.ApiPatchBucketsIDRequest) bool {
					body := in.GetPatchBucketRequest()
					return assert.Equal(t, "123", in.GetBucketID()) &&
						assert.NotNil(t, body) &&
						assert.Nil(t, body.Name) &&
						assert.Nil(t, body.Description) &&
						assert.Len(t, body.GetRetentionRules(), 1) &&
						assert.Nil(t, body.GetRetentionRules()[0].ShardGroupDurationSeconds) &&
						assert.Equal(t, int64(3*7*24*3600), *body.GetRetentionRules()[0].EverySeconds)
				})).Return(api.Bucket{
					Id:    api.PtrString("123"),
					Name:  "my-bucket",
					OrgID: api.PtrString("456"),
					RetentionRules: []api.RetentionRule{
						{EverySeconds: int64(3 * 7 * 24 * 3600)},
					},
				}, nil)
			},
			expectedStdoutPattern: `123\s+my-bucket\s+504h0m0s\s+n/a\s+456`,
		},
		{
			name: "shard-group duration",
			params: bucket.BucketsUpdateParams{
				BucketParams:       clients.BucketParams{BucketID: "123"},
				ShardGroupDuration: "10h30m",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PatchBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiPatchBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().PatchBucketsIDExecute(tmock.MatchedBy(func(in api.ApiPatchBucketsIDRequest) bool {
					body := in.GetPatchBucketRequest()
					return assert.Equal(t, "123", in.GetBucketID()) &&
						assert.NotNil(t, body) &&
						assert.Nil(t, body.Name) &&
						assert.Nil(t, body.Description) &&
						assert.Len(t, body.GetRetentionRules(), 1) &&
						assert.Nil(t, body.GetRetentionRules()[0].EverySeconds) &&
						assert.Equal(t, int64(10*3600+30*60), *body.GetRetentionRules()[0].ShardGroupDurationSeconds)
				})).Return(api.Bucket{
					Id:    api.PtrString("123"),
					Name:  "my-bucket",
					OrgID: api.PtrString("456"),
					RetentionRules: []api.RetentionRule{
						{ShardGroupDurationSeconds: api.PtrInt64(10*3600 + 30*60)},
					},
				}, nil)
			},
			expectedStdoutPattern: `123\s+my-bucket\s+infinite\s+10h30m0s\s+456`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()
			client := mock.NewMockBucketsApi(ctrl)
			if tc.registerBucketExpectations != nil {
				tc.registerBucketExpectations(t, client)
			}
			cli := bucket.Client{
				CLI:        clients.CLI{StdIO: stdio},
				BucketsApi: client,
			}

			err := cli.Update(context.Background(), &tc.params)
			require.NoError(t, err)
			testutils.MatchLines(t, []string{
				`ID\s+Name\s+Retention\s+Shard group duration\s+Organization ID\s+Schema Type`,
				tc.expectedStdoutPattern,
			}, strings.Split(writtenBytes.String(), "\n"))
		})
	}
}
