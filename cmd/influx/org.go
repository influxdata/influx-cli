package main

import (
	"github.com/influxdata/influx-cli/v2/clients/org"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/urfave/cli"
)

func newOrgCmd() cli.Command {
	return cli.Command{
		Name:    "org",
		Aliases: []string{"organization"},
		Usage:   "Organization management commands",
		Subcommands: []cli.Command{
			newOrgCreateCmd(),
			newOrgDeleteCmd(),
			newOrgListCmd(),
			newOrgMembersCmd(),
			newOrgUpdateCmd(),
		},
	}
}

func newOrgCreateCmd() cli.Command {
	var params org.CreateParams
	return cli.Command{
		Name:   "create",
		Usage:  "Create organization",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "Name to set on the new organization",
				Required:    true,
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "Description to set on the new organization",
				Destination: &params.Description,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newOrgDeleteCmd() cli.Command {
	var id influxid.ID
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete organization",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:   "id, i",
				Usage:  "The organization ID",
				EnvVar: "INFLUX_ORG_ID",
				Value:  &id,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.Delete(getContext(ctx), id)
		},
	}
}

func newOrgListCmd() cli.Command {
	var params org.ListParams
	return cli.Command{
		Name:    "list",
		Aliases: []string{"find", "ls"},
		Usage:   "List organizations",
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The organization name",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.Name,
			},
			&cli.GenericFlag{
				Name:   "id, i",
				Usage:  "The organization ID",
				EnvVar: "INFLUX_ORG_ID",
				Value:  &params.ID,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newOrgUpdateCmd() cli.Command {
	var params org.UpdateParams
	return cli.Command{
		Name:   "update",
		Usage:  "Update organization",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.GenericFlag{
				Name:     "id, i",
				Usage:    "The organization ID",
				EnvVar:   "INFLUX_ORG_ID",
				Required: true,
				Value:    &params.ID,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "New name to set on the organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "New description to set on the organization",
				EnvVar:      "INFLUX_ORG_DESCRIPTION",
				Destination: &params.Description,
			},
		),
		Action: func(ctx *cli.Context) error {
			client := org.Client{
				CLI:              getCLI(ctx),
				OrganizationsApi: getAPI(ctx).OrganizationsApi,
			}
			return client.Update(getContext(ctx), &params)
		},
	}
}
