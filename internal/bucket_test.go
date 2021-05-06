package internal_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBucketsCreate(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		configOrgName              string
		params                     internal.BucketsCreateParams
		registerOrgExpectations    func(*testing.T, *mock.MockOrganizationsApi)
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPattern      string
		expectedInErr              string
	}{
		{
			name: "minimal",
			params: internal.BucketsCreateParams{
				OrgID: "123",
				Name:  "my-bucket",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PostBuckets(gomock.Any()).Return(api.ApiPostBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().
					PostBucketsExecute(tmock.MatchedBy(func(in api.ApiPostBucketsRequest) bool {
						body := in.GetPostBucketRequest()
						return assert.NotNil(t, body) &&
							assert.Equal(t, "123", body.OrgID) &&
							assert.Equal(t, "my-bucket", body.Name) &&
							assert.Nil(t, body.Description) &&
							assert.Empty(t, body.RetentionRules)
					})).
					Return(api.Bucket{
						Id:             api.PtrString("456"),
						OrgID:          api.PtrString("123"),
						Name:           "my-bucket",
						RetentionRules: nil,
					}, nil)
			},
			expectedStdoutPattern: "456\\s+my-bucket\\s+infinite\\s+n/a\\s+123",
		},
		{
			name: "fully specified",
			params: internal.BucketsCreateParams{
				OrgID:              "123",
				Name:               "my-bucket",
				Description:        "my cool bucket",
				Retention:          "24h",
				ShardGroupDuration: "1h",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PostBuckets(gomock.Any()).Return(api.ApiPostBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().
					PostBucketsExecute(tmock.MatchedBy(func(in api.ApiPostBucketsRequest) bool {
						body := in.GetPostBucketRequest()
						return assert.NotNil(t, body) &&
							assert.Equal(t, "123", body.OrgID) &&
							assert.Equal(t, "my-bucket", body.Name) &&
							assert.Equal(t, "my cool bucket", *body.Description) &&
							assert.Len(t, body.RetentionRules, 1) &&
							assert.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds) &&
							assert.Equal(t, int64(3600), *body.RetentionRules[0].ShardGroupDurationSeconds)
					})).
					Return(api.Bucket{
						Id:    api.PtrString("456"),
						OrgID: api.PtrString("123"),
						Name:  "my-bucket",
						RetentionRules: []api.RetentionRule{
							{EverySeconds: 86400, ShardGroupDurationSeconds: api.PtrInt64(3600)},
						},
					}, nil)
			},
			expectedStdoutPattern: "456\\s+my-bucket\\s+24h0m0s\\s+1h0m0s\\s+123",
		},
		{
			name: "retention but not shard-group duration",
			params: internal.BucketsCreateParams{
				OrgID:     "123",
				Name:      "my-bucket",
				Retention: "24h",
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PostBuckets(gomock.Any()).Return(api.ApiPostBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().
					PostBucketsExecute(tmock.MatchedBy(func(in api.ApiPostBucketsRequest) bool {
						body := in.GetPostBucketRequest()
						return assert.NotNil(t, body) &&
							assert.Equal(t, "123", body.OrgID) &&
							assert.Equal(t, "my-bucket", body.Name) &&
							assert.Nil(t, body.Description) &&
							assert.Len(t, body.RetentionRules, 1) &&
							assert.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds) &&
							assert.Nil(t, body.RetentionRules[0].ShardGroupDurationSeconds)
					})).
					Return(api.Bucket{
						Id:             api.PtrString("456"),
						OrgID:          api.PtrString("123"),
						Name:           "my-bucket",
						RetentionRules: []api.RetentionRule{{EverySeconds: 86400}},
					}, nil)
			},
		},
		{
			name: "create bucket with explicit schema",
			params: internal.BucketsCreateParams{
				OrgID:      "123",
				Name:       "my-bucket",
				SchemaType: api.SCHEMATYPE_EXPLICIT,
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PostBuckets(gomock.Any()).Return(api.ApiPostBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().
					PostBucketsExecute(tmock.MatchedBy(func(in api.ApiPostBucketsRequest) bool {
						body := in.GetPostBucketRequest()
						return assert.NotNil(t, body) &&
							assert.Equal(t, "123", body.OrgID) &&
							assert.Equal(t, "my-bucket", body.Name) &&
							assert.Nil(t, body.Description) &&
							assert.Empty(t, body.RetentionRules) &&
							assert.Equal(t, api.SCHEMATYPE_EXPLICIT, *body.SchemaType)
					})).
					Return(api.Bucket{
						Id:         api.PtrString("456"),
						OrgID:      api.PtrString("123"),
						Name:       "my-bucket",
						SchemaType: api.SCHEMATYPE_EXPLICIT.Ptr(),
					}, nil)
			},
			expectedStdoutPattern: `456\s+my-bucket\s+infinite\s+n/a\s+123\s+explicit`,
		},
		{
			name: "look up org by name",
			params: internal.BucketsCreateParams{
				OrgName:            "my-org",
				Name:               "my-bucket",
				Description:        "my cool bucket",
				Retention:          "24h",
				ShardGroupDuration: "1h",
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).
					Return(api.Organizations{
						Orgs: &[]api.Organization{{Id: api.PtrString("123")}},
					}, nil)
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PostBuckets(gomock.Any()).Return(api.ApiPostBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().
					PostBucketsExecute(tmock.MatchedBy(func(in api.ApiPostBucketsRequest) bool {
						body := in.GetPostBucketRequest()
						return assert.NotNil(t, body) &&
							assert.Equal(t, "123", body.OrgID) &&
							assert.Equal(t, "my-bucket", body.Name) &&
							assert.Equal(t, "my cool bucket", *body.Description) &&
							assert.Len(t, body.RetentionRules, 1) &&
							assert.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds) &&
							assert.Equal(t, int64(3600), *body.RetentionRules[0].ShardGroupDurationSeconds)
					})).
					Return(api.Bucket{
						Id:    api.PtrString("456"),
						OrgID: api.PtrString("123"),
						Name:  "my-bucket",
						RetentionRules: []api.RetentionRule{
							{EverySeconds: 86400, ShardGroupDurationSeconds: api.PtrInt64(3600)},
						},
					}, nil)
			},
			expectedStdoutPattern: "456\\s+my-bucket\\s+24h0m0s\\s+1h0m0s\\s+123",
		},
		{
			name: "look up org by name from config",
			params: internal.BucketsCreateParams{
				Name:               "my-bucket",
				Description:        "my cool bucket",
				Retention:          "24h",
				ShardGroupDuration: "1h",
			},
			configOrgName: "my-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).
					Return(api.Organizations{
						Orgs: &[]api.Organization{{Id: api.PtrString("123")}},
					}, nil)
			},
			registerBucketExpectations: func(t *testing.T, bucketsApi *mock.MockBucketsApi) {
				bucketsApi.EXPECT().PostBuckets(gomock.Any()).Return(api.ApiPostBucketsRequest{ApiService: bucketsApi})
				bucketsApi.EXPECT().
					PostBucketsExecute(tmock.MatchedBy(func(in api.ApiPostBucketsRequest) bool {
						body := in.GetPostBucketRequest()
						return assert.NotNil(t, body) &&
							assert.Equal(t, "123", body.OrgID) &&
							assert.Equal(t, "my-bucket", body.Name) &&
							assert.Equal(t, "my cool bucket", *body.Description) &&
							assert.Len(t, body.RetentionRules, 1) &&
							assert.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds) &&
							assert.Equal(t, int64(3600), *body.RetentionRules[0].ShardGroupDurationSeconds)
					})).
					Return(api.Bucket{
						Id:    api.PtrString("456"),
						OrgID: api.PtrString("123"),
						Name:  "my-bucket",
						RetentionRules: []api.RetentionRule{
							{EverySeconds: 86400, ShardGroupDurationSeconds: api.PtrInt64(3600)},
						},
					}, nil)
			},
			expectedStdoutPattern: "456\\s+my-bucket\\s+24h0m0s\\s+1h0m0s\\s+123",
		},
		{
			name: "no org specified",
			params: internal.BucketsCreateParams{
				Name:               "my-bucket",
				Description:        "my cool bucket",
				Retention:          "24h",
				ShardGroupDuration: "1h",
			},
			expectedInErr: "must specify org ID or org name",
		},
		{
			name: "no such org",
			params: internal.BucketsCreateParams{
				Name:               "my-bucket",
				OrgName:            "fake-org",
				Description:        "my cool bucket",
				Retention:          "24h",
				ShardGroupDuration: "1h",
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(gomock.Any()).Return(api.Organizations{}, nil)
			},
			expectedInErr: "no organization found",
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

			orgApi := mock.NewMockOrganizationsApi(ctrl)
			if tc.registerOrgExpectations != nil {
				tc.registerOrgExpectations(t, orgApi)
			}
			bucketApi := mock.NewMockBucketsApi(ctrl)
			if tc.registerBucketExpectations != nil {
				tc.registerBucketExpectations(t, bucketApi)
			}
			cli := internal.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio}
			clients := internal.BucketsClients{
				OrgApi:    orgApi,
				BucketApi: bucketApi,
			}

			err := cli.BucketsCreate(context.Background(), &clients, &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, writtenBytes.String())
				return
			}
			require.NoError(t, err)
			outLines := strings.Split(writtenBytes.String(), "\n")
			if outLines[len(outLines)-1] == "" {
				outLines = outLines[:len(outLines)-1]
			}
			require.Equal(t, 2, len(outLines))
			require.Regexp(t, "ID\\s+Name\\s+Retention\\s+Shard group duration\\s+Organization ID", outLines[0])
			require.Regexp(t, tc.expectedStdoutPattern, outLines[1])
		})
	}
}

func TestBucketsList(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		configOrgName              string
		params                     internal.BucketsListParams
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPatterns     []string
		expectedInErr              string
	}{
		{
			name: "by ID",
			params: internal.BucketsListParams{
				ID: "123",
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
				"123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
			},
		},
		{
			name: "by name",
			params: internal.BucketsListParams{
				Name: "my-bucket",
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
				"123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
			},
		},
		{
			name: "override org by ID",
			params: internal.BucketsListParams{
				OrgID: "456",
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
			params: internal.BucketsListParams{
				OrgName: "my-org",
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
				"123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
				"999\\s+bucket2\\s+infinite\\s+1m0s\\s+456",
			},
		},
		{
			name: "list multiple bucket schema types",
			params: internal.BucketsListParams{
				OrgName: "my-org",
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
			stdio := mock.NewMockStdIO(ctrl)
			bytesWritten := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(bytesWritten.Write).AnyTimes()
			cli := internal.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio}

			client := mock.NewMockBucketsApi(ctrl)
			if tc.registerBucketExpectations != nil {
				tc.registerBucketExpectations(t, client)
			}

			err := cli.BucketsList(context.Background(), client, &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, bytesWritten.String())
				return
			}
			require.NoError(t, err)
			outLines := strings.Split(bytesWritten.String(), "\n")
			if outLines[len(outLines)-1] == "" {
				outLines = outLines[:len(outLines)-1]
			}
			require.Equal(t, len(tc.expectedStdoutPatterns)+1, len(outLines))
			require.Regexp(t, "ID\\s+Name\\s+Retention\\s+Shard group duration\\s+Organization ID", outLines[0])
			for i, pattern := range tc.expectedStdoutPatterns {
				require.Regexp(t, pattern, outLines[i+1])
			}
		})
	}
}

func TestBucketsUpdate(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		params                     internal.BucketsUpdateParams
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPattern      string
	}{
		{
			name: "name",
			params: internal.BucketsUpdateParams{
				ID:   "123",
				Name: "cold-storage",
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
			expectedStdoutPattern: "123\\s+cold-storage\\s+infinite\\s+n/a\\s+456",
		},
		{
			name: "description",
			params: internal.BucketsUpdateParams{
				ID:          "123",
				Description: "a very useful description",
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
			expectedStdoutPattern: "123\\s+my-bucket\\s+infinite\\s+n/a\\s+456",
		},
		{
			name: "retention",
			params: internal.BucketsUpdateParams{
				ID:        "123",
				Retention: "3w",
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
			expectedStdoutPattern: "123\\s+my-bucket\\s+504h0m0s\\s+n/a\\s+456",
		},
		{
			name: "shard-group duration",
			params: internal.BucketsUpdateParams{
				ID:                 "123",
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
			expectedStdoutPattern: "123\\s+my-bucket\\s+infinite\\s+10h30m0s\\s+456",
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
			cli := internal.CLI{StdIO: stdio}
			client := mock.NewMockBucketsApi(ctrl)
			if tc.registerBucketExpectations != nil {
				tc.registerBucketExpectations(t, client)
			}

			err := cli.BucketsUpdate(context.Background(), client, &tc.params)
			require.NoError(t, err)
			outLines := strings.Split(writtenBytes.String(), "\n")
			if outLines[len(outLines)-1] == "" {
				outLines = outLines[:len(outLines)-1]
			}
			require.Equal(t, 2, len(outLines))
			require.Regexp(t, "ID\\s+Name\\s+Retention\\s+Shard group duration\\s+Organization ID", outLines[0])
			require.Regexp(t, tc.expectedStdoutPattern, outLines[1])
		})
	}
}

func TestBucketsDelete(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		configOrgName              string
		params                     internal.BucketsDeleteParams
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPattern      string
		expectedInErr              string
	}{
		{
			name:          "by ID",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
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
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456\\s+implicit",
		},
		{
			name:          "by name and org ID",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
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
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456\\s+implicit",
		},
		{
			name:          "by name and org name",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
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
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456\\s+implicit",
		},
		{
			name:          "by name in default org",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
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
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456\\s+implicit",
		},
		{
			name: "by name without org",
			params: internal.BucketsDeleteParams{
				Name: "my-bucket",
			},
			expectedInErr: "must specify org ID or org name",
		},
		{
			name: "no such bucket",
			params: internal.BucketsDeleteParams{
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
			cli := internal.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio}
			client := mock.NewMockBucketsApi(ctrl)
			if tc.registerBucketExpectations != nil {
				tc.registerBucketExpectations(t, client)
			}

			err := cli.BucketsDelete(context.Background(), client, &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, writtenBytes.String())
				return
			}
			require.NoError(t, err)
			outLines := strings.Split(writtenBytes.String(), "\n")
			if outLines[len(outLines)-1] == "" {
				outLines = outLines[:len(outLines)-1]
			}
			require.Regexp(t, `ID\s+Name\s+Retention\s+Shard group duration\s+Organization ID\s+Schema Type\s+Deleted`, outLines[0])
			require.Regexp(t, tc.expectedStdoutPattern+"\\s+true", outLines[1])
		})
	}
}
