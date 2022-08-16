package main

import (
	v1dbrps "github.com/influxdata/influx-cli/v2/clients/v1_dbrps"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newV1DBRPCmd() cli.Command {
	return cli.Command{
		Name:   "dbrp",
		Usage:  "Commands to manage database and retention policy mappings for v1 APIs",
		Before: middleware.NoArgs,
		Subcommands: []cli.Command{
			newV1DBRPListCmd(),
			newV1DBRPCreateCmd(),
			newV1DBRPDeleteCmd(),
			newV1DBRPUpdateCmd(),
		},
	}
}

func newV1DBRPListCmd() cli.Command {
	var params v1dbrps.ListParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)

	return cli.Command{
		Name:  "list",
		Usage: "List database and retention policy mappings",
		Description: `List database and retention policy mappings, both standard and virtual.

   Virtual DBRP mappings (InfluxDB OSS only) are created automatically using the bucket name.
   Virtual mappings are read-only. To modify a virtual DBRP mapping, create a new, explicit DBRP mapping.
   For more information, see https://docs.influxdata.com/influxdb/latest/query-data/influxql/dbrp/`,
		Aliases: []string{"find", "ls"},
		Flags: append(
			flags,
			&cli.StringFlag{
				Name:        "bucket-id",
				Usage:       "Limit results to the matching bucket id",
				Destination: &params.BucketID,
			},
			&cli.StringFlag{
				Name:        "db",
				Usage:       "Limit results to the matching database name",
				Destination: &params.DB,
			},
			&cli.BoolFlag{
				Name:        "default",
				Usage:       "Limit results to default mappings",
				Destination: &params.Default,
			},
			&cli.StringFlag{
				Name:        "id",
				Usage:       "Limit results to a single mapping",
				Destination: &params.ID,
			},
			&cli.StringFlag{
				Name:        "rp",
				Usage:       "Limit results to the matching retention policy name",
				Destination: &params.RP,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if err := checkOrgFlags(&params.OrgParams); err != nil {
				return err
			}
			api := getAPI(ctx)
			client := v1dbrps.Client{
				CLI:      getCLI(ctx),
				DBRPsApi: api.DBRPsApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newV1DBRPCreateCmd() cli.Command {
	var params v1dbrps.CreateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)

	return cli.Command{
		Name:  "create",
		Usage: "Create a database and retention policy mapping to an existing bucket",
		Flags: append(
			flags,
			&cli.StringFlag{
				Name:        "bucket-id",
				Usage:       "The ID of the bucket to be mapped",
				Destination: &params.BucketID,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "db",
				Usage:       "The name of the database",
				Destination: &params.DB,
				Required:    true,
			},
			&cli.BoolFlag{
				Name:        "default",
				Usage:       "Identify this retention policy as the default for the database",
				Destination: &params.Default,
			},
			&cli.StringFlag{
				Name:        "rp",
				Usage:       "The name of the retention policy",
				Destination: &params.RP,
				Required:    true,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if err := checkOrgFlags(&params.OrgParams); err != nil {
				return err
			}
			api := getAPI(ctx)
			client := v1dbrps.Client{
				CLI:              getCLI(ctx),
				DBRPsApi:         api.DBRPsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newV1DBRPDeleteCmd() cli.Command {
	var params v1dbrps.DeleteParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)

	return cli.Command{
		Name:  "delete",
		Usage: "Delete a database and retention policy mapping",
		Flags: append(
			flags,
			&cli.StringFlag{
				Name:        "id",
				Usage:       "The ID of the mapping to delete",
				Destination: &params.ID,
				Required:    true,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if err := checkOrgFlags(&params.OrgParams); err != nil {
				return err
			}
			api := getAPI(ctx)
			client := v1dbrps.Client{
				CLI:      getCLI(ctx),
				DBRPsApi: api.DBRPsApi,
			}
			return client.Delete(getContext(ctx), &params)
		},
	}
}

func newV1DBRPUpdateCmd() cli.Command {
	var params v1dbrps.UpdateParams
	flags := append(commonFlags(), getOrgFlags(&params.OrgParams)...)

	return cli.Command{
		Name:  "update",
		Usage: "Update a database and retention policy mapping",
		Flags: append(
			flags,
			&cli.StringFlag{
				Name:        "id",
				Usage:       "The ID of the mapping to update",
				Destination: &params.ID,
				Required:    true,
			},
			&cli.BoolFlag{
				Name:        "default",
				Usage:       "Set this mapping's retention policy as the default for the mapping's database",
				Destination: &params.Default,
			},
			&cli.StringFlag{
				Name:        "rp",
				Usage:       "The updated name of the retention policy",
				Destination: &params.RP,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			if err := checkOrgFlags(&params.OrgParams); err != nil {
				return err
			}
			api := getAPI(ctx)
			client := v1dbrps.Client{
				CLI:      getCLI(ctx),
				DBRPsApi: api.DBRPsApi,
			}
			return client.Update(getContext(ctx), &params)
		},
	}
}
