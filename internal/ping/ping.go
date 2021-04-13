package ping

import (
	"context"
	"fmt"
	"net/http"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/pkg/tracing"
)

type Client interface {
	GetHealth(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*http.Response, error)
}

// Ping checks the health of a remote InfluxDB instance.
func Ping(ctx tracing.Context, client Client) error {
	resp, err := client.GetHealth(ctx, &api.GetHealthParams{ZapTraceSpan: ctx.TraceId()})
	if err != nil {
		return fmt.Errorf("failed to make health check request: %w", err)
	}
	parsed, err := api.ParseGetHealthResponse(resp)
	if err != nil {
		return fmt.Errorf("failed to parse health check response: %w", err)
	}

	if parsed.StatusCode()/100 != 2 {
		return fmt.Errorf("health check failed (got status %d): %s", parsed.StatusCode(), string(parsed.Body))
	}
	if parsed.JSON200 == nil {
		return fmt.Errorf("health check returned malformed body: %s", string(parsed.Body))
	}
	if parsed.JSON200.Status != api.HealthCheckStatus_pass {
		return fmt.Errorf("health check failed: %s", *parsed.JSON200.Message)
	}

	return nil
}
