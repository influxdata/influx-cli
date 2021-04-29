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
}

func TestBucketsUpdate(t *testing.T) {
	t.Parallel()
}

func TestBucketsDelete(t *testing.T) {
	t.Parallel()
}
