package main

import (
	"errors"

	"github.com/influxdata/influx-cli/v2/clients/server_config"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newServerConfigCommand() cli.Command {
	var params server_config.ListParams
	return cli.Command{
		Name:  "server-config",
		Usage: "Display server config",
		Flags: append(
			commonFlags(),
			&cli.BoolFlag{
				Name:        "toml",
				Usage:       "Output configuration as TOML instead of JSON",
				Destination: &params.TOML,
			},
			&cli.BoolFlag{
				Name:        "yaml",
				Usage:       "Output configuration as YAML instead of JSON",
				Destination: &params.YAML,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if params.TOML && params.YAML {
				return errors.New("cannot specify both TOML and YAML simultaneously")
			}

			api := getAPI(ctx)
			client := server_config.Client{
				CLI:       getCLI(ctx),
				ConfigApi: api.ConfigApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}
