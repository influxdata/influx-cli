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
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBucketsDelete(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		configOrgName              string
		params                     bucket.BucketsDeleteParams
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPattern      string
		expectedInErr              string
	}{
		{
			name:          "by ID",
			configOrgName: "my-default-org",
			params: bucket.BucketsDeleteParams{
				ID: "123",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "123", *in.GetId()) &&
						assert.Nil(t, in.GetName()) &&
						assert.Nil(t, in.GetOrgID()) &&
						assert.Nil(t, in.GetOrg())
				})).Return(api.Buckets{
					Buckets: &[]api.Bucket{
						{
							Id:    api.PtrString("123"),
							Name:  "my-bucket",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
						},
					},
				}, nil)

				bucketsApi.EXPECT().DeleteBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiDeleteBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().DeleteBucketsIDExecute(tmock.MatchedBy(func(in api.ApiDeleteBucketsIDRequest) bool {
					return assert.Equal(t, "123", in.GetBucketID())
				})).Return(nil)
			},
			expectedStdoutPattern: `123\s+my-bucket\s+1h0m0s\s+n/a\s+456\s+implicit`,
		},
		{
			name:          "by name and org ID",
			configOrgName: "my-default-org",
			params: bucket.BucketsDeleteParams{
				Name:  "my-bucket",
				OrgID: "456",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Nil(t, in.GetId()) &&
						assert.Equal(t, "my-bucket", *in.GetName()) &&
						assert.Equal(t, "456", *in.GetOrgID()) &&
						assert.Nil(t, in.GetOrg())
				})).Return(api.Buckets{
					Buckets: &[]api.Bucket{
						{
							Id:    api.PtrString("123"),
							Name:  "my-bucket",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
						},
					},
				}, nil)

				bucketsApi.EXPECT().DeleteBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiDeleteBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().DeleteBucketsIDExecute(tmock.MatchedBy(func(in api.ApiDeleteBucketsIDRequest) bool {
					return assert.Equal(t, "123", in.GetBucketID())
				})).Return(nil)
			},
			expectedStdoutPattern: `123\s+my-bucket\s+1h0m0s\s+n/a\s+456\s+implicit`,
		},
		{
			name:          "by name and org name",
			configOrgName: "my-default-org",
			params: bucket.BucketsDeleteParams{
				Name:    "my-bucket",
				OrgName: "my-org",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Nil(t, in.GetId()) &&
						assert.Equal(t, "my-bucket", *in.GetName()) &&
						assert.Nil(t, in.GetOrgID()) &&
						assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Buckets{
					Buckets: &[]api.Bucket{
						{
							Id:    api.PtrString("123"),
							Name:  "my-bucket",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
						},
					},
				}, nil)

				bucketsApi.EXPECT().DeleteBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiDeleteBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().DeleteBucketsIDExecute(tmock.MatchedBy(func(in api.ApiDeleteBucketsIDRequest) bool {
					return assert.Equal(t, "123", in.GetBucketID())
				})).Return(nil)
			},
			expectedStdoutPattern: `123\s+my-bucket\s+1h0m0s\s+n/a\s+456\s+implicit`,
		},
		{
			name:          "by name in default org",
			configOrgName: "my-default-org",
			params: bucket.BucketsDeleteParams{
				Name: "my-bucket",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Nil(t, in.GetId()) &&
						assert.Equal(t, "my-bucket", *in.GetName()) &&
						assert.Nil(t, in.GetOrgID()) &&
						assert.Equal(t, "my-default-org", *in.GetOrg())
				})).Return(api.Buckets{
					Buckets: &[]api.Bucket{
						{
							Id:    api.PtrString("123"),
							Name:  "my-bucket",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
						},
					},
				}, nil)

				bucketsApi.EXPECT().DeleteBucketsID(gomock.Any(), gomock.Eq("123")).
					Return(api.ApiDeleteBucketsIDRequest{ApiService: bucketsApi}.BucketID("123"))
				bucketsApi.EXPECT().DeleteBucketsIDExecute(tmock.MatchedBy(func(in api.ApiDeleteBucketsIDRequest) bool {
					return assert.Equal(t, "123", in.GetBucketID())
				})).Return(nil)
			},
			expectedStdoutPattern: `123\s+my-bucket\s+1h0m0s\s+n/a\s+456\s+implicit`,
		},
		{
			name: "by name without org",
			params: bucket.BucketsDeleteParams{
				Name: "my-bucket",
			},
			expectedInErr: "must specify org ID or org name",
		},
		{
			name: "no such bucket",
			params: bucket.BucketsDeleteParams{
				ID: "123",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(gomock.Any()).Return(api.Buckets{}, nil)
			},
			expectedInErr: "not found",
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
				CLI:        clients.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio},
				BucketsApi: client,
			}

			err := cli.Delete(context.Background(), &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, writtenBytes.String())
				return
			}
			require.NoError(t, err)
			testutils.MatchLines(t, []string{
				`ID\s+Name\s+Retention\s+Shard group duration\s+Organization ID\s+Schema Type\s+Deleted`,
				tc.expectedStdoutPattern + `\s+true`,
			}, strings.Split(writtenBytes.String(), "\n"))
		})
	}
}
