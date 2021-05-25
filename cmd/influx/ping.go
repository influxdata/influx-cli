package main

import (
	"github.com/influxdata/influx-cli/v2/clients/ping"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newPingCmd() *cli.Command {
	return &cli.Command{
		Name:   "ping",
		Usage:  "Check the InfluxDB /health endpoint",
		Before: middleware.WithBeforeFns(withCli(), withApi(false)),
		Flags:  coreFlags(),
		Action: func(ctx *cli.Context) error {
			client := ping.Client{
				CLI:       getCLI(ctx),
				HealthApi: getAPINoToken(ctx).HealthApi,
			}
			return client.Ping(ctx.Context)
		},
	}
}
