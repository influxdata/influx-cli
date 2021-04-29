package main

import (
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/urfave/cli/v2"
)

var bucketCmd = cli.Command{
	Name:  "bucket",
	Usage: "Bucket management commands",
	Subcommands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create bucket",
			Flags: append(
				commonFlags,
				&cli.StringFlag{
					Name:     "name",
					Usage:    "New bucket name",
					Aliases:  []string{"n"},
					EnvVars:  []string{"INFLUX_BUCKET_NAME"},
					Required: true,
				},
				&cli.StringFlag{
					Name:    "description",
					Usage:   "Description of the bucket that will be created",
					Aliases: []string{"d"},
				},
				&cli.StringFlag{
					Name:        "retention",
					Usage:       "Duration bucket will retain data, or 0 for infinite",
					Aliases:     []string{"r"},
					DefaultText: "infinite",
				},
				&cli.StringFlag{
					Name:        "shard-group-duration",
					Usage:       "Shard group duration used internally by the storage engine",
					DefaultText: "calculated from retention",
				},
				&cli.StringFlag{
					Name:    "org-id",
					Usage:   "The ID of the organization",
					EnvVars: []string{"INFLUX_ORG_ID"},
				},
				&cli.StringFlag{
					Name:    "org",
					Usage:   "The name of the organization",
					Aliases: []string{"o"},
					EnvVars: []string{"INFLUX_ORG"},
				},
			),
			Action: func(ctx *cli.Context) error {
				cli, err := newCli(ctx)
				if err != nil {
					return err
				}
				client, err := newApiClient(ctx, cli, true)
				if err != nil {
					return err
				}
				clients := internal.BucketsClients{
					BucketApi: client.BucketsApi,
					OrgApi:    client.OrganizationsApi,
				}
				return cli.BucketsCreate(standardCtx(ctx), &clients, &internal.BucketsCreateParams{
					OrgID:              ctx.String("org-id"),
					OrgName:            ctx.String("org"),
					Name:               ctx.String("name"),
					Description:        ctx.String("description"),
					Retention:          ctx.String("retention"),
					ShardGroupDuration: ctx.String("shard-group-duration"),
				})
			},
		},
		{
			Name:  "delete",
			Usage: "Delete bucket",
			Flags: append(
				commonFlags,
				&cli.StringFlag{
					Name:    "id",
					Usage:   "The bucket ID, required if name isn't provided",
					Aliases: []string{"i"},
				},
				&cli.StringFlag{
					Name:    "name",
					Usage:   "The bucket name, org or org-id will be required by choosing this",
					Aliases: []string{"n"},
				},
				&cli.StringFlag{
					Name:    "org-id",
					Usage:   "The ID of the organization",
					EnvVars: []string{"INFLUX_ORG_ID"},
				},
				&cli.StringFlag{
					Name:    "org",
					Usage:   "The name of the organization",
					Aliases: []string{"o"},
					EnvVars: []string{"INFLUX_ORG"},
				},
			),
			Action: func(ctx *cli.Context) error {
				cli, err := newCli(ctx)
				if err != nil {
					return err
				}
				client, err := newApiClient(ctx, cli, true)
				if err != nil {
					return err
				}
				clients := internal.BucketsClients{
					BucketApi: client.BucketsApi,
					OrgApi:    client.OrganizationsApi,
				}
				return cli.BucketsDelete(standardCtx(ctx), &clients, &internal.BucketsDeleteParams{
					ID:      ctx.String("id"),
					Name:    ctx.String("name"),
					OrgID:   ctx.String("org-id"),
					OrgName: ctx.String("org"),
				})
			},
		},
		{
			Name:    "list",
			Usage:   "List buckets",
			Aliases: []string{"find", "ls"},
			Flags: append(
				commonFlags,
				&cli.StringFlag{
					Name:    "name",
					Usage:   "The bucket name",
					Aliases: []string{"n"},
					EnvVars: []string{"INFLUX_BUCKET_NAME"},
				},
				&cli.StringFlag{
					Name:    "org-id",
					Usage:   "The ID of the organization",
					EnvVars: []string{"INFLUX_ORG_ID"},
				},
				&cli.StringFlag{
					Name:    "org",
					Usage:   "The name of the organization",
					Aliases: []string{"o"},
					EnvVars: []string{"INFLUX_ORG"},
				},
				&cli.StringFlag{
					Name:    "id",
					Usage:   "The bucket ID",
					Aliases: []string{"i"},
				},
			),
			Action: func(ctx *cli.Context) error {
				cli, err := newCli(ctx)
				if err != nil {
					return err
				}
				client, err := newApiClient(ctx, cli, true)
				if err != nil {
					return err
				}
				return cli.BucketsList(standardCtx(ctx), client.BucketsApi, &internal.BucketsListParams{
					ID:      ctx.String("id"),
					Name:    ctx.String("name"),
					OrgID:   ctx.String("org-id"),
					OrgName: ctx.String("org"),
				})
			},
		},
		{
			Name:  "update",
			Usage: "Update bucket",
			Flags: append(
				commonFlags,
				&cli.StringFlag{
					Name:    "name",
					Usage:   "New name to set on the bucket",
					Aliases: []string{"n"},
					EnvVars: []string{"INFLUX_BUCKET_NAME"},
				},
				&cli.StringFlag{
					Name:     "id",
					Usage:    "The bucket ID",
					Aliases:  []string{"i"},
					Required: true,
				},
				&cli.StringFlag{
					Name:    "description",
					Usage:   "New description to set on the bucket",
					Aliases: []string{"d"},
				},
				&cli.StringFlag{
					Name:    "retention",
					Usage:   "New retention duration to set on the bucket, or 0 for infinite",
					Aliases: []string{"r"},
				},
				&cli.StringFlag{
					Name:  "shard-group-duration",
					Usage: "New shard group duration to set on the bucket, or 0 to have the server calculate a value",
				},
			),
			Action: func(ctx *cli.Context) error {
				cli, err := newCli(ctx)
				if err != nil {
					return err
				}
				client, err := newApiClient(ctx, cli, true)
				if err != nil {
					return err
				}
				return cli.BucketsUpdate(standardCtx(ctx), client.BucketsApi, &internal.BucketsUpdateParams{
					ID:                 ctx.String("id"),
					Name:               ctx.String("name"),
					Description:        ctx.String("description"),
					Retention:          ctx.String("retention"),
					ShardGroupDuration: ctx.String("shard-group-duration"),
				})
			},
		},
	},
}
