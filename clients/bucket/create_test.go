package bucket_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/bucket"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBucketsCreate(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                       string
		configOrgName              string
		params                     bucket.BucketsCreateParams
		registerOrgExpectations    func(*testing.T, *mock.MockOrganizationsApi)
		registerBucketExpectations func(*testing.T, *mock.MockBucketsApi)
		expectedStdoutPattern      string
		expectedInErr              string
	}{
		{
			name: "minimal",
			params: bucket.BucketsCreateParams{
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
			expectedStdoutPattern: `456\s+my-bucket\s+infinite\s+n/a\s+123`,
		},
		{
			name: "fully specified",
			params: bucket.BucketsCreateParams{
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
			expectedStdoutPattern: `456\s+my-bucket\s+24h0m0s\s+1h0m0s\s+123`,
		},
		{
			name: "retention but not shard-group duration",
			params: bucket.BucketsCreateParams{
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
			params: bucket.BucketsCreateParams{
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
			params: bucket.BucketsCreateParams{
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
			expectedStdoutPattern: `456\s+my-bucket\s+24h0m0s\s+1h0m0s\s+123`,
		},
		{
			name: "look up org by name from config",
			params: bucket.BucketsCreateParams{
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
			expectedStdoutPattern: `456\s+my-bucket\s+24h0m0s\s+1h0m0s\s+123`,
		},
		{
			name: "no org specified",
			params: bucket.BucketsCreateParams{
				Name:               "my-bucket",
				Description:        "my cool bucket",
				Retention:          "24h",
				ShardGroupDuration: "1h",
			},
			expectedInErr: "must specify org ID or org name",
		},
		{
			name: "no such org",
			params: bucket.BucketsCreateParams{
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
			cli := bucket.Client{
				CLI:              clients.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio},
				OrganizationsApi: orgApi,
				BucketsApi:       bucketApi,
			}

			err := cli.Create(context.Background(), &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, writtenBytes.String())
				return
			}
			require.NoError(t, err)
			testutils.MatchLines(t, []string{
				`ID\s+Name\s+Retention\s+Shard group duration\s+Organization ID\s+Schema Type`,
				tc.expectedStdoutPattern,
			}, strings.Split(writtenBytes.String(), "\n"))
		})
	}
}
