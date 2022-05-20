package main

import (
	"github.com/fatih/color"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	repl "github.com/influxdata/influx-cli/v2/clients/v1_repl"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

type Client struct {
	clients.CLI
	api.LegacyQueryApi
}

func newReplCmd() cli.Command {
	var orgParams clients.OrgParams
	persistentQueryParams := repl.DefaultPersistentQueryParams()
	return cli.Command{
		Name:        "repl",
		Usage:       "Start an InfluxQL REPL",
		Description: "Start an InfluxQL REPL",
		// ArgsUsage:   "[query literal or '-' for stdin]",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			append(commonFlagsNoPrint(), getOrgFlags(&orgParams)...),
			// &cli.StringFlag{
			// 	Name:      "file, f",
			// 	Usage:     "Path to Flux query file",
			// 	TakesFile: true,
			// },
			// &cli.BoolFlag{
			// 	Name:  "raw, r",
			// 	Usage: "Display raw query results",
			// },
			// &cli.StringSliceFlag{
			// 	Name:  "profilers, p",
			// 	Usage: "Names of Flux profilers to enable",
			// },
		),
		Action: func(ctx *cli.Context) error {
			if err := checkOrgFlags(&orgParams); err != nil {
				return err
			}
			api := getAPI(ctx)
			c := repl.Client{
				CLI:                   getCLI(ctx),
				PersistentQueryParams: persistentQueryParams,
				PingApi:               api.PingApi,
				LegacyQueryApi:        api.LegacyQueryApi,
				OrganizationsApi:      api.OrganizationsApi,
			}
			color.Cyan("InfluxQL Shell %s", version)
			return c.Create(getContext(ctx))
		},
	}
}
