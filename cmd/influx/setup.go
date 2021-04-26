package main

import (
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/urfave/cli/v2"
)

var setupCmd = cli.Command{
	Name:  "setup",
	Usage: "Setup instance with initial user, org, bucket",
	Flags: append(
		commonFlagsNoToken,
		&cli.StringFlag{
			Name:    "username",
			Usage:   "Name of initial user to create",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    "password",
			Usage:   "Password to set on initial user",
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:        tokenFlag,
			Usage:       "Auth token to set on the initial user",
			Aliases:     []string{"t"},
			EnvVars:     []string{"INFLUX_TOKEN"},
			DefaultText: "auto-generated",
		},
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Name of initial organization to create",
			Aliases: []string{"o"},
		},
		&cli.StringFlag{
			Name:    "bucket",
			Usage:   "Name of initial bucket to create",
			Aliases: []string{"b"},
		},
		&cli.StringFlag{
			Name:        "retention",
			Usage:       "Duration initial bucket will retain data, or 0 for infinite",
			Aliases:     []string{"r"},
			DefaultText: "infinite",
		},
		&cli.BoolFlag{
			Name:    "force",
			Usage:   "Skip confirmation prompt",
			Aliases: []string{"f"},
		},
		&cli.StringFlag{
			Name:    "name",
			Usage:   "Name to set on CLI config generated for the InfluxDB instance, required if other configs exist",
			Aliases: []string{"n"},
		},
		&cli.BoolFlag{
			Name:    printJsonFlag,
			Usage:   "Output data as JSON",
			EnvVars: []string{"INFLUX_OUTPUT_JSON"},
		},
		&cli.BoolFlag{
			Name:    hideHeadersFlag,
			Usage:   "Hide the table headers in output data",
			EnvVars: []string{"INFLUX_HIDE_HEADERS"},
		},
	),
	Action: func(ctx *cli.Context) error {
		cli, err := newCli(ctx)
		if err != nil {
			return err
		}
		client, err := newApiClient(ctx, cli, false)
		if err != nil {
			return err
		}
		return cli.Setup(standardCtx(ctx), client.SetupApi, &internal.SetupParams{
			Username:   ctx.String("username"),
			Password:   ctx.String("password"),
			AuthToken:  ctx.String(tokenFlag),
			Org:        ctx.String("org"),
			Bucket:     ctx.String("bucket"),
			Retention:  ctx.String("retention"),
			Force:      ctx.Bool("force"),
			ConfigName: ctx.String("name"),
		})
	},
}
