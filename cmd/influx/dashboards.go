package main

import (
	"github.com/influxdata/influx-cli/v2/clients/dashboards"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newDashboardsCommand() cli.Command {
	var params dashboards.Params
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)
	flags = append(flags, &cli.StringSliceFlag{
		Name:  "id, i",
		Usage: "Dashboard ID to retrieve",
	})
	return cli.Command{
		Name:  "dashboards",
		Usage: "List Dashboard(s).",
		Description: `List Dashboard(s).

Examples:
	# list all known Dashboards
	influx dashboards

	# list all known Dashboards matching ids
	influx dashboards --id $ID1 --id $ID2

	# list all known Dashboards matching ids shorts
	influx dashboards -i $ID1 -i $ID2
`,
		Flags:  flags,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			rawIds := ctx.StringSlice("id")
			params.Ids = rawIds

			api := getAPI(ctx)
			client := dashboards.Client{
				CLI:           getCLI(ctx),
				DashboardsApi: api.DashboardsApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}
