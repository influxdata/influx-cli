package main

import (
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newSetupCmd() *cli.Command {
	var params internal.SetupParams
	return &cli.Command{
		Name:   "setup",
		Usage:  "Setup instance with initial user, org, bucket",
		Before: middleware.WithBeforeFns(withCli(), withApi(false)),
		Flags: append(
			commonFlagsNoToken,
			&cli.StringFlag{
				Name:        "username",
				Usage:       "Name of initial user to create",
				Aliases:     []string{"u"},
				Destination: &params.Username,
			},
			&cli.StringFlag{
				Name:        "password",
				Usage:       "Password to set on initial user",
				Aliases:     []string{"p"},
				Destination: &params.Password,
			},
			&cli.StringFlag{
				Name:        tokenFlag,
				Usage:       "Auth token to set on the initial user",
				Aliases:     []string{"t"},
				EnvVars:     []string{"INFLUX_TOKEN"},
				DefaultText: "auto-generated",
				Destination: &params.AuthToken,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "Name of initial organization to create",
				Aliases:     []string{"o"},
				Destination: &params.Org,
			},
			&cli.StringFlag{
				Name:        "bucket",
				Usage:       "Name of initial bucket to create",
				Aliases:     []string{"b"},
				Destination: &params.Bucket,
			},
			&cli.StringFlag{
				Name:        "retention",
				Usage:       "Duration initial bucket will retain data, or 0 for infinite",
				Aliases:     []string{"r"},
				DefaultText: "infinite",
				Destination: &params.Retention,
			},
			&cli.BoolFlag{
				Name:        "force",
				Usage:       "Skip confirmation prompt",
				Aliases:     []string{"f"},
				Destination: &params.Force,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "Name to set on CLI config generated for the InfluxDB instance, required if other configs exist",
				Aliases:     []string{"n"},
				Destination: &params.ConfigName,
			},
		),
		Action: func(ctx *cli.Context) error {
			return getCLI(ctx).Setup(ctx.Context, getAPINoToken(ctx).SetupApi, &params)
		},
	}
}
