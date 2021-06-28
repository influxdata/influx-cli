package main

import (
	"github.com/influxdata/influx-cli/v2/clients/setup"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newSetupCmd() cli.Command {
	var params setup.Params
	return cli.Command{
		Name:   "setup",
		Usage:  "Setup instance with initial user, org, bucket",
		Before: middleware.WithBeforeFns(withCli(), withApi(false)),
		Flags: append(
			commonFlagsNoToken(),
			&cli.StringFlag{
				Name:        "username, u",
				Usage:       "Name of initial user to create",
				Destination: &params.Username,
			},
			&cli.StringFlag{
				Name:        "password, p",
				Usage:       "Password to set on initial user",
				Destination: &params.Password,
			},
			&cli.StringFlag{
				Name:        tokenFlagName + ", t",
				Usage:       "Auth token to set on the initial user",
				EnvVar:      "INFLUX_TOKEN",
				Destination: &params.AuthToken,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "Name of initial organization to create",
				Destination: &params.Org,
			},
			&cli.StringFlag{
				Name:        "bucket, b",
				Usage:       "Name of initial bucket to create",
				Destination: &params.Bucket,
			},
			&cli.StringFlag{
				Name:        "retention, r",
				Usage:       "Duration initial bucket will retain data, or 0 for infinite",
				Destination: &params.Retention,
			},
			&cli.BoolFlag{
				Name:        "force, f",
				Usage:       "Skip confirmation prompt",
				Destination: &params.Force,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Name to set on CLI config generated for the InfluxDB instance, required if other configs exist",
				Destination: &params.ConfigName,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := setup.Client{
				CLI:      getCLI(ctx),
				SetupApi: getAPINoToken(ctx).SetupApi,
			}
			return client.Setup(getContext(ctx), &params)
		},
	}
}
