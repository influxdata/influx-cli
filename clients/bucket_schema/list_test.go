package bucket_schema_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/bucket_schema"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClient_List(t *testing.T) {
	t.Parallel()

	var (
		orgID         = influxid.MustIDFromString("deadf00dbaadf00d")
		bucketID      = influxid.MustIDFromString("f00ddeadf00dbaad")
		measurementID = influxid.MustIDFromString("1010f00ddeedfeed")
		createdAt     = time.Date(2004, 4, 9, 2, 15, 0, 0, time.UTC)
		updatedAt     = time.Date(2009, 9, 1, 2, 15, 0, 0, time.UTC)
	)

	type setupArgs struct {
		buckets *mock.MockBucketsApi
		schemas *mock.MockBucketSchemasApi
		cli     clients.CLI
		params  bucket_schema.ListParams
		cols    []api.MeasurementSchemaColumn
		stdio   *mock.MockStdIO
	}

	type optFn func(t *testing.T, a *setupArgs)

	type args struct {
		OrgName        string
		OrgID          influxid.ID
		BucketName     string
		BucketID       influxid.ID
		Name           string
		ExtendedOutput bool
	}

	withArgs := func(args args) optFn {
		return func(t *testing.T, a *setupArgs) {
			t.Helper()
			a.params = bucket_schema.ListParams{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams:    clients.OrgParams{OrgName: args.OrgName, OrgID: args.OrgID},
					BucketParams: clients.BucketParams{BucketName: args.BucketName, BucketID: args.BucketID},
				},
				Name:           args.Name,
				ExtendedOutput: args.ExtendedOutput,
			}
		}
	}

	expGetBuckets := func(n ...string) optFn {
		return func(t *testing.T, a *setupArgs) {
			t.Helper()
			require.True(t, len(n) <= 1, "either zero or one bucket name")
			var buckets []api.Bucket
			if len(n) == 1 {
				bucket := api.NewBucket(n[0], nil)
				bucket.SetOrgID(orgID.String())
				bucket.SetId(bucketID.String())
				bucket.SetName(n[0])
				buckets = []api.Bucket{*bucket}
			}

			req := api.ApiGetBucketsRequest{ApiService: a.buckets}

			a.buckets.EXPECT().
				GetBuckets(gomock.Any()).
				Return(req)

			a.buckets.EXPECT().
				GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					matchOrg := (in.GetOrg() != nil && *in.GetOrg() == a.params.OrgName) ||
						(in.GetOrgID() != nil && a.params.OrgID.Valid() && *in.GetOrgID() == a.params.OrgID.String())
					matchBucket := (in.GetName() != nil && *in.GetName() == a.params.BucketName) ||
						(in.GetId() != nil && a.params.BucketID.Valid() && *in.GetId() == a.params.BucketID.String())
					return matchOrg && matchBucket
				})).
				Return(api.Buckets{Buckets: &buckets}, nil)
		}
	}

	withCols := func(p string) optFn {
		return func(t *testing.T, a *setupArgs) {
			data, err := os.ReadFile(filepath.Join("testdata", p))
			require.NoError(t, err)

			var f bucket_schema.ColumnsFormat
			decoder, err := f.DecoderFn(p)
			require.NoError(t, err)
			cols, err := decoder(bytes.NewReader(data))
			require.NoError(t, err)
			a.cols = cols
		}
	}

	expGetMeasurementSchemas := func() optFn {
		return func(t *testing.T, a *setupArgs) {
			t.Helper()

			req := api.ApiGetMeasurementSchemasRequest{ApiService: a.schemas}.BucketID(bucketID.String())

			a.schemas.EXPECT().
				GetMeasurementSchemas(gomock.Any(), bucketID.String()).
				Return(req)

			a.schemas.EXPECT().
				GetMeasurementSchemasExecute(tmock.MatchedBy(func(in api.ApiGetMeasurementSchemasRequest) bool {
					return (in.GetOrgID() != nil && *in.GetOrgID() == orgID.String()) &&
						in.GetBucketID() == bucketID.String() &&
						(in.GetName() != nil && *in.GetName() == a.params.Name)
				})).
				Return(api.MeasurementSchemaList{
					MeasurementSchemas: []api.MeasurementSchema{
						{
							Id:        measurementID.String(),
							Name:      a.params.Name,
							Columns:   a.cols,
							CreatedAt: createdAt,
							UpdatedAt: updatedAt,
						},
					},
				}, nil)
		}
	}

	opts := func(opts ...optFn) []optFn { return opts }

	lines := func(lines ...string) []string { return lines }

	cases := []struct {
		name     string
		opts     []optFn
		expErr   string
		expLines []string
	}{
		{
			name: "org arg missing",
			opts: opts(
				withArgs(args{BucketName: "my-bucket"}),
			),
			expErr: "org missing: specify org ID or org name",
		},
		{
			name: "bucket args missing",
			opts: opts(
				withArgs(args{OrgName: "my-org"}),
			),
			expErr: "bucket missing: specify bucket ID or bucket name",
		},
		{
			name: "bucket not found",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketName: "my-bucket"}),
				expGetBuckets(),
			),
			expErr: `bucket "my-bucket" not found`,
		},
		{
			name: "bucket not found by id",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketID: influxid.MustIDFromString("baadf00d7777deed")}),
				expGetBuckets(),
			),
			expErr: `bucket "baadf00d7777deed" not found`,
		},
		{
			name: "list succeeds with org name and bucket name",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketName: "my-bucket", Name: "cpu"}),
				withCols("columns.csv"),
				expGetBuckets("my-bucket"),
				expGetMeasurementSchemas(),
			),
			expLines: lines(
				`^ID\s+Measurement Name\s+Bucket ID$`,
				fmt.Sprintf(`^%s\s+cpu\s+%s`, measurementID, bucketID),
			),
		},
		{
			name: "list succeeds with org id and bucket name",
			opts: opts(
				withArgs(args{OrgID: orgID, BucketName: "my-bucket", Name: "cpu"}),
				withCols("columns.csv"),
				expGetBuckets("my-bucket"),
				expGetMeasurementSchemas(),
			),
			expLines: lines(
				`^ID\s+Measurement Name\s+Bucket ID$`,
				fmt.Sprintf(`^%s\s+cpu\s+%s`, measurementID, bucketID),
			),
		},
		{
			name: "list succeeds with org name and bucket id",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketID: bucketID, Name: "cpu"}),
				withCols("columns.csv"),
				expGetBuckets("my-bucket"),
				expGetMeasurementSchemas(),
			),
			expLines: lines(
				`^ID\s+Measurement Name\s+Bucket ID$`,
				fmt.Sprintf(`^%s\s+cpu\s+%s`, measurementID, bucketID),
			),
		},
		{
			name: "list succeeds with org id and bucket id",
			opts: opts(
				withArgs(args{OrgID: orgID, BucketID: bucketID, Name: "cpu"}),
				withCols("columns.csv"),
				expGetBuckets("my-bucket"),
				expGetMeasurementSchemas(),
			),
			expLines: lines(
				`^ID\s+Measurement Name\s+Bucket ID$`,
				fmt.Sprintf(`^%s\s+cpu\s+%s`, measurementID, bucketID),
			),
		},
		{
			name: "list succeeds extended output",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketName: "my-bucket", Name: "cpu", ExtendedOutput: true}),
				withCols("columns.csv"),
				expGetBuckets("my-bucket"),
				expGetMeasurementSchemas(),
			),
			expLines: lines(
				`^ID\s+Measurement Name\s+Column Name\s+Column Type\s+Column Data Type\s+Bucket ID$`,
				fmt.Sprintf(`^%s\s+cpu\s+time\s+timestamp\s+%s`, measurementID, bucketID),
				fmt.Sprintf(`^%s\s+cpu\s+host\s+tag\s+%s`, measurementID, bucketID),
				fmt.Sprintf(`^%s\s+cpu\s+usage_user\s+field\s+float\s+%s`, measurementID, bucketID),
			),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockIO := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			mockIO.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

			args := &setupArgs{
				buckets: mock.NewMockBucketsApi(ctrl),
				schemas: mock.NewMockBucketSchemasApi(ctrl),
				stdio:   mockIO,
				cli:     clients.CLI{StdIO: mockIO},
			}

			for _, opt := range tc.opts {
				opt(t, args)
			}

			c := bucket_schema.Client{
				BucketsApi:       args.buckets,
				BucketSchemasApi: args.schemas,
				CLI:              args.cli,
			}

			err := c.List(context.Background(), args.params)
			if tc.expErr != "" {
				assert.EqualError(t, err, tc.expErr)
			} else {
				require.NoError(t, err)
				testutils.MatchLines(t, tc.expLines, strings.Split(writtenBytes.String(), "\n"))
			}
		})
	}
}
