package main

import (
	"github.com/influxdata/influx-cli/v2/internal/cmd/org"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/urfave/cli/v2"
)

func newOrgCmd() *cli.Command {
	return &cli.Command{
		Name:    "org",
		Aliases: []string{"organization"},
		Usage:   "Organization management commands",
		Subcommands: []*cli.Command{
			newOrgCreateCmd(),
			newOrgDeleteCmd(),
			newOrgListCmd(),
			newOrgUpdateCmd(),
		},
	}
}

func newOrgCreateCmd() *cli.Command {
	var params org.CreateParams
	return &cli.Command{
		Name:   "create",
		Usage:  "Create organization",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name",
				Usage:       "Name to set on the new organization",
				Aliases:     []string{"n"},
				Required:    true,
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description",
				Usage:       "Description to set on the new organization",
				Aliases:     []string{"d"},
				Destination: &params.Description,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.Create(ctx.Context, &params)
		},
	}
}

func newOrgDeleteCmd() *cli.Command {
	var id influxid.ID
	return &cli.Command{
		Name:   "delete",
		Usage:  "Delete organization",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:    "id",
				Usage:   "The organization ID",
				Aliases: []string{"i"},
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &id,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.Delete(ctx.Context, id)
		},
	}
}

func newOrgListCmd() *cli.Command {
	var params org.ListParams
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"find", "ls"},
		Usage:   "List organizations",
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The organization name",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.Name,
			},
			&cli.GenericFlag{
				Name:    "id",
				Usage:   "The organization ID",
				Aliases: []string{"i"},
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &params.ID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.List(ctx.Context, &params)
		},
	}
}

func newOrgUpdateCmd() *cli.Command {
	var params org.UpdateParams
	return &cli.Command{
		Name:   "update",
		Usage:  "Update organization",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:     "id",
				Usage:    "The organization ID",
				Aliases:  []string{"i"},
				EnvVars:  []string{"INFLUX_ORG_ID"},
				Required: true,
				Value:    &params.ID,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "New name to set on the organization",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description",
				Usage:       "New description to set on the organization",
				Aliases:     []string{"d"},
				EnvVars:     []string{"INFLUX_ORG_DESCRIPTION"},
				Destination: &params.Description,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.Update(ctx.Context, &params)
		},
	}
}
