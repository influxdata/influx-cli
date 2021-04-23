package internal

import (
	"context"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

// Ping checks the health of a remote InfluxDB instance.
func (c *CLI) Ping(ctx context.Context, client api.HealthApi) error {
	req := client.GetHealth(ctx)
	if c.TraceId != "" {
		req = req.ZapTraceSpan(c.TraceId)
	}
	if _, _, err := client.GetHealthExecute(req); err != nil {
		return err
	}

	_, err := c.StdIO.Write([]byte("OK\n"))
	return err
}
