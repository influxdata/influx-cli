package measurement_schema

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
)

type Client struct {
	BucketApi        api.BucketsApi
	BucketSchemasApi api.BucketSchemasApi
	CLI              *internal.CLI
}

type orgBucketID struct {
	OrgID    string
	BucketID string
}

func (c Client) resolveMeasurement(ctx context.Context, ids orgBucketID, name string) (string, error) {
	res, _, err := c.BucketSchemasApi.
		GetMeasurementSchemas(ctx, ids.BucketID).
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

func (c Client) resolveOrgBucketIds(ctx context.Context, params internal.OrgBucketParams) (*orgBucketID, error) {
	if params.OrgID.Valid() && params.BucketID.Valid() {
		return &orgBucketID{OrgID: params.OrgID.String(), BucketID: params.BucketID.String()}, nil
	}

	if params.BucketName == "" {
		return nil, errors.New("bucket missing: specify bucket ID or bucket name")
	}

	if !params.OrgID.Valid() && params.OrgName == "" && c.CLI.ActiveConfig.Org == "" {
		return nil, errors.New("org missing: specify org ID or org name")
	}

	req := c.BucketApi.GetBuckets(ctx).Name(params.BucketName)
	if params.OrgID.Valid() {
		req = req.OrgID(params.OrgID.String())
	} else if params.OrgName != "" {
		req = req.Org(params.OrgName)
	} else {
		req = req.Org(c.CLI.ActiveConfig.Org)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to find bucket %q: %w", params.BucketName, err)
	}
	buckets := resp.GetBuckets()
	if len(buckets) == 0 {
		return nil, fmt.Errorf("no %q not found", params.BucketName)
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

type CreateParams struct {
	internal.OrgBucketParams
	Name           string
	Stdin          io.Reader
	ColumnsFile    string
	ColumnsFormat  ColumnsFormat
	ExtendedOutput bool
}

func (c Client) Create(ctx context.Context, params CreateParams) error {
	cols, err := c.readColumns(params.Stdin, params.ColumnsFormat, params.ColumnsFile)
	if err != nil {
		return err
	}

	ids, err := c.resolveOrgBucketIds(ctx, params.OrgBucketParams)
	if err != nil {
		return err
	}

	res, _, err := c.BucketSchemasApi.
		CreateMeasurementSchema(ctx, ids.BucketID).
		OrgID(ids.OrgID).
		MeasurementSchemaCreateRequest(api.MeasurementSchemaCreateRequest{
			Name:    params.Name,
			Columns: cols,
		}).
		Execute()

	if err != nil {
		return fmt.Errorf("failed to create measurement: %w", err)
	}

	return c.printMeasurements(ids.BucketID, []api.MeasurementSchema{res}, params.ExtendedOutput)
}

type UpdateParams struct {
	internal.OrgBucketParams
	Name           string
	ID             string
	Stdin          io.Reader
	ColumnsFile    string
	ColumnsFormat  ColumnsFormat
	ExtendedOutput bool
}

func (c Client) Update(ctx context.Context, params UpdateParams) error {
	cols, err := c.readColumns(params.Stdin, params.ColumnsFormat, params.ColumnsFile)
	if err != nil {
		return err
	}

	ids, err := c.resolveOrgBucketIds(ctx, params.OrgBucketParams)
	if err != nil {
		return err
	}

	var id string
	if params.ID == "" && params.Name == "" {
		return errors.New("measurement id or name required")
	} else if params.ID != "" {
		id = params.ID
	} else {
		id, err = c.resolveMeasurement(ctx, *ids, params.Name)
		if err != nil {
			return err
		}
	}

	res, _, err := c.BucketSchemasApi.
		UpdateMeasurementSchema(ctx, ids.BucketID, id).
		OrgID(ids.OrgID).
		MeasurementSchemaUpdateRequest(api.MeasurementSchemaUpdateRequest{
			Columns: cols,
		}).
		Execute()

	if err != nil {
		return fmt.Errorf("failed to update measurement schema: %w", err)
	}

	return c.printMeasurements(ids.BucketID, []api.MeasurementSchema{res}, params.ExtendedOutput)
}

type ListParams struct {
	internal.OrgBucketParams
	Name           string
	ExtendedOutput bool
}

func (c Client) List(ctx context.Context, params ListParams) error {
	ids, err := c.resolveOrgBucketIds(ctx, params.OrgBucketParams)
	if err != nil {
		return err
	}

	req := c.BucketSchemasApi.
		GetMeasurementSchemas(ctx, ids.BucketID).
		OrgID(ids.OrgID)

	if params.Name != "" {
		req = req.Name(params.Name)
	}

	res, _, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list measurement schemas: %w", err)
	}
	return c.printMeasurements(ids.BucketID, res.MeasurementSchemas, params.ExtendedOutput)
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

	if c.CLI.PrintAsJSON {
		return c.CLI.PrintJSON(m)
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

	return c.CLI.PrintTable(headers, rows...)
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
