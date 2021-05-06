package bucket_schema

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/internal/cmd"
)

type ListParams struct {
	cmd.OrgBucketParams
	Name           string
	ExtendedOutput bool
}

func (c Client) List(ctx context.Context, params ListParams) error {
	ids, err := c.resolveOrgBucketIds(ctx, params.OrgBucketParams)
	if err != nil {
		return err
	}

	req := c.GetMeasurementSchemas(ctx, ids.BucketID).OrgID(ids.OrgID)

	if params.Name != "" {
		req = req.Name(params.Name)
	}

	res, err := req.Execute()
	if err != nil {
		return fmt.Errorf("failed to list measurement schemas: %w", err)
	}
	return c.printMeasurements(ids.BucketID, res.MeasurementSchemas, params.ExtendedOutput)
}
