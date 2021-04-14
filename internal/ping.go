package internal

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

type Client interface {
	GetHealthWithResponse(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*api.GetHealthResponse, error)
}

// Ping checks the health of a remote InfluxDB instance.
func (c *CLI) Ping(ctx context.Context, client Client) error {
	resp, err := client.GetHealthWithResponse(ctx, &api.GetHealthParams{ZapTraceSpan: c.TraceId})
	if err != nil {
		return fmt.Errorf("failed to make health check request: %w", err)
	}

	var failureMessage string
	if resp.JSON503 != nil {
		if resp.JSON503.Message != nil {
			failureMessage = *resp.JSON503.Message
		} else {
			failureMessage = "status 503"
		}
	} else if resp.JSONDefault != nil {
		failureMessage = resp.JSONDefault.Error()
	} else if resp.JSON200.Status != api.HealthCheckStatus_pass {
		if resp.JSON200.Message != nil {
			failureMessage = *resp.JSON200.Message
		} else {
			failureMessage = fmt.Sprintf("check %s failed", resp.JSON200.Name)
		}
	}

	if failureMessage != "" {
		return fmt.Errorf("health check failed: %s", failureMessage)
	}
	c.Stdout.Write([]byte("OK\n"))
	return nil
}
