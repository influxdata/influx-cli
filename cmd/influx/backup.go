package main

import (
	"compress/gzip"
	"errors"
	"fmt"

	"github.com/influxdata/influx-cli/v2/clients/backup"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

type gzipCompressionLevel int

func (cl *gzipCompressionLevel) Set(v string) error {
	switch v {
	case "default":
		*cl = gzipCompressionLevel(gzip.DefaultCompression)
	case "full":
		*cl = gzipCompressionLevel(gzip.BestCompression)
	case "speedy":
		*cl = gzipCompressionLevel(gzip.BestSpeed)
	default:
		return fmt.Errorf("unknown compression level: %q, required: [default, full, speedy]", v)
	}
	return nil
}

func (cl gzipCompressionLevel) String() string {
	switch int(cl) {
	case gzip.BestCompression:
		return "full"
	case gzip.BestSpeed:
		return "speedy"
	default:
		return "default"
	}
}

func newBackupCmd() cli.Command {
	var params backup.Params
	// Default to gzipping local files.
	params.Compression = br.GzipCompression

	var compressionLevel gzipCompressionLevel

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
				Usage: "Compression to use for local backup files, either 'none' or 'gzip'",
				Value: &params.Compression,
			},
			&cli.GenericFlag{
				Name:  "gzip-compression-level",
				Usage: "The level of gzip compression for backup files: 'default', 'full' (best compression), or 'speedy' (fastest)",
				Value: &compressionLevel,
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
			params.GzipCompressionLevel = int(compressionLevel)

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
