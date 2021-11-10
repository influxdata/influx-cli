package main

import (
	"errors"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/query"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newQueryCmd() cli.Command {
	var orgParams clients.OrgParams
	return cli.Command{
		Name:        "query",
		Usage:       "Execute a Flux query",
		Description: "Execute a Flux query provided via the first argument, a file, or stdin",
		ArgsUsage:   "[query literal or '-' for stdin]",
		Before:      middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			append(commonFlagsNoPrint(), getOrgFlags(&orgParams)...),
			&cli.StringFlag{
				Name:      "file, f",
				Usage:     "Path to Flux query file",
				TakesFile: true,
			},
			&cli.BoolFlag{
				Name:  "raw, r",
				Usage: "Display raw query results",
			},
			&cli.StringSliceFlag{
				Name:  "profilers, p",
				Usage: "Names of Flux profilers to enable",
			},
		),
		Action: func(ctx *cli.Context) error {
			queryString, err := clients.ReadQuery(ctx.String("file"), ctx.Args())
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
			return client.Query(getContext(ctx), &params)
		},
	}
}
