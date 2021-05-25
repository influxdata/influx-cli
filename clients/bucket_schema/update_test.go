package bucket_schema_test

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/bucket_schema"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClient_Update(t *testing.T) {
	t.Parallel()

	var (
		orgID         = "dead"
		bucketID      = "f00d"
		measurementID = "1010"
		createdAt     = time.Date(2004, 4, 9, 2, 15, 0, 0, time.UTC)
		updatedAt     = time.Date(2009, 9, 1, 2, 15, 0, 0, time.UTC)
	)

	type setupArgs struct {
		buckets *mock.MockBucketsApi
		schemas *mock.MockBucketSchemasApi
		cli     clients.CLI
		params  bucket_schema.UpdateParams
		cols    []api.MeasurementSchemaColumn
		stdio   *mock.MockStdIO
	}

	type optFn func(t *testing.T, a *setupArgs)

	type args struct {
		OrgName        string
		BucketName     string
		Name           string
		ColumnsFile    string
		ExtendedOutput bool
	}

	withArgs := func(args args) optFn {
		return func(t *testing.T, a *setupArgs) {
			t.Helper()
			var colFile string
			if args.ColumnsFile != "" {
				colFile = filepath.Join("testdata", args.ColumnsFile)
			}

			a.params = bucket_schema.UpdateParams{
				OrgBucketParams: clients.OrgBucketParams{
					OrgParams:    clients.OrgParams{OrgName: args.OrgName},
					BucketParams: clients.BucketParams{BucketName: args.BucketName},
				},
				Name:           args.Name,
				ColumnsFile:    colFile,
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
				bucket.SetOrgID(orgID)
				bucket.SetId(bucketID)
				bucket.SetName(n[0])
				buckets = []api.Bucket{*bucket}
			}

			req := api.ApiGetBucketsRequest{ApiService: a.buckets}

			a.buckets.EXPECT().
				GetBuckets(gomock.Any()).
				Return(req)

			a.buckets.EXPECT().
				GetBucketsExecute(tmock.MatchedBy(func(in api.ApiGetBucketsRequest) bool {
					return (in.GetOrg() != nil && *in.GetOrg() == a.params.OrgName) &&
						(in.GetName() != nil && *in.GetName() == a.params.BucketName)
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

	expGetMeasurementSchema := func() optFn {
		return func(t *testing.T, a *setupArgs) {
			t.Helper()

			req := api.ApiGetMeasurementSchemasRequest{ApiService: a.schemas}.BucketID(bucketID)

			a.schemas.EXPECT().
				GetMeasurementSchemas(gomock.Any(), bucketID).
				Return(req)

			a.schemas.EXPECT().
				GetMeasurementSchemasExecute(tmock.MatchedBy(func(in api.ApiGetMeasurementSchemasRequest) bool {
					return (in.GetOrgID() != nil && *in.GetOrgID() == orgID) &&
						in.GetBucketID() == bucketID &&
						(in.GetName() != nil && *in.GetName() == a.params.Name)
				})).
				Return(api.MeasurementSchemaList{
					MeasurementSchemas: []api.MeasurementSchema{
						{
							Id:        measurementID,
							Name:      a.params.Name,
							Columns:   a.cols,
							CreatedAt: createdAt,
							UpdatedAt: updatedAt,
						},
					},
				}, nil)
		}
	}

	expUpdate := func() optFn {
		return func(t *testing.T, a *setupArgs) {
			t.Helper()

			req := api.ApiUpdateMeasurementSchemaRequest{ApiService: a.schemas}.BucketID(bucketID).MeasurementID(measurementID)

			a.schemas.EXPECT().
				UpdateMeasurementSchema(gomock.Any(), bucketID, measurementID).
				Return(req)

			a.schemas.EXPECT().
				UpdateMeasurementSchemaExecute(tmock.MatchedBy(func(in api.ApiUpdateMeasurementSchemaRequest) bool {
					return cmp.Equal(in.GetOrgID(), &orgID) &&
						cmp.Equal(in.GetBucketID(), bucketID) &&
						cmp.Equal(in.GetMeasurementID(), measurementID) &&
						cmp.Equal(in.GetMeasurementSchemaUpdateRequest().Columns, a.cols)
				})).
				Return(api.MeasurementSchema{
					Id:        measurementID,
					Name:      a.params.Name,
					Columns:   a.cols,
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
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
			name:   "unable to guess from stdin",
			expErr: `unable to guess format for file "stdin"`,
		},
		{
			name: "org arg missing",
			opts: opts(
				withArgs(args{BucketName: "my-bucket", ColumnsFile: "columns.csv"}),
			),
			expErr: "org missing: specify org ID or org name",
		},
		{
			name: "bucket arg missing",
			opts: opts(
				withArgs(args{OrgName: "my-org", ColumnsFile: "columns.csv"}),
			),
			expErr: "bucket missing: specify bucket ID or bucket name",
		},
		{
			name: "bucket not found",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketName: "my-bucket", ColumnsFile: "columns.csv"}),
				expGetBuckets(),
			),
			expErr: `bucket "my-bucket" not found`,
		},
		{
			name: "update succeeds",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketName: "my-bucket", Name: "cpu", ColumnsFile: "columns.csv"}),
				withCols("columns.csv"),

				expGetBuckets("my-bucket"),
				expGetMeasurementSchema(),
				expUpdate(),
			),
			expLines: lines(
				`^ID\s+Measurement Name\s+Bucket ID$`,
				`^1010\s+cpu\s+f00d$`,
			),
		},
		{
			name: "update succeeds extended output",
			opts: opts(
				withArgs(args{OrgName: "my-org", BucketName: "my-bucket", Name: "cpu", ColumnsFile: "columns.csv", ExtendedOutput: true}),
				withCols("columns.csv"),
				expGetBuckets("my-bucket"),
				expGetMeasurementSchema(),
				expUpdate(),
			),
			expLines: lines(
				`^ID\s+Measurement Name\s+Column Name\s+Column Type\s+Column Data Type\s+Bucket ID$`,
				`^1010\s+cpu\s+time\s+timestamp\s+f00d$`,
				`^1010\s+cpu\s+host\s+tag\s+f00d$`,
				`^1010\s+cpu\s+usage_user\s+field\s+float\s+f00d$`,
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

			err := c.Update(context.Background(), args.params)
			if tc.expErr != "" {
				assert.EqualError(t, err, tc.expErr)
			} else {
				require.NoError(t, err)
				testutils.MatchLines(t, tc.expLines, strings.Split(writtenBytes.String(), "\n"))
			}
		})
	}
}
