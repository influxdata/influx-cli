package main

import (
	"errors"

	"github.com/influxdata/influx-cli/v2/clients/backup"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newBackupCmd() cli.Command {
	var params backup.Params
	// Default to gzipping local files.
	params.Compression = br.GzipCompression

	return cli.Command{
		Name:  "backup",
		Usage: "Backup database",
		Description: `Backs up InfluxDB to a directory

Examples:
	# backup all data
	influx backup /path/to/backup
`,
		ArgsUsage: "path",
		Before:    middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			append(commonFlagsNoPrint(), getOrgFlags(&params.OrgParams)...),
			&cli.StringFlag{
				Name:        "bucket-id",
				Usage:       "The ID of the bucket to backup",
				Destination: &params.BucketID,
			},
			&cli.StringFlag{
				Name:        "bucket, b",
				Usage:       "The name of the bucket to backup",
				Destination: &params.BucketName,
			},
			&cli.GenericFlag{
				Name:  "compression",
				Usage: "Compression to use for local backup files on the client-side, either 'none' or 'gzip'. When -gzip-compression-level is set to 'none' this defaults to 'none'",
				Value: &params.Compression,
			},
			&cli.StringFlag{
				Name:        "gzip-compression-level",
				Usage:       "The level of gzip compression for server-side backup: 'default', 'full' (best compression), 'speedy' (fastest), or 'none'",
				Destination: &params.GzipCompressionLevel,
			},
		),
		Action: func(ctx *cli.Context) error {
			if err := checkOrgFlags(&params.OrgParams); err != nil {
				return err
			}
			if ctx.NArg() != 1 {
				return errors.New("backup path must be specified as a single positional argument")
			}
			params.Path = ctx.Args().Get(0)

			// If the user requested no server-side compression and didn't
			// explicitly set local compression, skip local gzip too.
			if params.GzipCompressionLevel == "none" && !ctx.IsSet("compression") {
				params.Compression = br.NoCompression
			}

			api := getAPI(ctx)
			client := backup.Client{
				CLI:       getCLI(ctx),
				BackupApi: api.BackupApi,
				HealthApi: api.HealthApi,
			}
			return client.Backup(getContext(ctx), &params)
		},
	}
}
