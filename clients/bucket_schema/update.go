package bucket_schema

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/internal/api"
)

type UpdateParams struct {
	clients.OrgBucketParams
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

	res, err := c.UpdateMeasurementSchema(ctx, ids.BucketID, id).
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
