package internal

import (
	"context"
	"fmt"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

// Ping checks the health of a remote InfluxDB instance.
func (c *CLI) Ping(ctx context.Context, client api.HealthApi) error {
	req := client.GetHealth(ctx)
	if c.TraceId != "" {
		req = req.ZapTraceSpan(c.TraceId)
	}
	resp, _, err := client.GetHealthExecute(req)
	if err != nil {
		return fmt.Errorf("failed to make health check request: %w", err)
	}

	if resp.Status == api.HEALTHCHECKSTATUS_FAIL {
		var message string
		if resp.Message != nil {
			message = *resp.Message
		} else {
			message = fmt.Sprintf("check %s failed", resp.Name)
		}
		return fmt.Errorf("health check failed: %s", message)
	}

	c.Stdout.Write([]byte("OK\n"))
	return nil
}
