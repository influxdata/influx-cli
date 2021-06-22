package main

import "github.com/urfave/cli/v2"

func newV1SubCommand() *cli.Command {
	return &cli.Command{
		Name:  "v1",
		Usage: "InfluxDB v1 management commands",
		Subcommands: []*cli.Command{
			newV1DBRPCmd(),
		},
	}
}
