package ping

import (
	"context"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
)

type Client struct {
	internal.CLI
	api.HealthApi
}

// Ping checks the health of a remote InfluxDB instance.
func (c Client) Ping(ctx context.Context) error {
	if _, err := c.GetHealth(ctx).Execute(); err != nil {
		return err
	}
	_, err := c.StdIO.Write([]byte("OK\n"))
	return err
}
