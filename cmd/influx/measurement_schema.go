package main

import (
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/cmd/measurement_schema"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/urfave/cli/v2"
)

func withMSClient() cli.BeforeFunc {
	return middleware.WithBeforeFns(
		withCli(),
		withApi(),
		func(ctx *cli.Context) error {
			c := getCLI(ctx)
			client := getAPI(ctx)
			ctx.App.Metadata["measurement_schema"] = measurement_schema.Client{
				BucketApi:        client.BucketsApi,
				BucketSchemasApi: client.BucketSchemasApi,
				CLI:              c,
			}
			return nil
		})
}

func getMSClient(ctx *cli.Context) measurement_schema.Client {
	i, ok := ctx.App.Metadata["measurement_schema"].(measurement_schema.Client)
	if !ok {
		panic("missing measurement schema client")
	}
	return i
}

func newMeasurementSchemaCmd() *cli.Command {
	return &cli.Command{
		Name:  "bucket-schema",
		Usage: "Bucket schema management commands",
		Subcommands: []*cli.Command{
			newMeasurementSchemaCreateCmd(),
			newMeasurementSchemaUpdateCmd(),
			newMeasurementSchemaListCmd(),
		},
	}
}

func newMeasurementSchemaCreateCmd() *cli.Command {
	var params struct {
		internal.OrgBucketParams
		Name           string
		ColumnsFile    string
		ColumnsFormat  measurement_schema.ColumnsFormat
		ExtendedOutput bool
	}
	return &cli.Command{
		Name:   "create",
		Usage:  "Create a measurement schema for a bucket",
		Before: withMSClient(),
		Flags: append(
			commonFlags,
			append(
				getOrgBucketFlags(&params.OrgBucketParams),
				&cli.StringFlag{
					Name:        "name",
					Usage:       "Name of the measurement",
					Destination: &params.Name,
				},
				&cli.StringFlag{
					Name:        "columns-file",
					Usage:       "A path referring to list of column definitions",
					Destination: &params.ColumnsFile,
				},
				&cli.GenericFlag{
					Name:        "columns-format",
					Usage:       "The format of the columns file. \"auto\" will attempt to guess the format.",
					DefaultText: "auto",
					Value:       &params.ColumnsFormat,
				},
				&cli.BoolFlag{
					Name:        "extended-output",
					Usage:       "Print column information for each measurement",
					Aliases:     []string{"x"},
					Destination: &params.ExtendedOutput,
				},
			)...,
		),
		Action: func(ctx *cli.Context) error {
			return getMSClient(ctx).
				Create(ctx.Context, measurement_schema.CreateParams{
					OrgBucketParams: params.OrgBucketParams,
					Name:            params.Name,
					Stdin:           ctx.App.Reader,
					ColumnsFile:     params.ColumnsFile,
					ColumnsFormat:   params.ColumnsFormat,
					ExtendedOutput:  params.ExtendedOutput,
				})
		},
	}
}

func newMeasurementSchemaUpdateCmd() *cli.Command {
	var params struct {
		internal.OrgBucketParams
		ID             influxid.ID
		Name           string
		ColumnsFile    string
		ColumnsFormat  measurement_schema.ColumnsFormat
		ExtendedOutput bool
	}
	return &cli.Command{
		Name:   "update",
		Usage:  "Update a measurement schema for a bucket",
		Before: withMSClient(),
		Flags: append(
			commonFlags,
			append(
				getOrgBucketFlags(&params.OrgBucketParams),
				&cli.GenericFlag{
					Name:  "id",
					Usage: "ID of the measurement",
					Value: &params.ID,
				},
				&cli.StringFlag{
					Name:        "name",
					Usage:       "Name of the measurement",
					Destination: &params.Name,
				},
				&cli.StringFlag{
					Name:        "columns-file",
					Usage:       "A path referring to list of column definitions",
					Destination: &params.ColumnsFile,
				},
				&cli.GenericFlag{
					Name:        "columns-format",
					Usage:       "The format of the columns file. \"auto\" will attempt to guess the format.",
					DefaultText: "auto",
					Value:       &params.ColumnsFormat,
				},
				&cli.BoolFlag{
					Name:        "extended-output",
					Usage:       "Print column information for each measurement",
					Aliases:     []string{"x"},
					Destination: &params.ExtendedOutput,
				},
			)...,
		),
		Action: func(ctx *cli.Context) error {
			return getMSClient(ctx).
				Update(ctx.Context, measurement_schema.UpdateParams{
					OrgBucketParams: params.OrgBucketParams,
					ID:              params.ID.String(),
					Name:            params.Name,
					Stdin:           ctx.App.Reader,
					ColumnsFile:     params.ColumnsFile,
					ColumnsFormat:   params.ColumnsFormat,
					ExtendedOutput:  params.ExtendedOutput,
				})
		},
	}
}

func newMeasurementSchemaListCmd() *cli.Command {
	var params measurement_schema.ListParams
	return &cli.Command{
		Name:   "list",
		Usage:  "List schemas for a bucket",
		Before: withMSClient(),
		Flags: append(
			commonFlags,
			append(
				getOrgBucketFlags(&params.OrgBucketParams),
				&cli.StringFlag{
					Name:        "name",
					Usage:       "Name of single measurement to find",
					Destination: &params.Name,
				},
				&cli.BoolFlag{
					Name:        "extended-output",
					Usage:       "Print column information for each measurement",
					Aliases:     []string{"x"},
					Destination: &params.ExtendedOutput,
				},
			)...,
		),
		Action: func(ctx *cli.Context) error {
			return getMSClient(ctx).List(ctx.Context, params)
		},
	}
}
