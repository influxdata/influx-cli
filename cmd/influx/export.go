package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients/export"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"
)

func newExportCmd() *cli.Command {
	var params struct {
		out            string
		stackId        string
		resourceType   string
		bucketIds      string
		bucketNames    string
		checkIds       string
		checkNames     string
		dashboardIds   string
		dashboardNames string
		endpointIds    string
		endpointNames  string
		labelIds       string
		labelNames     string
		ruleIds        string
		ruleNames      string
		taskIds        string
		taskNames      string
		telegrafIds    string
		telegrafNames  string
		variableIds    string
		variableNames  string
	}
	return &cli.Command{
		Name:  "export",
		Usage: "Export existing resources as a template",
		Description: `The export command provides a mechanism to export existing resources to a
template. Each template resource kind is supported via flags.

Examples:
	# export buckets by ID
	influx export --buckets=$ID1,$ID2,$ID3

	# export buckets, labels, and dashboards by ID
	influx export \
		--buckets=$BID1,$BID2,$BID3 \
		--labels=$LID1,$LID2,$LID3 \
		--dashboards=$DID1,$DID2,$DID3

	# export all resources for a stack
	influx export --stack-id $STACK_ID

	# export a stack with resources not associated with the stack
	influx export --stack-id $STACK_ID --buckets $BUCKET_ID

All of the resources are supported via the examples provided above. Provide the
resource flag and then provide the IDs.

For information about exporting InfluxDB templates, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/export/`,
		Subcommands: []*cli.Command{
			newExportAllCmd(),
			newExportStackCmd(),
		},
		Flags: append(
			commonFlagsNoPrint(),
			&cli.StringFlag{
				Name:        "file",
				Usage:       "Output file for created template; defaults to std out if no file provided; the extension of provided file (.yml/.json) will dictate encoding",
				Aliases:     []string{"f"},
				Destination: &params.out,
			},
			&cli.StringFlag{
				Name:        "stack-id",
				Usage:       "ID for stack to include in export",
				Destination: &params.stackId,
			},
			&cli.StringFlag{
				Name:        "resource-type",
				Usage:       "If specified, strings on stdin/positional args will be treated as IDs of the given type",
				Destination: &params.resourceType,
			},
			&cli.StringFlag{
				Name:        "buckets",
				Usage:       "List of bucket ids comma separated",
				Destination: &params.bucketIds,
			},
			&cli.StringFlag{
				Name:        "checks",
				Usage:       "List of check ids comma separated",
				Destination: &params.checkIds,
			},
			&cli.StringFlag{
				Name:        "dashboards",
				Usage:       "List of dashboard ids comma separated",
				Destination: &params.dashboardIds,
			},
			&cli.StringFlag{
				Name:        "endpoints",
				Usage:       "List of notification endpoint ids comma separated",
				Destination: &params.endpointIds,
			},
			&cli.StringFlag{
				Name:        "labels",
				Usage:       "List of label ids comma separated",
				Destination: &params.labelIds,
			},
			&cli.StringFlag{
				Name:        "rules",
				Usage:       "List of notification rule ids comma separated",
				Destination: &params.ruleIds,
			},
			&cli.StringFlag{
				Name:        "tasks",
				Usage:       "List of task ids comma separated",
				Destination: &params.taskIds,
			},
			&cli.StringFlag{
				Name:        "telegraf-configs",
				Usage:       "List of telegraf config ids comma separated",
				Destination: &params.telegrafIds,
			},
			&cli.StringFlag{
				Name:        "variables",
				Usage:       "List of variable ids comma separated",
				Destination: &params.variableIds,
			},
			&cli.StringFlag{
				Name:        "bucket-names",
				Usage:       "List of bucket names comma separated",
				Destination: &params.bucketNames,
			},
			&cli.StringFlag{
				Name:        "check-names",
				Usage:       "List of check names comma separated",
				Destination: &params.checkNames,
			},
			&cli.StringFlag{
				Name:        "dashboard-names",
				Usage:       "List of dashboard names comma separated",
				Destination: &params.dashboardNames,
			},
			&cli.StringFlag{
				Name:        "endpoint-names",
				Usage:       "List of notification endpoint names comma separated",
				Destination: &params.endpointNames,
			},
			&cli.StringFlag{
				Name:        "label-names",
				Usage:       "List of label names comma separated",
				Destination: &params.labelNames,
			},
			&cli.StringFlag{
				Name:        "rule-names",
				Usage:       "List of notification rule names comma separated",
				Destination: &params.ruleNames,
			},
			&cli.StringFlag{
				Name:        "task-names",
				Usage:       "List of task names comma separated",
				Destination: &params.taskNames,
			},
			&cli.StringFlag{
				Name:        "telegraf-config-names",
				Usage:       "List of telegraf config names comma separated",
				Destination: &params.telegrafNames,
			},
			&cli.StringFlag{
				Name:        "variable-names",
				Usage:       "List of variable names comma separated",
				Destination: &params.variableNames,
			},
		),
		ArgsUsage: "[resource-id]...",
		Before:    middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			parsedParams := export.Params{
				StackId: params.stackId,
				IdsPerType: map[string][]string{
					"Bucket":               splitNonEmpty(params.bucketIds),
					"Check":                splitNonEmpty(params.checkIds),
					"Dashboard":            splitNonEmpty(params.dashboardIds),
					"NotificationEndpoint": splitNonEmpty(params.endpointIds),
					"Label":                splitNonEmpty(params.labelIds),
					"NotificationRule":     splitNonEmpty(params.ruleIds),
					"Task":                 splitNonEmpty(params.taskIds),
					"Telegraf":             splitNonEmpty(params.telegrafIds),
					"Variable":             splitNonEmpty(params.variableIds),
				},
				NamesPerType: map[string][]string{
					"Bucket":               splitNonEmpty(params.bucketNames),
					"Check":                splitNonEmpty(params.checkNames),
					"Dashboard":            splitNonEmpty(params.dashboardNames),
					"NotificationEndpoint": splitNonEmpty(params.endpointNames),
					"Label":                splitNonEmpty(params.labelNames),
					"NotificationRule":     splitNonEmpty(params.ruleNames),
					"Task":                 splitNonEmpty(params.taskNames),
					"Telegraf":             splitNonEmpty(params.telegrafNames),
					"Variable":             splitNonEmpty(params.variableNames),
				},
			}

			outParams, closer, err := parseOutParams(params.out)
			if closer != nil {
				defer closer()
			}
			if err != nil {
				return err
			}
			parsedParams.OutParams = outParams

			if params.resourceType != "" {
				ids := ctx.Args().Slice()

				// Read any IDs from stdin.
				// !IsTerminal detects when some other process is piping into this command.
				if !isatty.IsTerminal(os.Stdin.Fd()) {
					inBytes, err := io.ReadAll(os.Stdin)
					if err != nil {
						return fmt.Errorf("failed to read args from std in: %w", err)
					}
					ids = append(ids, strings.Fields(string(inBytes))...)
				}

				if _, ok := parsedParams.IdsPerType[params.resourceType]; !ok {
					parsedParams.IdsPerType[params.resourceType] = []string{}
				}
				parsedParams.IdsPerType[params.resourceType] = append(parsedParams.IdsPerType[params.resourceType], ids...)

				if _, ok := parsedParams.NamesPerType[params.resourceType]; !ok {
					parsedParams.NamesPerType[params.resourceType] = []string{}
				}
				parsedParams.NamesPerType[params.resourceType] = append(parsedParams.NamesPerType[params.resourceType], ids...)
			} else if ctx.NArg() > 0 {
				return fmt.Errorf("must specify --resource-type when passing IDs as args")
			}

			client := export.Client{
				CLI:          getCLI(ctx),
				TemplatesApi: getAPI(ctx).TemplatesApi,
			}
			return client.Export(ctx.Context, &parsedParams)
		},
	}
}

func newExportAllCmd() *cli.Command {
	var params struct {
		out     string
		orgId   string
		orgName string
		filters cli.StringSlice
	}
	return &cli.Command{
		Name:  "all",
		Usage: "Export all existing resources for an organization as a template",
		Description: `The export all command will export all resources for an organization. The
command also provides a mechanism to filter by label name or resource kind.

Examples:
	# Export all resources for an organization
	influx export all --org $ORG_NAME

	# Export all bucket resources
	influx export all --org $ORG_NAME --filter=kind=Bucket

	# Export all resources associated with label Foo
	influx export all --org $ORG_NAME --filter=labelName=Foo

	# Export all bucket resources and filter by label Foo
	influx export all --org $ORG_NAME \
		--filter kind=Bucket \
		--filter labelName=Foo

	# Export all bucket or dashboard resources and filter by label Foo.
	# note: like filters are unioned and filter types are intersections.
	#		This example will export a resource if it is a dashboard or
	#		bucket and has an associated label of Foo.
	influx export all --org $ORG_NAME \
		--filter kind=Bucket \
		--filter kind=Dashboard \
		--filter labelName=Foo

For information about exporting InfluxDB templates, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/export/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/export/all/
`,
		Flags: append(
			commonFlagsNoPrint(),
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVars:     []string{"INFLUX_ORG_ID"},
				Destination: &params.orgId,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &params.orgName,
			},
			&cli.StringFlag{
				Name:        "file",
				Usage:       "Output file for created template; defaults to std out if no file provided; the extension of provided file (.yml/.json) will dictate encoding",
				Aliases:     []string{"f"},
				Destination: &params.out,
			},
			&cli.StringSliceFlag{
				Name:        "filter",
				Usage:       "Filter exported resources by labelName or resourceKind (format: labelName=example)",
				Destination: &params.filters,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			parsedParams := export.AllParams{
				OrgId:   params.orgId,
				OrgName: params.orgName,
			}

			for _, filter := range params.filters.Value() {
				components := strings.Split(filter, "=")
				if len(components) != 2 {
					return fmt.Errorf("invalid filter %q, must have format `type=example`", filter)
				}
				switch key, val := components[0], components[1]; key {
				case "labelName":
					parsedParams.LabelFilters = append(parsedParams.LabelFilters, val)
				case "kind", "resourceKind":
					parsedParams.KindFilters = append(parsedParams.KindFilters, val)
				default:
					return fmt.Errorf("invalid filter provided %q; filter must be 1 in [labelName, resourceKind]", filter)
				}
			}

			outParams, closer, err := parseOutParams(params.out)
			if closer != nil {
				defer closer()
			}
			if err != nil {
				return err
			}
			parsedParams.OutParams = outParams

			apiClient := getAPI(ctx)
			client := export.Client{
				CLI:              getCLI(ctx),
				TemplatesApi:     apiClient.TemplatesApi,
				OrganizationsApi: apiClient.OrganizationsApi,
			}
			return client.ExportAll(ctx.Context, &parsedParams)
		},
	}
}

func newExportStackCmd() *cli.Command {
	var params struct {
		out string
	}

	return &cli.Command{
		Name:  "stack",
		Usage: "Export all resources associated with a stack as a template",
		Description: `The influx export stack command exports all resources 
associated with a stack as a template. All metadata.name fields remain the same.

Example:
	# Export a stack as a template
	influx export stack $STACK_ID

For information about exporting InfluxDB templates, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/export/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/export/stack/
`,
		Flags: append(
			commonFlagsNoPrint(),
			&cli.StringFlag{
				Name:        "file",
				Usage:       "Output file for created template; defaults to std out if no file provided; the extension of provided file (.yml/.json) will dictate encoding",
				Aliases:     []string{"f"},
				Destination: &params.out,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() != 1 {
				return fmt.Errorf("incorrect number of arguments, expected one for <stack-id>")
			}

			parsedParams := export.StackParams{
				StackId: ctx.Args().Get(0),
			}

			outParams, closer, err := parseOutParams(params.out)
			if closer != nil {
				defer closer()
			}
			if err != nil {
				return err
			}
			parsedParams.OutParams = outParams

			apiClient := getAPI(ctx)
			client := export.Client{
				CLI:              getCLI(ctx),
				TemplatesApi:     apiClient.TemplatesApi,
				OrganizationsApi: apiClient.OrganizationsApi,
			}
			return client.ExportStack(ctx.Context, &parsedParams)
		},
	}
}

func splitNonEmpty(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

func parseOutParams(outPath string) (export.OutParams, func(), error) {
	if outPath == "" {
		return export.OutParams{Out: os.Stdout, Encoding: export.YamlEncoding}, nil, nil
	}

	f, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return export.OutParams{}, nil, fmt.Errorf("failed to open output path %q: %w", outPath, err)
	}
	params := export.OutParams{Out: f}
	switch filepath.Ext(outPath) {
	case ".json":
		params.Encoding = export.JsonEncoding
	default:
		params.Encoding = export.YamlEncoding
	}

	return params, func() { _ = f.Close() }, nil
}
