package ping

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/kit/tracing"
)

type Client interface {
	GetHealthWithResponse(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*api.GetHealthResponse, error)
}

// Ping checks the health of a remote InfluxDB instance.
func Ping(ctx tracing.Context, client Client) error {
	resp, err := client.GetHealthWithResponse(ctx, &api.GetHealthParams{ZapTraceSpan: ctx.TraceId()})
	if err != nil {
		return err
	}

	if resp.StatusCode()/100 != 2 {
		return fmt.Errorf("health check failed (got status %d): %s", resp.StatusCode(), string(resp.Body))
	}
	if resp.JSON200 == nil {
		return fmt.Errorf("got unexpected response body for healthcheck: %s", string(resp.Body))
	}
	if resp.JSON200.Status != api.HealthCheckStatus_pass {
		return fmt.Errorf("health check failed: %s", *resp.JSON200.Message)
	}

	return nil
}
