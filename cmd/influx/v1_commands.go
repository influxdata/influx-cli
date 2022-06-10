package main

import (
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newV1SubCommand() cli.Command {
	return cli.Command{
		Name:   "v1",
		Usage:  "InfluxDB v1 management commands",
		Before: middleware.NoArgs,
		Subcommands: []cli.Command{
			newV1DBRPCmd(),
			newV1AuthCommand(),
			newV1ReplCmd(),
		},
	}
}
