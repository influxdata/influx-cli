package main

import (
	"github.com/fatih/color"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	shell "github.com/influxdata/influx-cli/v2/clients/v1_shell"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

type Client struct {
	clients.CLI
	api.LegacyQueryApi
}

func newV1ShellCmd() cli.Command {
	var orgParams clients.OrgParams
	persistentQueryParams := shell.DefaultPersistentQueryParams()
	return cli.Command{
		Name:        "shell",
		Usage:       "Start an InfluxQL shell",
		Description: "Start an InfluxQL shell",
		Before:      middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags:       append(commonFlagsNoPrint(), getOrgFlags(&orgParams)...),
		Action: func(ctx *cli.Context) error {
			if err := checkOrgFlags(&orgParams); err != nil {
				return err
			}
			api := getAPI(ctx)
			c := shell.Client{
				CLI:                   getCLI(ctx),
				PersistentQueryParams: persistentQueryParams,
				PingApi:               api.PingApi,
				LegacyQueryApi:        api.LegacyQueryApi,
				OrganizationsApi:      api.OrganizationsApi,
				LegacyWriteApi:        api.LegacyWriteApi,
				DBRPsApi:              api.DBRPsApi,
			}
			color.Cyan("InfluxQL Shell %s", version)
			return c.Create(getContext(ctx))
		},
	}
}
