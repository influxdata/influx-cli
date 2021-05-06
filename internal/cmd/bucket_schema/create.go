package bucket_schema

import (
	"context"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
)

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

	res, err := c.CreateMeasurementSchema(ctx, ids.BucketID).
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
