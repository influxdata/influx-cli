package main

import (
	"github.com/influxdata/influx-cli/v2/clients/trino"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func withTrinoClient() cli.BeforeFunc {
	return middleware.WithBeforeFns(
		withCli(),
		withApiV1(true),
		func(ctx *cli.Context) error {
			client := getAPIV1(ctx)
			ctx.App.Metadata["trino_client"] = trino.Client{
				QueryApi: client.QueryApi,
				CLI:      getCLI(ctx),
			}
			return nil
		})
}

func getTrinoClient(ctx *cli.Context) trino.Client {
	i, ok := ctx.App.Metadata["trino_client"].(trino.Client)
	if !ok {
		panic("missing Trino client")
	}
	return i
}

func newTrinoCmd() cli.Command {
	return cli.Command{
		Name:  "trino",
		Usage: "Trino management and query commands",
		Subcommands: cli.Commands{
			newTrinoInferSchemaCmd(),
		},
	}
}

func newTrinoInferSchemaCmd() cli.Command {
	params := trino.InferSchemaParams{}

	return cli.Command{
		Name:   "infer-schema",
		Usage:  "Infer schema for provided InfluxQL database",
		Before: withTrinoClient(),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "database",
				Usage:       "Database name",
				EnvVar:      "INFLUX_V1_DB",
				Destination: &params.DB,
				Required:    true,
			},
		),
		Action: func(ctx *cli.Context) error {
			return getTrinoClient(ctx).InferSchema(getContext(ctx), &params)
		},
	}
}
