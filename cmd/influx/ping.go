package main

import (
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newPingCmd() *cli.Command {
	return &cli.Command{
		Name:   "ping",
		Usage:  "Check the InfluxDB /health endpoint",
		Before: middleware.WithBeforeFns(withCli(), withApi(false)),
		Flags:  coreFlags,
		Action: func(ctx *cli.Context) error {
			return getCLI(ctx).Ping(ctx.Context, getAPINoToken(ctx).HealthApi)
		},
	}
}
