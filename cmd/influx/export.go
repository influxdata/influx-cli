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
		resourceType   export.ResourceType
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
			&cli.GenericFlag{
				Name:  "resource-type",
				Usage: "If specified, strings on stdin/positional args will be treated as IDs of the given type",
				Value: &params.resourceType,
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
				StackId:        params.stackId,
				BucketIds:      splitNonEmpty(params.bucketIds),
				BucketNames:    splitNonEmpty(params.bucketNames),
				CheckIds:       splitNonEmpty(params.checkIds),
				CheckNames:     splitNonEmpty(params.checkNames),
				DashboardIds:   splitNonEmpty(params.dashboardIds),
				DashboardNames: splitNonEmpty(params.dashboardNames),
				EndpointIds:    splitNonEmpty(params.endpointIds),
				EndpointNames:  splitNonEmpty(params.endpointNames),
				LabelIds:       splitNonEmpty(params.labelIds),
				LabelNames:     splitNonEmpty(params.labelNames),
				RuleIds:        splitNonEmpty(params.ruleIds),
				RuleNames:      splitNonEmpty(params.ruleNames),
				TaskIds:        splitNonEmpty(params.taskIds),
				TaskNames:      splitNonEmpty(params.taskNames),
				TelegrafIds:    splitNonEmpty(params.telegrafIds),
				TelegrafNames:  splitNonEmpty(params.telegrafNames),
				VariableIds:    splitNonEmpty(params.variableIds),
				VariableNames:  splitNonEmpty(params.variableNames),
			}

			outParams, closer, err := parseOutParams(params.out)
			if closer != nil {
				defer closer()
			}
			if err != nil {
				return err
			}
			parsedParams.OutParams = outParams

			if params.resourceType != export.TypeUnset {
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

				switch params.resourceType {
				case export.TypeBucket:
					parsedParams.BucketIds = append(parsedParams.BucketIds, ids...)
				case export.TypeCheck:
					parsedParams.CheckIds = append(parsedParams.CheckIds, ids...)
				case export.TypeDashboard:
					parsedParams.DashboardIds = append(parsedParams.DashboardIds, ids...)
				case export.TypeLabel:
					parsedParams.LabelIds = append(parsedParams.LabelIds, ids...)
				case export.TypeNotificationEndpoint:
					parsedParams.EndpointIds = append(parsedParams.EndpointIds, ids...)
				case export.TypeNotificationRule:
					parsedParams.RuleIds = append(parsedParams.RuleIds, ids...)
				case export.TypeTask:
					parsedParams.TaskIds = append(parsedParams.TaskIds, ids...)
				case export.TypeTelegraf:
					parsedParams.TelegrafIds = append(parsedParams.TelegrafIds, ids...)
				case export.TypeVariable:
					parsedParams.VariableIds = append(parsedParams.VariableIds, ids...)

				// NOTE: The API doesn't support filtering by these resource subtypes,
				// and instead converts them to the parent type. For example,
				// `--resource-type notificationEndpointHTTP` gets translated to a filter
				// on all notification endpoints on the server-side. I think this was
				// intentional since the 2.0.x CLI didn't expose flags to filter on subtypes,
				// but a bug/oversight in its parsing still allowed the subtypes through
				// when passing IDs over stdin.
				// Instead of allowing the type-filter to be silently converted by the server,
				// we catch the previously-allowed subtypes here and return a (hopefully) useful
				// error suggesting the correct flag to use.
				case export.TypeCheckDeadman, export.TypeCheckThreshold:
					return fmt.Errorf("filtering on resource-type %q is not supported by the API. Use resource-type %q instead", params.resourceType, export.TypeCheck)
				case export.TypeNotificationEndpointHTTP, export.TypeNotificationEndpointPagerDuty, export.TypeNotificationEndpointSlack:
					return fmt.Errorf("filtering on resource-type %q is not supported by the API. Use resource-type %q instead", params.resourceType, export.TypeNotificationEndpoint)

				default:
				}
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
