package bucket_schema

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	api.BucketsApi
	api.BucketSchemasApi
	clients.CLI
}

type orgBucketID struct {
	OrgID    string
	BucketID string
}

func (c Client) resolveMeasurement(ctx context.Context, ids orgBucketID, name string) (string, error) {
	res, err := c.GetMeasurementSchemas(ctx, ids.BucketID).
		OrgID(ids.OrgID).
		Name(name).
		Execute()
	if err != nil {
		return "", fmt.Errorf("failed to find measurement schema: %w", err)
	}

	if len(res.MeasurementSchemas) == 0 {
		return "", fmt.Errorf("measurement schema %q not found", name)
	}

	return res.MeasurementSchemas[0].Id, nil
}

func (c Client) resolveOrgBucketIds(ctx context.Context, params clients.OrgBucketParams) (*orgBucketID, error) {
	if params.BucketName == "" && !params.BucketID.Valid() {
		return nil, errors.New("bucket missing: specify bucket ID or bucket name")
	}

	if !params.OrgID.Valid() && params.OrgName == "" && c.ActiveConfig.Org == "" {
		return nil, errors.New("org missing: specify org ID or org name")
	}

	req := c.GetBuckets(ctx)
	if params.BucketID.Valid() {
		req = req.Id(params.BucketID.String())
	} else {
		req = req.Name(params.BucketName)
	}
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	} else if params.OrgName != "" {
		req = req.Org(params.OrgName)
	} else {
		req = req.Org(c.ActiveConfig.Org)
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to find bucket %q: %w", params.BucketName, err)
	}
	buckets := resp.GetBuckets()
	if len(buckets) == 0 {
		return nil, fmt.Errorf("bucket %q not found", params.BucketName)
	}

	return &orgBucketID{OrgID: buckets[0].GetOrgID(), BucketID: buckets[0].GetId()}, nil
}

func (c Client) readColumns(stdin io.Reader, f ColumnsFormat, path string) ([]api.MeasurementSchemaColumn, error) {
	var (
		r    io.Reader
		name string
	)

	if path == "" {
		r = stdin
		name = "stdin"
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("unable to read file %q: %w", path, err)
		}
		r = bytes.NewReader(data)
		name = path
	}

	reader, err := f.DecoderFn(name)
	if err != nil {
		return nil, err
	}

	return reader(r)
}

// Constants for table column headers
const (
	IDHdr              = "ID"
	MeasurementNameHdr = "Measurement Name"
	ColumnNameHdr      = "Column Name"
	ColumnTypeHdr      = "Column Type"
	ColumnDataTypeHdr  = "Column Data Type"
	BucketIDHdr        = "Bucket ID"
)

func (c Client) printMeasurements(bucketID string, m []api.MeasurementSchema, extended bool) error {
	if len(m) == 0 {
		return nil
	}

	if c.PrintAsJSON {
		return c.PrintJSON(m)
	}

	var headers []string
	if extended {
		headers = []string{
			IDHdr,
			MeasurementNameHdr,
			ColumnNameHdr,
			ColumnTypeHdr,
			ColumnDataTypeHdr,
			BucketIDHdr,
		}
	} else {
		headers = []string{
			IDHdr,
			MeasurementNameHdr,
			BucketIDHdr,
		}
	}

	var makeRow measurementRowFn
	if extended {
		makeRow = makeExtendedMeasurementRows
	} else {
		makeRow = makeMeasurementRows
	}

	var rows []map[string]interface{}

	for i := range m {
		rows = append(rows, makeRow(bucketID, &m[i])...)
	}

	return c.PrintTable(headers, rows...)
}

type measurementRowFn func(bucketID string, m *api.MeasurementSchema) []map[string]interface{}

func makeMeasurementRows(bucketID string, m *api.MeasurementSchema) []map[string]interface{} {
	return []map[string]interface{}{
		{
			IDHdr:              m.Id,
			MeasurementNameHdr: m.Name,
			BucketIDHdr:        bucketID,
		},
	}
}

func makeExtendedMeasurementRows(bucketID string, m *api.MeasurementSchema) []map[string]interface{} {
	rows := make([]map[string]interface{}, 0, len(m.Columns))

	for i := range m.Columns {
		col := &m.Columns[i]
		rows = append(rows, map[string]interface{}{
			IDHdr:              m.Id,
			MeasurementNameHdr: m.Name,
			ColumnNameHdr:      col.Name,
			ColumnTypeHdr:      col.Type,
			ColumnDataTypeHdr:  col.GetDataType(),
			BucketIDHdr:        bucketID,
		})
	}
	return rows
}
