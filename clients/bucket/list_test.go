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
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBucketsList(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		configOrgName              string
		params                     bucket.BucketsListParams
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPatterns     []string
		expectedInErr              string
	}{
		{
			name: "by ID",
			params: bucket.BucketsListParams{
				OrgBucketParams: clients.OrgBucketParams{
					BucketParams: clients.BucketParams{BucketID: "123"},
				},
			},
			configOrgName: "my-default-org",
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "123", *in.GetId()) &&
						assert.Equal(t, "my-default-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetName()) &&
						assert.Nil(t, in.GetOrgID())
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
			},
			expectedStdoutPatterns: []string{
				`123\s+my-bucket\s+1h0m0s\s+n/a\s+456`,
			},
		},
		{
			name: "by name",
			params: bucket.BucketsListParams{
				OrgBucketParams: clients.OrgBucketParams{
					BucketParams: clients.BucketParams{BucketName: "my-bucket"},
				},
			},
			configOrgName: "my-default-org",
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "my-bucket", *in.GetName()) &&
						assert.Equal(t, "my-default-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetId()) &&
						assert.Nil(t, in.GetOrgID())
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
			},
			expectedStdoutPatterns: []string{
				`123\s+my-bucket\s+1h0m0s\s+n/a\s+456`,
			},
		},
		{
			name: "pagination via 'offset'",
			params: bucket.BucketsListParams{
				PageSize: 2,
				Offset:   1,
			},
			configOrgName: "my-default-org",
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "my-default-org", *in.GetOrg()) &&
						assert.Equal(t, int32(2), *in.GetLimit()) &&
						assert.Equal(t, int32(1), *in.GetOffset())
				})).Return(api.Buckets{
					Buckets: &[]api.Bucket{
						{
							Id:    api.PtrString("222"),
							Name:  "my-bucket2",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 2400},
							},
						},
						{
							Id:    api.PtrString("333"),
							Name:  "my-bucket3",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
						},
					},
				}, nil)
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "my-default-org", *in.GetOrg()) &&
						assert.Equal(t, int32(2), *in.GetLimit()) &&
						assert.Equal(t, int32(3), *in.GetOffset())
				})).Return(api.Buckets{
					Buckets: &[]api.Bucket{
						{
							Id:    api.PtrString("444"),
							Name:  "my-bucket4",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 4800},
							},
						},
					},
				}, nil)
			},
			expectedStdoutPatterns: []string{
				`222\s+my-bucket2\s+40m0s\s+n/a\s+456`,
				`333\s+my-bucket3\s+1h0m0s\s+n/a\s+456`,
				`444\s+my-bucket4\s+1h20m0s\s+n/a\s+456`,
			},
		},
		{
			name: "override org by ID",
			params: bucket.BucketsListParams{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{OrgID: "456"},
				},
			},
			configOrgName: "my-default-org",
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "456", *in.GetOrgID()) &&
						assert.Nil(t, in.GetId()) &&
						assert.Nil(t, in.GetOrg()) &&
						assert.Nil(t, in.GetName())
				})).Return(api.Buckets{}, nil)
			},
		},
		{
			name: "override org by name",
			params: bucket.BucketsListParams{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{OrgName: "my-org"},
				},
				Limit: 2,
			},
			configOrgName: "my-default-org",
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetId()) &&
						assert.Nil(t, in.GetName()) &&
						assert.Nil(t, in.GetOrgID())
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
						{
							Id:    api.PtrString("999"),
							Name:  "bucket2",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 0, ShardGroupDurationSeconds: api.PtrInt64(60)},
							},
						},
					},
				}, nil)
			},
			expectedStdoutPatterns: []string{
				`123\s+my-bucket\s+1h0m0s\s+n/a\s+456`,
				`999\s+bucket2\s+infinite\s+1m0s\s+456`,
			},
		},
		{
			name: "list multiple bucket schema types",
			params: bucket.BucketsListParams{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{OrgName: "my-org"},
				},
				Limit: 3,
			},
			configOrgName: "my-default-org",
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().GetBuckets(gomock.Any()).Return(api.ApiGetBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetId()) &&
						assert.Nil(t, in.GetName()) &&
						assert.Nil(t, in.GetOrgID())
				})).Return(api.Buckets{
					Buckets: &[]api.Bucket{
						{
							Id:    api.PtrString("001"),
							Name:  "omit-schema-type",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
						},
						{
							Id:    api.PtrString("002"),
							Name:  "implicit-schema-type",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
							SchemaType: api.SCHEMATYPE_IMPLICIT.Ptr(),
						},
						{
							Id:    api.PtrString("003"),
							Name:  "explicit-schema-type",
							OrgID: api.PtrString("456"),
							RetentionRules: []api.RetentionRule{
								{EverySeconds: 3600},
							},
							SchemaType: api.SCHEMATYPE_EXPLICIT.Ptr(),
						},
					},
				}, nil)
			},
			expectedStdoutPatterns: []string{
				`001\s+omit-schema-type\s+1h0m0s\s+n/a\s+456\s+implicit`,
				`002\s+implicit-schema-type\s+1h0m0s\s+n/a\s+456\s+implicit`,
				`003\s+explicit-schema-type\s+1h0m0s\s+n/a\s+456\s+explicit`,
			},
		},
		{
			name:          "no org specified",
			expectedInErr: "must specify org ID or org name",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			client := mock.NewMockBucketsApi(ctrl)
			if tc.registerBucketExpectations != nil {
				tc.registerBucketExpectations(t, client)
			}
			stdio := mock.NewMockStdIO(ctrl)
			bytesWritten := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(bytesWritten.Write).AnyTimes()
			cli := bucket.Client{
				CLI:        clients.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio},
				BucketsApi: client,
			}

			err := cli.List(context.Background(), &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, bytesWritten.String())
				return
			}
			require.NoError(t, err)
			testutils.MatchLines(t, append(
				[]string{`ID\s+Name\s+Retention\s+Shard group duration\s+Organization ID\s+Schema Type`},
				tc.expectedStdoutPatterns...,
			), strings.Split(bytesWritten.String(), "\n"))
		})
	}
}
