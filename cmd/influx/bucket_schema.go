package main

import (
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/bucket_schema"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/urfave/cli/v2"
)

func withBucketSchemaClient() cli.BeforeFunc {
	return middleware.WithBeforeFns(
		withCli(),
		withApi(true),
		func(ctx *cli.Context) error {
			client := getAPI(ctx)
			ctx.App.Metadata["measurement_schema"] = bucket_schema.Client{
				BucketsApi:       client.BucketsApi,
				BucketSchemasApi: client.BucketSchemasApi.OnlyCloud(),
				CLI:              getCLI(ctx),
			}
			return nil
		})
}

func getBucketSchemaClient(ctx *cli.Context) bucket_schema.Client {
	i, ok := ctx.App.Metadata["measurement_schema"].(bucket_schema.Client)
	if !ok {
		panic("missing measurement schema clients")
	}
	return i
}

func newBucketSchemaCmd() *cli.Command {
	return &cli.Command{
		Name:  "bucket-schema",
		Usage: "Bucket schema management commands",
		Subcommands: []*cli.Command{
			newBucketSchemaCreateCmd(),
			newBucketSchemaUpdateCmd(),
			newBucketSchemaListCmd(),
		},
	}
}

func newBucketSchemaCreateCmd() *cli.Command {
	var params struct {
		clients.OrgBucketParams
		Name           string
		ColumnsFile    string
		ColumnsFormat  bucket_schema.ColumnsFormat
		ExtendedOutput bool
	}
	return &cli.Command{
		Name:   "create",
		Usage:  "Create a measurement schema for a bucket",
		Before: withBucketSchemaClient(),
		Flags: append(
			commonFlags(),
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
			return getBucketSchemaClient(ctx).
				Create(ctx.Context, bucket_schema.CreateParams{
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

func newBucketSchemaUpdateCmd() *cli.Command {
	var params struct {
		clients.OrgBucketParams
		ID             influxid.ID
		Name           string
		ColumnsFile    string
		ColumnsFormat  bucket_schema.ColumnsFormat
		ExtendedOutput bool
	}
	return &cli.Command{
		Name:   "update",
		Usage:  "Update a measurement schema for a bucket",
		Before: withBucketSchemaClient(),
		Flags: append(
			commonFlags(),
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
			return getBucketSchemaClient(ctx).
				Update(ctx.Context, bucket_schema.UpdateParams{
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

func newBucketSchemaListCmd() *cli.Command {
	var params bucket_schema.ListParams
	return &cli.Command{
		Name:   "list",
		Usage:  "List schemas for a bucket",
		Before: withBucketSchemaClient(),
		Flags: append(
			commonFlags(),
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
			return getBucketSchemaClient(ctx).List(ctx.Context, params)
		},
	}
}
