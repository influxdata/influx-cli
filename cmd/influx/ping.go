package main

import "github.com/urfave/cli/v2"

var pingCmd = cli.Command{
	Name:  "ping",
	Usage: "Check the InfluxDB /health endpoint",
	Flags: coreFlags,
	Action: func(ctx *cli.Context) error {
		cli, err := newCli(ctx)
		if err != nil {
			return err
		}
		client, err := newApiClient(ctx, cli, false)
		if err != nil {
			return err
		}
		return cli.Ping(standardCtx(ctx), client.HealthApi)
	},
}
