package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/batcher"
	"github.com/influxdata/influx-cli/v2/internal/linereader"
	"github.com/influxdata/influx-cli/v2/internal/throttler"
	"github.com/urfave/cli/v2"
)

var writeFlags = append(
	commonFlags,
	&cli.StringFlag{
		Name:    "bucket-id",
		Usage:   "The ID of destination bucket",
		EnvVars: []string{"INFLUX_BUCKET_ID"},
	},
	&cli.StringFlag{
		Name:    "bucket",
		Usage:   "The name of destination bucket",
		Aliases: []string{"b"},
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
		Name:    "precision",
		Usage:   "Precision of the timestamps of the lines",
		Aliases: []string{"p"},
		EnvVars: []string{"INFLUX_PRECISION"},
		Value:   "ns",
	},
	&cli.StringFlag{
		Name:        "format",
		Usage:       "Input format, either 'lp' (Line Protocol) or 'csv' (Comma Separated Values)",
		DefaultText: "'lp' unless '.csv' extension",
	},
	&cli.StringSliceFlag{
		Name:  "header",
		Usage: "Header prepends lines to input data",
	},
	&cli.StringSliceFlag{
		Name:      "file",
		Usage:     "The path to the file to import",
		Aliases:   []string{"f"},
		TakesFile: true,
	},
	&cli.StringSliceFlag{
		Name:    "url",
		Usage:   "The URL to import data from",
		Aliases: []string{"u"},
	},
	&cli.BoolFlag{
		Name:  "debug",
		Usage: "Log CSV columns to stderr before reading data rows",
	},
	&cli.BoolFlag{
		Name:  "skipRowOnError",
		Usage: "Log CSV data errors to stderr and continue with CSV processing",
	},
	// NOTE: The old CLI allowed this flag to be used as an int _or_ a bool, with the bool form being
	// short-hand for N=1. urfave/cli isn't that flexible.
	&cli.IntFlag{
		Name:  "skipHeader",
		Usage: "Skip the first <n> rows from input data",
	},
	&cli.IntFlag{
		Name:  "max-line-length",
		Usage: "Specifies the maximum number of bytes that can be read for a single line",
		Value: 16_000_000,
	},
	&cli.BoolFlag{
		Name:   "xIgnoreDataTypeInColumnName",
		Usage:  "Ignores dataType which could be specified after ':' in column name",
		Hidden: true,
	},
	&cli.StringFlag{
		Name:  "encoding",
		Usage: "Character encoding of input files or stdin",
		Value: "UTF-8",
	},
	&cli.StringFlag{
		Name:      "errors-file",
		Usage:     "The path to the file to write rejected rows to",
		TakesFile: true,
	},
	&cli.StringFlag{
		Name:        "rate-limit",
		Usage:       `Throttles write, examples: "5 MB / 5 min" , "17kBs"`,
		DefaultText: "no throttling",
	},
	&cli.StringFlag{
		Name:        "compression",
		Usage:       "Input compression, either 'none' or 'gzip'",
		DefaultText: "'none' unless an input has a '.gz' extension",
	},
)

var writeCmd = cli.Command{
	Name:        "write",
	Usage:       "Write points to InfluxDB",
	Description: "Write data to InfluxDB via stdin, or add an entire file specified with the -f flag",
	Flags:       writeFlags,
	Action: func(ctx *cli.Context) error {
		format, err := parseFormat(ctx.String("format"))
		if err != nil {
			return err
		}
		compression, err := parseCompression(ctx.String("compression"))
		if err != nil {
			return err
		}
		precision, err := parsePrecision(ctx.String("precision"))
		if err != nil {
			return err
		}

		var errorOut io.Writer
		if ctx.IsSet("errors-file") {
			errorFile, err := os.Open(ctx.String("errors-file"))
			if err != nil {
				return fmt.Errorf("failed to open errors-file: %w", err)
			}
			defer errorFile.Close()
			errorOut = errorFile
		}

		throttler, err := throttler.NewThrottler(ctx.String("rate-limit"))
		if err != nil {
			return err
		}
		cli, err := newCli(ctx)
		if err != nil {
			return err
		}
		client, err := newApiClient(ctx, cli, true)
		if err != nil {
			return err
		}
		writeClients := &internal.WriteClients{
			Client: client.WriteApi,
			Reader: &linereader.MultiInputLineReader{
				StdIn:                      os.Stdin,
				HttpClient:                 http.DefaultClient,
				ErrorOut:                   errorOut,
				Args:                       ctx.Args().Slice(),
				Files:                      ctx.StringSlice("file"),
				URLs:                       ctx.StringSlice("url"),
				Format:                     format,
				Compression:                compression,
				Encoding:                   ctx.String("encoding"),
				Headers:                    ctx.StringSlice("header"),
				SkipRowOnError:             ctx.Bool("skipRowOnError"),
				SkipHeader:                 ctx.Int("skipHeader"),
				IgnoreDataTypeInColumnName: ctx.Bool("xIgnoreDataTypeInColumnName"),
				Debug:                      ctx.Bool("debug"),
			},
			Throttler: throttler,
			Writer: &batcher.BufferBatcher{
				MaxFlushBytes:    batcher.DefaultMaxBytes,
				MaxFlushInterval: batcher.DefaultInterval,
				MaxLineLength:    ctx.Int("max-line-length"),
			},
		}

		return cli.Write(standardCtx(ctx), writeClients, &internal.WriteParams{
			BucketID:   ctx.String("bucket-id"),
			BucketName: ctx.String("bucket"),
			OrgID:      ctx.String("org-id"),
			OrgName:    ctx.String("org"),
			Precision:  precision,
		})
	},
	Subcommands: []*cli.Command{
		{
			Name:        "dryrun",
			Usage:       "Write to stdout instead of InfluxDB",
			Description: "Write protocol lines to stdout instead of InfluxDB. Troubleshoot conversion from CSV to line protocol",
			Flags:       writeFlags,
			Action: func(ctx *cli.Context) error {
				format, err := parseFormat(ctx.String("format"))
				if err != nil {
					return err
				}
				compression, err := parseCompression(ctx.String("compression"))
				if err != nil {
					return err
				}

				var errorOut io.Writer
				if ctx.IsSet("errors-file") {
					errorFile, err := os.Open(ctx.String("errors-file"))
					if err != nil {
						return fmt.Errorf("failed to open errors-file: %w", err)
					}
					defer errorFile.Close()
					errorOut = errorFile
				}

				cli, err := newCli(ctx)
				if err != nil {
					return err
				}
				reader := &linereader.MultiInputLineReader{
					StdIn:                      os.Stdin,
					HttpClient:                 http.DefaultClient,
					ErrorOut:                   errorOut,
					Args:                       ctx.Args().Slice(),
					Files:                      ctx.StringSlice("file"),
					URLs:                       ctx.StringSlice("url"),
					Format:                     format,
					Compression:                compression,
					Encoding:                   ctx.String("encoding"),
					Headers:                    ctx.StringSlice("header"),
					SkipRowOnError:             ctx.Bool("skipRowOnError"),
					SkipHeader:                 ctx.Int("skipHeader"),
					IgnoreDataTypeInColumnName: ctx.Bool("xIgnoreDataTypeInColumnName"),
					Debug:                      ctx.Bool("debug"),
				}

				return cli.WriteDryRun(standardCtx(ctx), reader)
			},
		},
	},
}

func parseFormat(f string) (linereader.InputFormat, error) {
	switch f {
	case "":
		return linereader.InputFormatDerived, nil
	case "lp":
		return linereader.InputFormatLP, nil
	case "csv":
		return linereader.InputFormatCSV, nil
	default:
		return 0, fmt.Errorf("unsupported format: %q", f)
	}
}

func parseCompression(c string) (linereader.InputCompression, error) {
	switch c {
	case "":
		return linereader.InputCompressionDerived, nil
	case "none":
		return linereader.InputCompressionNone, nil
	case "gzip":
		return linereader.InputCompressionGZIP, nil
	default:
		return 0, fmt.Errorf("unsupported compression: %q", c)
	}
}

func parsePrecision(p string) (api.WritePrecision, error) {
	switch p {
	case "ms":
		return api.WRITEPRECISION_MS, nil
	case "s":
		return api.WRITEPRECISION_S, nil
	case "us":
		return api.WRITEPRECISION_US, nil
	case "ns":
		return api.WRITEPRECISION_NS, nil
	default:
		return "", fmt.Errorf("unsupported precision: %q", p)
	}
}
