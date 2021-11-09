package main

import (
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients/bucket"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newBucketCmd() cli.Command {
	return cli.Command{
		Name:   "bucket",
		Usage:  "Bucket management commands",
		Before: middleware.NoArgs,
		Subcommands: []cli.Command{
			newBucketCreateCmd(),
			newBucketDeleteCmd(),
			newBucketListCmd(),
			newBucketUpdateCmd(),
		},
	}
}

func newBucketCreateCmd() cli.Command {
	params := bucket.BucketsCreateParams{
		SchemaType: api.SCHEMATYPE_IMPLICIT,
	}
	return cli.Command{
		Name:   "create",
		Usage:  "Create bucket",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "New bucket name",
				EnvVar:      "INFLUX_BUCKET_NAME",
				Destination: &params.Name,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "Description of the bucket that will be created",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "retention, r",
				Usage:       "Duration bucket will retain data, or 0 for infinite",
				Destination: &params.Retention,
			},
			&cli.StringFlag{
				Name:        "shard-group-duration",
				Usage:       "Shard group duration used internally by the storage engine",
				Destination: &params.ShardGroupDuration,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The name of the organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.GenericFlag{
				Name:  "schema-type",
				Usage: "The schema type (implicit, explicit)",
				Value: &params.SchemaType,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := bucket.Client{
				CLI:              getCLI(ctx),
				BucketsApi:       api.BucketsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Create(getContext(ctx), &params)
		},
	}
}

func newBucketDeleteCmd() cli.Command {
	var params bucket.BucketsDeleteParams
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete bucket",
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id, i",
				Usage:       "The bucket ID, required if name isn't provided",
				Destination: &params.ID,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The bucket name, org or org-id will be required by choosing this",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The name of the organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := bucket.Client{
				CLI:              getCLI(ctx),
				BucketsApi:       api.BucketsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Delete(getContext(ctx), &params)
		},
	}
}

func newBucketListCmd() cli.Command {
	var params bucket.BucketsListParams
	return cli.Command{
		Name:    "list",
		Usage:   "List buckets",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "id, i",
				Usage:       "The bucket ID, required if name isn't provided",
				Destination: &params.ID,
			},
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "The bucket name, org or org-id will be required by choosing this",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The name of the organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.OrgName,
			},
			&cli.IntFlag{
				Name:        "limit",
				Usage:       "Total number of buckets to fetch from the server, or 0 to return all buckets",
				Destination: &params.Limit,
			},
			&cli.IntFlag{
				Name:        "offset",
				Usage:       "Number of buckets to skip over in the list",
				Destination: &params.Offset,
			},
			&cli.IntFlag{
				Name:        "page-size",
				Usage:       "Number of buckets to fetch per request to the server",
				Value:       20,
				Destination: &params.PageSize,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := bucket.Client{
				CLI:              getCLI(ctx),
				BucketsApi:       api.BucketsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newBucketUpdateCmd() cli.Command {
	var params bucket.BucketsUpdateParams
	return cli.Command{
		Name:    "update",
		Usage:   "Update bucket",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "New name to set on the bucket",
				EnvVar:      "INFLUX_BUCKET_NAME",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "id, i",
				Usage:       "The bucket ID",
				Required:    true,
				Destination: &params.ID,
			},
			&cli.StringFlag{
				Name:        "description, d",
				Usage:       "New description to set on the bucket",
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "retention, r",
				Usage:       "New retention duration to set on the bucket, or 0 for infinite",
				Destination: &params.Retention,
			},
			&cli.StringFlag{
				Name:        "shard-group-duration",
				Usage:       "New shard group duration to set on the bucket, or 0 to have the server calculate a value",
				Destination: &params.ShardGroupDuration,
			},
		),
		Action: func(ctx *cli.Context) error {
			api := getAPI(ctx)
			client := bucket.Client{
				CLI:              getCLI(ctx),
				BucketsApi:       api.BucketsApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Update(getContext(ctx), &params)
		},
	}
}
