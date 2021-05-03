package internal_test

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBucketsCreate(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name                  string
		configOrgName         string
		params                internal.BucketsCreateParams
		buildOrgLookupFn      func(*testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error)
		buildBucketCreateFn   func(*testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error)
		expectedStdoutPattern string
		expectedInErr         string
	}{
		{
			name: "minimal",
			params: internal.BucketsCreateParams{
				OrgID: "123",
				Name:  "my-bucket",
			},
			buildOrgLookupFn: func(t *testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
				return func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
					return api.Organizations{}, nil, errors.New("unexpected org lookup call")
				}
			},
			buildBucketCreateFn: func(t *testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
					body := req.GetPostBucketRequest()
					require.NotNil(t, body)
					require.Equal(t, "123", body.OrgID)
					require.Equal(t, "my-bucket", body.Name)
					require.Empty(t, body.RetentionRules)

					return api.Bucket{
						Id:             api.PtrString("456"),
						OrgID:          api.PtrString("123"),
						Name:           "my-bucket",
						RetentionRules: nil,
					}, nil, nil
				}
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
			buildOrgLookupFn: func(t *testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
				return func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
					return api.Organizations{}, nil, errors.New("unexpected org lookup call")
				}
			},
			buildBucketCreateFn: func(t *testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
					body := req.GetPostBucketRequest()
					require.NotNil(t, body)
					require.Equal(t, "123", body.OrgID)
					require.Equal(t, "my-bucket", body.Name)
					require.Equal(t, "my cool bucket", *body.Description)
					require.Equal(t, 1, len(body.RetentionRules))
					require.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds)
					require.Equal(t, int64(3600), *body.RetentionRules[0].ShardGroupDurationSeconds)

					return api.Bucket{
						Id:             api.PtrString("456"),
						OrgID:          api.PtrString("123"),
						Name:           "my-bucket",
						RetentionRules: body.RetentionRules,
					}, nil, nil
				}
			},
			expectedStdoutPattern: "456\\s+my-bucket\\s+24h0m0s\\s+1h0m0s\\s+123",
		},
		{
			name: "retention but not shard-group duration",
			params: internal.BucketsCreateParams{
				OrgID:       "123",
				Name:        "my-bucket",
				Retention:   "24h",
			},
			buildOrgLookupFn: func(t *testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
				return func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
					return api.Organizations{}, nil, errors.New("unexpected org lookup call")
				}
			},
			buildBucketCreateFn: func(t *testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
					body := req.GetPostBucketRequest()
					require.NotNil(t, body)
					require.Equal(t, "123", body.OrgID)
					require.Equal(t, "my-bucket", body.Name)
					require.Nil(t, body.Description)
					require.Equal(t, 1, len(body.RetentionRules))
					require.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds)
					require.Nil(t, body.RetentionRules[0].ShardGroupDurationSeconds)

					return api.Bucket{
						Id:             api.PtrString("456"),
						OrgID:          api.PtrString("123"),
						Name:           "my-bucket",
						RetentionRules: body.RetentionRules,
					}, nil, nil
				}
			},
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
			buildOrgLookupFn: func(t *testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
				return func(req api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
					require.Equal(t, "my-org", *req.GetOrg())
					return api.Organizations{
						Orgs: &[]api.Organization{{Id: api.PtrString("123")}},
					}, nil, nil
				}
			},
			buildBucketCreateFn: func(t *testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
					body := req.GetPostBucketRequest()
					require.NotNil(t, body)
					require.Equal(t, "123", body.OrgID)
					require.Equal(t, "my-bucket", body.Name)
					require.Equal(t, "my cool bucket", *body.Description)
					require.Equal(t, 1, len(body.RetentionRules))
					require.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds)
					require.Equal(t, int64(3600), *body.RetentionRules[0].ShardGroupDurationSeconds)

					return api.Bucket{
						Id:             api.PtrString("456"),
						OrgID:          api.PtrString("123"),
						Name:           "my-bucket",
						RetentionRules: body.RetentionRules,
					}, nil, nil
				}
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
			buildOrgLookupFn: func(t *testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
				return func(req api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
					require.Equal(t, "my-org", *req.GetOrg())
					return api.Organizations{
						Orgs: &[]api.Organization{{Id: api.PtrString("123")}},
					}, nil, nil
				}
			},
			buildBucketCreateFn: func(t *testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
					body := req.GetPostBucketRequest()
					require.NotNil(t, body)
					require.Equal(t, "123", body.OrgID)
					require.Equal(t, "my-bucket", body.Name)
					require.Equal(t, "my cool bucket", *body.Description)
					require.Equal(t, 1, len(body.RetentionRules))
					require.Equal(t, int64(86400), body.RetentionRules[0].EverySeconds)
					require.Equal(t, int64(3600), *body.RetentionRules[0].ShardGroupDurationSeconds)

					return api.Bucket{
						Id:             api.PtrString("456"),
						OrgID:          api.PtrString("123"),
						Name:           "my-bucket",
						RetentionRules: body.RetentionRules,
					}, nil, nil
				}
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
			buildOrgLookupFn: func(t *testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
				return func(req api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
					return api.Organizations{}, nil, errors.New("shouldn't be called")
				}
			},
			buildBucketCreateFn: func(t *testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
					return api.Bucket{}, nil, errors.New("shouldn't be called")
				}
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
			buildOrgLookupFn: func(t *testing.T) func(api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
				return func(req api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
					return api.Organizations{}, nil, nil
				}
			},
			buildBucketCreateFn: func(t *testing.T) func(api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPostBucketsRequest) (api.Bucket, *http.Response, error) {
					return api.Bucket{}, nil, errors.New("shouldn't be called")
				}
			},
			expectedInErr: "no organization found",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stdio := mock.NewMockStdio(nil, false)
			cli := internal.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio}
			clients := internal.BucketsClients{
				OrgApi: &mock.OrganizationsApi{
					GetOrgsExecuteFn: tc.buildOrgLookupFn(t),
				},
				BucketApi: &mock.BucketsApi{
					PostBucketsExecuteFn: tc.buildBucketCreateFn(t),
				},
			}

			err := cli.BucketsCreate(context.Background(), &clients, &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, stdio.Stdout())
				return
			}
			require.NoError(t, err)
			outLines := strings.Split(stdio.Stdout(), "\n")
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
		name                   string
		configOrgName          string
		params                 internal.BucketsListParams
		buildBucketLookupFn    func(*testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error)
		expectedStdoutPatterns []string
		expectedInErr          string
	}{
		{
			name: "by ID",
			params: internal.BucketsListParams{
				ID: "123",
			},
			configOrgName: "my-default-org",
			buildBucketLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "123", *req.GetId())
					require.Equal(t, "my-default-org", *req.GetOrg())
					require.Nil(t, req.GetName())
					require.Nil(t, req.GetOrgID())
					return api.Buckets{
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
					}, nil, nil
				}
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
			buildBucketLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "my-bucket", *req.GetName())
					require.Equal(t, "my-default-org", *req.GetOrg())
					require.Nil(t, req.GetId())
					require.Nil(t, req.GetOrgID())
					return api.Buckets{
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
					}, nil, nil
				}
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
			buildBucketLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "456", *req.GetOrgID())
					require.Nil(t, req.GetId())
					require.Nil(t, req.GetOrg())
					require.Nil(t, req.GetOrg())
					return api.Buckets{}, nil, nil
				}
			},
		},
		{
			name: "override org by name",
			params: internal.BucketsListParams{
				OrgName: "my-org",
			},
			configOrgName: "my-default-org",
			buildBucketLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "my-org", *req.GetOrg())
					require.Nil(t, req.GetId())
					require.Nil(t, req.GetName())
					require.Nil(t, req.GetOrgID())
					return api.Buckets{
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
					}, nil, nil
				}
			},
			expectedStdoutPatterns: []string{
				"123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
				"999\\s+bucket2\\s+infinite\\s+1m0s\\s+456",
			},
		},
		{
			name:          "no org specified",
			expectedInErr: "must specify org ID or org name",
			buildBucketLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					return api.Buckets{}, nil, errors.New("shouldn't be called")
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stdio := mock.NewMockStdio(nil, false)
			cli := internal.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio}
			client := mock.BucketsApi{
				GetBucketsExecuteFn: tc.buildBucketLookupFn(t),
			}

			err := cli.BucketsList(context.Background(), &client, &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, stdio.Stdout())
				return
			}
			require.NoError(t, err)
			outLines := strings.Split(stdio.Stdout(), "\n")
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
		name                  string
		params                internal.BucketsUpdateParams
		buildBucketUpdateFn   func(*testing.T) func(api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error)
		expectedStdoutPattern string
	}{
		{
			name: "name",
			params: internal.BucketsUpdateParams{
				ID:   "123",
				Name: "cold-storage",
			},
			buildBucketUpdateFn: func(t *testing.T) func(api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
					require.Equal(t, "123", req.GetBucketID())
					body := req.GetPatchBucketRequest()
					require.Equal(t, "cold-storage", body.GetName())
					require.Empty(t, body.GetDescription())
					require.Empty(t, body.GetRetentionRules())

					return api.Bucket{
						Id:    api.PtrString("123"),
						Name:  "cold-storage",
						OrgID: api.PtrString("456"),
					}, nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+cold-storage\\s+infinite\\s+n/a\\s+456",
		},
		{
			name: "description",
			params: internal.BucketsUpdateParams{
				ID:          "123",
				Description: "a very useful description",
			},
			buildBucketUpdateFn: func(t *testing.T) func(api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
					require.Equal(t, "123", req.GetBucketID())
					body := req.GetPatchBucketRequest()
					require.Equal(t, "a very useful description", body.GetDescription())
					require.Empty(t, body.GetName())
					require.Empty(t, body.GetRetentionRules())

					return api.Bucket{
						Id:          api.PtrString("123"),
						Name:        "my-bucket",
						Description: api.PtrString("a very useful description"),
						OrgID:       api.PtrString("456"),
					}, nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+my-bucket\\s+infinite\\s+n/a\\s+456",
		},
		{
			name: "retention",
			params: internal.BucketsUpdateParams{
				ID:        "123",
				Retention: "3w",
			},
			buildBucketUpdateFn: func(t *testing.T) func(api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
					require.Equal(t, "123", req.GetBucketID())
					body := req.GetPatchBucketRequest()
					require.Len(t, body.GetRetentionRules(), 1)
					rule := body.GetRetentionRules()[0]
					require.Nil(t, rule.ShardGroupDurationSeconds)
					require.Equal(t, int64(3*7*24*3600), rule.GetEverySeconds())
					require.Nil(t, body.Name)
					require.Nil(t, body.Description)

					return api.Bucket{
						Id:    api.PtrString("123"),
						Name:  "my-bucket",
						OrgID: api.PtrString("456"),
						RetentionRules: []api.RetentionRule{
							{EverySeconds: rule.GetEverySeconds()},
						},
					}, nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+my-bucket\\s+504h0m0s\\s+n/a\\s+456",
		},
		{
			name: "shard-group duration",
			params: internal.BucketsUpdateParams{
				ID:                 "123",
				ShardGroupDuration: "10h30m",
			},
			buildBucketUpdateFn: func(t *testing.T) func(api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
				return func(req api.ApiPatchBucketsIDRequest) (api.Bucket, *http.Response, error) {
					require.Equal(t, "123", req.GetBucketID())
					body := req.GetPatchBucketRequest()
					require.Len(t, body.GetRetentionRules(), 1)
					rule := body.GetRetentionRules()[0]
					require.Nil(t, rule.EverySeconds)
					require.Equal(t, int64(10*3600+30*60), rule.GetShardGroupDurationSeconds())
					require.Nil(t, body.Name)
					require.Nil(t, body.Description)

					return api.Bucket{
						Id:    api.PtrString("123"),
						Name:  "my-bucket",
						OrgID: api.PtrString("456"),
						RetentionRules: []api.RetentionRule{
							{ShardGroupDurationSeconds: rule.ShardGroupDurationSeconds},
						},
					}, nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+my-bucket\\s+infinite\\s+10h30m0s\\s+456",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stdio := mock.NewMockStdio(nil, false)
			cli := internal.CLI{StdIO: stdio}
			client := mock.BucketsApi{
				PatchBucketsIDExecuteFn: tc.buildBucketUpdateFn(t),
			}

			err := cli.BucketsUpdate(context.Background(), &client, &tc.params)
			require.NoError(t, err)
			outLines := strings.Split(stdio.Stdout(), "\n")
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
		name                  string
		configOrgName         string
		params                internal.BucketsDeleteParams
		buildBucketsLookupFn  func(*testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error)
		buildBucketDeleteFn   func(*testing.T) func(api.ApiDeleteBucketsIDRequest) (*http.Response, error)
		expectedStdoutPattern string
		expectedInErr         string
	}{
		{
			name:          "by ID",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
				ID: "123",
			},
			buildBucketsLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "123", *req.GetId())
					require.Nil(t, req.GetName())
					require.Nil(t, req.GetOrgID())
					require.Nil(t, req.GetOrg())

					return api.Buckets{
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
					}, nil, nil
				}
			},
			buildBucketDeleteFn: func(t *testing.T) func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
				return func(req api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
					assert.Equal(t, "123", req.GetBucketID())
					return nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
		},
		{
			name:          "by name and org ID",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
				Name:  "my-bucket",
				OrgID: "456",
			},
			buildBucketsLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "my-bucket", *req.GetName())
					require.Equal(t, "456", *req.GetOrgID())
					require.Nil(t, req.GetId())
					require.Nil(t, req.GetOrg())

					return api.Buckets{
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
					}, nil, nil
				}
			},
			buildBucketDeleteFn: func(t *testing.T) func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
				return func(req api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
					assert.Equal(t, "123", req.GetBucketID())
					return nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
		},
		{
			name:          "by name and org name",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
				Name:    "my-bucket",
				OrgName: "my-org",
			},
			buildBucketsLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "my-bucket", *req.GetName())
					require.Equal(t, "my-org", *req.GetOrg())
					require.Nil(t, req.GetId())
					require.Nil(t, req.GetOrgID())

					return api.Buckets{
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
					}, nil, nil
				}
			},
			buildBucketDeleteFn: func(t *testing.T) func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
				return func(req api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
					assert.Equal(t, "123", req.GetBucketID())
					return nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
		},
		{
			name:          "by name in default org",
			configOrgName: "my-default-org",
			params: internal.BucketsDeleteParams{
				Name: "my-bucket",
			},
			buildBucketsLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(req api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					require.Equal(t, "my-bucket", *req.GetName())
					require.Equal(t, "my-default-org", *req.GetOrg())
					require.Nil(t, req.GetId())
					require.Nil(t, req.GetOrgID())

					return api.Buckets{
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
					}, nil, nil
				}
			},
			buildBucketDeleteFn: func(t *testing.T) func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
				return func(req api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
					assert.Equal(t, "123", req.GetBucketID())
					return nil, nil
				}
			},
			expectedStdoutPattern: "123\\s+my-bucket\\s+1h0m0s\\s+n/a\\s+456",
		},
		{
			name: "by name without org",
			params: internal.BucketsDeleteParams{
				Name: "my-bucket",
			},
			buildBucketsLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					return api.Buckets{}, nil, errors.New("shouldn't be called")
				}
			},
			buildBucketDeleteFn: func(t *testing.T) func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
				return func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
					return nil, errors.New("shouldn't be called")
				}
			},
			expectedInErr: "must specify org ID or org name",
		},
		{
			name: "no such bucket",
			params: internal.BucketsDeleteParams{
				ID: "123",
			},
			buildBucketsLookupFn: func(t *testing.T) func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
				return func(api.ApiGetBucketsRequest) (api.Buckets, *http.Response, error) {
					return api.Buckets{}, nil, nil
				}
			},
			buildBucketDeleteFn: func(t *testing.T) func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
				return func(api.ApiDeleteBucketsIDRequest) (*http.Response, error) {
					return nil, errors.New("shouldn't be called")
				}
			},
			expectedInErr: "not found",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stdio := mock.NewMockStdio(nil, false)
			cli := internal.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio}
			client := mock.BucketsApi{
				GetBucketsExecuteFn:      tc.buildBucketsLookupFn(t),
				DeleteBucketsIDExecuteFn: tc.buildBucketDeleteFn(t),
			}

			err := cli.BucketsDelete(context.Background(), &client, &tc.params)
			if tc.expectedInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedInErr)
				require.Empty(t, stdio.Stdout())
				return
			}
			require.NoError(t, err)
			outLines := strings.Split(stdio.Stdout(), "\n")
			if outLines[len(outLines)-1] == "" {
				outLines = outLines[:len(outLines)-1]
			}
			require.Regexp(t, "ID\\s+Name\\s+Retention\\s+Shard group duration\\s+Organization ID\\s+Deleted", outLines[0])
			require.Regexp(t, tc.expectedStdoutPattern+"\\s+true", outLines[1])
		})
	}
}
