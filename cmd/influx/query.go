package main

import (
	"errors"
	"strings"

	"github.com/influxdata/influx-cli/v2/internal/cmd"
	"github.com/influxdata/influx-cli/v2/internal/cmd/query"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newQueryCmd() *cli.Command {
	var orgParams cmd.OrgParams
	return &cli.Command{
		Name:        "query",
		Usage:       "Execute a Flux query",
		Description: "Execute a Flux query provided via the first argument, a file, or stdin",
		ArgsUsage:   "[query literal or '-' for stdin]",
		Before:      middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlagsNoPrint(),
			&cli.GenericFlag{
				Name:    "org-id",
				Usage:   "The ID of the organization",
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &orgParams.OrgID,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &orgParams.OrgName,
			},
			&cli.StringFlag{
				Name:    "file",
				Usage:   "Path to Flux query file",
				Aliases: []string{"f"},
			},
			&cli.BoolFlag{
				Name:    "raw",
				Usage:   "Display raw query results",
				Aliases: []string{"r"},
			},
			&cli.StringSliceFlag{
				Name:    "profilers",
				Usage:   "Names of Flux profilers to enable",
				Aliases: []string{"p"},
			},
		),
		Action: func(ctx *cli.Context) error {
			queryString, err := readQuery(ctx)
			if err != nil {
				return err
			}
			queryString = strings.TrimSpace(queryString)
			if queryString == "" {
				return errors.New("no query provided")
			}

			// The old CLI allowed specifying this either via repeated flags or
			// via a single flag w/ a comma-separated value.
			rawProfilers := ctx.StringSlice("profilers")
			var profilers []string
			for _, p := range rawProfilers {
				profilers = append(profilers, strings.Split(p, ",")...)
			}

			params := query.Params{
				OrgParams: orgParams,
				Query:     queryString,
				Profilers: profilers,
			}

			var printer query.ResultPrinter
			if ctx.Bool("raw") {
				printer = query.RawResultPrinter
			} else {
				printer = query.NewFormattingPrinter()
			}

			client := query.Client{
				CLI:           getCLI(ctx),
				QueryApi:      getAPI(ctx).QueryApi,
				ResultPrinter: printer,
			}
			return client.Query(ctx.Context, &params)
		},
	}
}
