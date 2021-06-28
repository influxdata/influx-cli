package main

import (
	"errors"

	"github.com/influxdata/influx-cli/v2/clients/restore"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newRestoreCmd() cli.Command {
	var params restore.Params

	return cli.Command{
		Name:  "restore",
		Usage: "Restores a backup directory to InfluxDB",
		Description: `Restore influxdb.

Examples:
	# backup all data
	influx restore /path/to/restore
`,
		ArgsUsage: "path",
		Before:    middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlagsNoPrint(),
			&cli.BoolFlag{
				Name:        "full",
				Usage:       "Fully restore and replace all data on server",
				Destination: &params.Full,
			},
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The original ID of the organization to restore",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The original name of the organization to restore",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.Org,
			},
			&cli.StringFlag{
				Name:        "bucket-id",
				Usage:       "The original ID of the bucket to restore",
				Destination: &params.BucketID,
			},
			&cli.StringFlag{
				Name:        "bucket, b",
				Usage:       "The original name of the bucket to restore",
				Destination: &params.Bucket,
			},
			&cli.StringFlag{
				Name:        "new-bucket",
				Usage:       "New name to use for the restored bucket",
				Destination: &params.NewBucketName,
			},
			&cli.StringFlag{
				Name:        "new-org",
				Usage:       "New name to use for the restored organization",
				Destination: &params.NewOrgName,
			},
		),
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() != 1 {
				return errors.New("restore path must be specified as a single positional argument")
			}
			params.Path = ctx.Args().Get(0)

			if params.Full && (params.Org != "" ||
				params.OrgID != "" ||
				params.Bucket != "" ||
				params.BucketID != "" ||
				params.NewOrgName != "" ||
				params.NewBucketName != "") {
				return errors.New("--full restore cannot be limited to a single org or bucket")
			}

			if params.NewOrgName != "" && params.OrgID == "" && params.Org == "" {
				return errors.New("--org-id or --org must be set to use --new-org")
			}
			if params.NewBucketName != "" && params.BucketID == "" && params.Bucket == "" {
				return errors.New("--bucket-id or --bucket must be set to use --new-bucket")
			}

			api := getAPI(ctx)
			client := restore.Client{
				CLI:              getCLI(ctx),
				RestoreApi:       api.RestoreApi.OnlyOSS(),
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Restore(getContext(ctx), &params)
		},
	}
}
