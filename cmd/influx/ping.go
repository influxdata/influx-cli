package main

import (
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

var pingCmd = cli.Command{
	Name:   "ping",
	Usage:  "Check the InfluxDB /health endpoint",
	Before: middleware.WithBeforeFns(withCli(), withApi()),
	Flags:  coreFlags,
	Action: func(ctx *cli.Context) error {
		cli := getCLI(ctx)
		client := getAPINoToken(ctx)
		return cli.Ping(ctx.Context, client.HealthApi)
	},
}
