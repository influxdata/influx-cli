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
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

type writeParams struct {
	Files       cli.StringSlice
	URLs        cli.StringSlice
	Format      linereader.InputFormat
	Compression linereader.InputCompression
	Encoding    string

	// CSV-specific options.
	Headers                    cli.StringSlice
	SkipRowOnError             bool
	SkipHeader                 int
	IgnoreDataTypeInColumnName bool
	Debug                      bool

	ErrorsFile    string
	MaxLineLength int
	RateLimit     throttler.BytesPerSec

	internal.WriteParams
}

func (p *writeParams) makeLineReader(args []string, errorOut io.Writer) *linereader.MultiInputLineReader {
	return &linereader.MultiInputLineReader{
		StdIn:                      os.Stdin,
		HttpClient:                 http.DefaultClient,
		ErrorOut:                   errorOut,
		Args:                       args,
		Files:                      p.Files.Value(),
		URLs:                       p.URLs.Value(),
		Format:                     p.Format,
		Compression:                p.Compression,
		Encoding:                   p.Encoding,
		Headers:                    p.Headers.Value(),
		SkipRowOnError:             p.SkipRowOnError,
		SkipHeader:                 p.SkipHeader,
		IgnoreDataTypeInColumnName: p.IgnoreDataTypeInColumnName,
		Debug:                      p.Debug,
	}
}

func (p *writeParams) makeErrorFile() (*os.File, error) {
	if p.ErrorsFile == "" {
		return nil, nil
	}
	errorFile, err := os.Open(p.ErrorsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open errors-file: %w", err)
	}
	return errorFile, nil
}

func (p *writeParams) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "bucket-id",
			Usage:       "The ID of destination bucket",
			EnvVars:     []string{"INFLUX_BUCKET_ID"},
			Destination: &p.BucketID,
		},
		&cli.StringFlag{
			Name:        "bucket",
			Usage:       "The name of destination bucket",
			Aliases:     []string{"b"},
			EnvVars:     []string{"INFLUX_BUCKET_NAME"},
			Destination: &p.BucketName,
		},
		&cli.StringFlag{
			Name:        "org-id",
			Usage:       "The ID of the organization",
			EnvVars:     []string{"INFLUX_ORG_ID"},
			Destination: &p.OrgID,
		},
		&cli.StringFlag{
			Name:        "org",
			Usage:       "The name of the organization",
			Aliases:     []string{"o"},
			EnvVars:     []string{"INFLUX_ORG"},
			Destination: &p.OrgName,
		},
		&cli.GenericFlag{
			Name:    "precision",
			Usage:   "Precision of the timestamps of the lines",
			Aliases: []string{"p"},
			EnvVars: []string{"INFLUX_PRECISION"},
			Value:   &p.Precision,
		},
		&cli.GenericFlag{
			Name:        "format",
			Usage:       "Input format, either 'lp' (Line Protocol) or 'csv' (Comma Separated Values)",
			DefaultText: "'lp' unless '.csv' extension",
			Value:       &p.Format,
		},
		&cli.StringSliceFlag{
			Name:        "header",
			Usage:       "Header prepends lines to input data",
			Destination: &p.Headers,
		},
		&cli.StringSliceFlag{
			Name:        "file",
			Usage:       "The path to the file to import",
			Aliases:     []string{"f"},
			TakesFile:   true,
			Destination: &p.Files,
		},
		&cli.StringSliceFlag{
			Name:        "url",
			Usage:       "The URL to import data from",
			Aliases:     []string{"u"},
			Destination: &p.URLs,
		},
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Log CSV columns to stderr before reading data rows",
			Destination: &p.Debug,
		},
		&cli.BoolFlag{
			Name:        "skipRowOnError",
			Usage:       "Log CSV data errors to stderr and continue with CSV processing",
			Destination: &p.SkipRowOnError,
		},
		// NOTE: The old CLI allowed this flag to be used as an int _or_ a bool, with the bool form being
		// short-hand for N=1. urfave/cli isn't that flexible.
		&cli.IntFlag{
			Name:        "skipHeader",
			Usage:       "Skip the first <n> rows from input data",
			Destination: &p.SkipHeader,
		},
		&cli.IntFlag{
			Name:        "max-line-length",
			Usage:       "Specifies the maximum number of bytes that can be read for a single line",
			Value:       16_000_000,
			Destination: &p.MaxLineLength,
		},
		&cli.BoolFlag{
			Name:        "xIgnoreDataTypeInColumnName",
			Usage:       "Ignores dataType which could be specified after ':' in column name",
			Hidden:      true,
			Destination: &p.IgnoreDataTypeInColumnName,
		},
		&cli.StringFlag{
			Name:        "encoding",
			Usage:       "Character encoding of input files or stdin",
			Value:       "UTF-8",
			Destination: &p.Encoding,
		},
		&cli.StringFlag{
			Name:        "errors-file",
			Usage:       "The path to the file to write rejected rows to",
			TakesFile:   true,
			Destination: &p.ErrorsFile,
		},
		&cli.GenericFlag{
			Name:        "rate-limit",
			Usage:       `Throttles write, examples: "5 MB / 5 min" , "17kBs"`,
			DefaultText: "no throttling",
			Value:       &p.RateLimit,
		},
		&cli.GenericFlag{
			Name:        "compression",
			Usage:       "Input compression, either 'none' or 'gzip'",
			DefaultText: "'none' unless an input has a '.gz' extension",
			Value:       &p.Compression,
		},
	}
}

func newWriteCmd() *cli.Command {
	params := writeParams{
		WriteParams: internal.WriteParams{
			Precision: api.WRITEPRECISION_NS,
		},
	}
	return &cli.Command{
		Name:        "write",
		Usage:       "Write points to InfluxDB",
		Description: "Write data to InfluxDB via stdin, or add an entire file specified with the -f flag",
		Before:      middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags:       append(commonFlagsNoPrint, params.Flags()...),
		Action: func(ctx *cli.Context) error {
			errorFile, err := params.makeErrorFile()
			if err != nil {
				return err
			}
			defer func() { _ = errorFile.Close() }()

			client := getAPI(ctx)
			writeClients := &internal.WriteClients{
				Client:    client.WriteApi,
				Reader:    params.makeLineReader(ctx.Args().Slice(), errorFile),
				Throttler: throttler.NewThrottler(params.RateLimit),
				Writer: &batcher.BufferBatcher{
					MaxFlushBytes:    batcher.DefaultMaxBytes,
					MaxFlushInterval: batcher.DefaultInterval,
					MaxLineLength:    params.MaxLineLength,
				},
			}

			cli := getCLI(ctx)
			return cli.Write(ctx.Context, writeClients, &params.WriteParams)
		},
		Subcommands: []*cli.Command{
			newWriteDryRun(),
		},
	}
}

func newWriteDryRun() *cli.Command {
	params := writeParams{
		WriteParams: internal.WriteParams{
			Precision: api.WRITEPRECISION_NS,
		},
	}

	return &cli.Command{
		Name:        "dryrun",
		Usage:       "Write to stdout instead of InfluxDB",
		Description: "Write protocol lines to stdout instead of InfluxDB. Troubleshoot conversion from CSV to line protocol",
		Before:      middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags:       append(commonFlagsNoPrint, params.Flags()...),
		Action: func(ctx *cli.Context) error {
			errorFile, err := params.makeErrorFile()
			if err != nil {
				return err
			}
			defer func() { _ = errorFile.Close() }()

			cli := getCLI(ctx)
			return cli.WriteDryRun(ctx.Context, params.makeLineReader(ctx.Args().Slice(), errorFile))
		},
	}
}
