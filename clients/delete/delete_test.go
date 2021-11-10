package delete_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/delete"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	id1 := "1111111111111111"
	id2 := "2222222222222222"

	start, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z")
	stop, _ := time.Parse(time.RFC3339Nano, "2021-01-01T00:00:00Z")

	testCases := []struct {
		name                 string
		params               delete.Params
		defaultOrgName       string
		registerExpectations func(*testing.T, *mock.MockDeleteApi)
		expectedErr          string
	}{
		{
			name: "by IDs",
			params: delete.Params{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{
						OrgID: id1,
					},
					BucketParams: clients.BucketParams{
						BucketID: id2,
					},
				},
				Start: start.Format(time.RFC3339Nano),
				Stop:  stop.Format(time.RFC3339Nano),
			},
			defaultOrgName: "my-default-org",
			registerExpectations: func(t *testing.T, delApi *mock.MockDeleteApi) {
				delApi.EXPECT().PostDelete(gomock.Any()).Return(api.ApiPostDeleteRequest{ApiService: delApi})
				delApi.EXPECT().PostDeleteExecute(tmock.MatchedBy(func(in api.ApiPostDeleteRequest) bool {
					body := in.GetDeletePredicateRequest()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1, *in.GetOrgID()) &&
						assert.Nil(t, in.GetOrg()) &&
						assert.Equal(t, id2, *in.GetBucketID()) &&
						assert.Nil(t, in.GetBucket()) &&
						assert.Equal(t, start, body.GetStart()) &&
						assert.Equal(t, stop, body.GetStop()) &&
						assert.Nil(t, body.Predicate)
				})).Return(nil)
			},
		},
		{
			name: "by org ID, bucket name",
			params: delete.Params{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{
						OrgID: id1,
					},
					BucketParams: clients.BucketParams{
						BucketName: "my-bucket",
					},
				},
				Start: start.Format(time.RFC3339Nano),
				Stop:  stop.Format(time.RFC3339Nano),
			},
			defaultOrgName: "my-default-org",
			registerExpectations: func(t *testing.T, delApi *mock.MockDeleteApi) {
				delApi.EXPECT().PostDelete(gomock.Any()).Return(api.ApiPostDeleteRequest{ApiService: delApi})
				delApi.EXPECT().PostDeleteExecute(tmock.MatchedBy(func(in api.ApiPostDeleteRequest) bool {
					body := in.GetDeletePredicateRequest()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1, *in.GetOrgID()) &&
						assert.Nil(t, in.GetOrg()) &&
						assert.Equal(t, "my-bucket", *in.GetBucket()) &&
						assert.Nil(t, in.GetBucketID()) &&
						assert.Equal(t, start, body.GetStart()) &&
						assert.Equal(t, stop, body.GetStop()) &&
						assert.Nil(t, body.Predicate)
				})).Return(nil)
			},
		},
		{
			name: "by org name, bucket ID",
			params: delete.Params{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{
						OrgName: "my-org",
					},
					BucketParams: clients.BucketParams{
						BucketID: id2,
					},
				},
				Start:     start.Format(time.RFC3339Nano),
				Stop:      stop.Format(time.RFC3339Nano),
				Predicate: `foo = "bar"`,
			},
			defaultOrgName: "my-default-org",
			registerExpectations: func(t *testing.T, delApi *mock.MockDeleteApi) {
				delApi.EXPECT().PostDelete(gomock.Any()).Return(api.ApiPostDeleteRequest{ApiService: delApi})
				delApi.EXPECT().PostDeleteExecute(tmock.MatchedBy(func(in api.ApiPostDeleteRequest) bool {
					body := in.GetDeletePredicateRequest()
					return assert.NotNil(t, body) &&
						assert.Equal(t, "my-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetOrgID()) &&
						assert.Equal(t, id2, *in.GetBucketID()) &&
						assert.Nil(t, in.GetBucket()) &&
						assert.Equal(t, start, body.GetStart()) &&
						assert.Equal(t, stop, body.GetStop()) &&
						assert.Equal(t, `foo = "bar"`, body.GetPredicate())
				})).Return(nil)
			},
		},
		{
			name: "by names",
			params: delete.Params{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{},
					BucketParams: clients.BucketParams{
						BucketName: "my-bucket",
					},
				},
				Start:     start.Format(time.RFC3339Nano),
				Stop:      stop.Format(time.RFC3339Nano),
				Predicate: `foo = "bar"`,
			},
			defaultOrgName: "my-default-org",
			registerExpectations: func(t *testing.T, delApi *mock.MockDeleteApi) {
				delApi.EXPECT().PostDelete(gomock.Any()).Return(api.ApiPostDeleteRequest{ApiService: delApi})
				delApi.EXPECT().PostDeleteExecute(tmock.MatchedBy(func(in api.ApiPostDeleteRequest) bool {
					body := in.GetDeletePredicateRequest()
					return assert.NotNil(t, body) &&
						assert.Equal(t, "my-default-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetOrgID()) &&
						assert.Equal(t, "my-bucket", *in.GetBucket()) &&
						assert.Nil(t, in.GetBucketID()) &&
						assert.Equal(t, start, body.GetStart()) &&
						assert.Equal(t, stop, body.GetStop()) &&
						assert.Equal(t, `foo = "bar"`, body.GetPredicate())
				})).Return(nil)
			},
		},
		{
			name:        "no org",
			expectedErr: clients.ErrMustSpecifyOrg.Error(),
		},
		{
			name:           "no bucket",
			defaultOrgName: "my-default-org",
			expectedErr:    clients.ErrMustSpecifyBucket.Error(),
		},
		{
			name: "bad start",
			params: delete.Params{
				OrgBucketParams: clients.OrgBucketParams{
					BucketParams: clients.BucketParams{
						BucketName: "my-bucket",
					},
				},
				Start:     "the beginning",
				Stop:      stop.Format(time.RFC3339Nano),
				Predicate: `foo = "bar"`,
			},
			defaultOrgName: "my-default-org",
			expectedErr:    `"the beginning" cannot be parsed`,
		},
		{
			name: "bad stop",
			params: delete.Params{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams: clients.OrgParams{},
					BucketParams: clients.BucketParams{
						BucketName: "my-bucket",
					},
				},
				Start:     start.Format(time.RFC3339Nano),
				Stop:      "the end",
				Predicate: `foo = "bar"`,
			},
			defaultOrgName: "my-default-org",
			expectedErr:    `"the end" cannot be parsed`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			api := mock.NewMockDeleteApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, api)
			}

			client := delete.Client{
				CLI:       clients.CLI{ActiveConfig: config.Config{Org: tc.defaultOrgName}},
				DeleteApi: api,
			}
			err := client.Delete(context.Background(), &tc.params)

			if tc.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
				return
			}

			require.NoError(t, err)
		})
	}
}
