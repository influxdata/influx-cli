package main

import (
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func withBucketsClient() cli.BeforeFunc {
	return middleware.WithBeforeFns(
		withCli(),
		withApi(true),
		func(ctx *cli.Context) error {
			client := getAPI(ctx)
			ctx.App.Metadata["bucketsClient"] = internal.BucketsClients{
				BucketApi: client.BucketsApi,
				OrgApi:    client.OrganizationsApi,
			}
			return nil
		},
	)
}

func getBucketsClient(ctx *cli.Context) internal.BucketsClients {
	i, ok := ctx.App.Metadata["bucketsClient"].(internal.BucketsClients)
	if !ok {
		panic("missing buckets client")
	}
	return i
}

func newBucketCmd() *cli.Command {
	return &cli.Command{
		Name:  "bucket",
		Usage: "Bucket management commands",
		Subcommands: []*cli.Command{
			newBucketCreateCmd(),
			newBucketDeleteCmd(),
			newBucketListCmd(),
			newBucketUpdateCmd(),
		},
	}

}

func newBucketCreateCmd() *cli.Command {
	var params internal.BucketsCreateParams
	return &cli.Command{
		Name:   "create",
		Usage:  "Create bucket",
		Before: withBucketsClient(),
		Flags: append(
			commonFlags,
			&cli.StringFlag{
				Name:        "name",
				Usage:       "New bucket name",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_BUCKET_NAME"},
				Destination: &params.Name,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "description",
				Usage:       "Description of the bucket that will be created",
				Aliases:     []string{"d"},
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "retention",
				Usage:       "Duration bucket will retain data, or 0 for infinite",
				Aliases:     []string{"r"},
				DefaultText: "infinite",
				Destination: &params.Retention,
			},
			&cli.StringFlag{
				Name:        "shard-group-duration",
				Usage:       "Shard group duration used internally by the storage engine",
				DefaultText: "calculated from retention",
				Destination: &params.ShardGroupDuration,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVars:     []string{"INFLUX_ORG_ID"},
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
		),
		Action: func(ctx *cli.Context) error {
			clients := getBucketsClient(ctx)
			return getCLI(ctx).BucketsCreate(ctx.Context, &clients, &params)
		},
	}
}

func newBucketDeleteCmd() *cli.Command {
	var params internal.BucketsDeleteParams
	return &cli.Command{
		Name:   "delete",
		Usage:  "Delete bucket",
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags,
			&cli.StringFlag{
				Name:        "id",
				Usage:       "The bucket ID, required if name isn't provided",
				Aliases:     []string{"i"},
				Destination: &params.ID,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The bucket name, org or org-id will be required by choosing this",
				Aliases:     []string{"n"},
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVars:     []string{"INFLUX_ORG_ID"},
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
		),
		Action: func(ctx *cli.Context) error {
			return getCLI(ctx).BucketsDelete(ctx.Context, getAPI(ctx).BucketsApi, &params)
		},
	}
}

func newBucketListCmd() *cli.Command {
	var params internal.BucketsListParams
	return &cli.Command{
		Name:    "list",
		Usage:   "List buckets",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags,
			&cli.StringFlag{
				Name:        "id",
				Usage:       "The bucket ID, required if name isn't provided",
				Aliases:     []string{"i"},
				Destination: &params.ID,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "The bucket name, org or org-id will be required by choosing this",
				Aliases:     []string{"n"},
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVars:     []string{"INFLUX_ORG_ID"},
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.OrgName,
			},
		),
		Action: func(ctx *cli.Context) error {
			return getCLI(ctx).BucketsList(ctx.Context, getAPI(ctx).BucketsApi, &params)
		},
	}
}

func newBucketUpdateCmd() *cli.Command {
	var params internal.BucketsUpdateParams
	return &cli.Command{
		Name:    "update",
		Usage:   "Update bucket",
		Aliases: []string{"find", "ls"},
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags,
			&cli.StringFlag{
				Name:        "name",
				Usage:       "New name to set on the bucket",
				Aliases:     []string{"n"},
				EnvVars:     []string{"INFLUX_BUCKET_NAME"},
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "id",
				Usage:       "The bucket ID",
				Aliases:     []string{"i"},
				Required:    true,
				Destination: &params.ID,
			},
			&cli.StringFlag{
				Name:        "description",
				Usage:       "New description to set on the bucket",
				Aliases:     []string{"d"},
				Destination: &params.Description,
			},
			&cli.StringFlag{
				Name:        "retention",
				Usage:       "New retention duration to set on the bucket, or 0 for infinite",
				Aliases:     []string{"r"},
				Destination: &params.Retention,
			},
			&cli.StringFlag{
				Name:        "shard-group-duration",
				Usage:       "New shard group duration to set on the bucket, or 0 to have the server calculate a value",
				Destination: &params.ShardGroupDuration,
			},
		),
		Action: func(ctx *cli.Context) error {
			return getCLI(ctx).BucketsUpdate(ctx.Context, getAPI(ctx).BucketsApi, &params)
		},
	}
}
