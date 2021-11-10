package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/apply"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/template"
	"github.com/mattn/go-isatty"
	"github.com/urfave/cli"
)

func newApplyCmd() cli.Command {
	var params struct {
		orgParams      clients.OrgParams
		stackId        string
		inPaths        cli.StringSlice
		inUrls         cli.StringSlice
		recursive      bool
		encoding       template.Encoding
		noColor        bool
		noTableBorders bool
		quiet          bool
		force          string
		secrets        cli.StringSlice
		envVars        cli.StringSlice
		filters        cli.StringSlice
	}
	return cli.Command{
		Name:  "apply",
		Usage: "Apply a template to manage resources",
		Description: `The apply command applies InfluxDB template(s). Use the command to create new
resources via a declarative template. The apply command can consume templates
via file(s), url(s), stdin, or any combination of the 3. Each run of the apply
command ensures that all templates applied are applied in unison as a transaction.
If any unexpected errors are discovered then all side effects are rolled back.

Examples:
	# Apply a template via a file
	influx apply -f $PATH_TO_TEMPLATE/template.json

	# Apply a stack that has associated templates. In this example the stack has a remote
	# template associated with it.
	influx apply --stack-id $STACK_ID

	# Apply a template associated with a stack. Stacks make template application idempotent.
	influx apply -f $PATH_TO_TEMPLATE/template.json --stack-id $STACK_ID

	# Apply multiple template files together (mix of yaml and json)
	influx apply \
		-f $PATH_TO_TEMPLATE/template_1.json \
		-f $PATH_TO_TEMPLATE/template_2.yml

	# Apply a template from a url
	influx apply -u https://raw.githubusercontent.com/influxdata/community-templates/master/docker/docker.yml

	# Apply a template from STDIN
	cat $TEMPLATE.json | influx apply --encoding json

	# Applying a directory of templates, takes everything from provided directory
	influx apply -f $PATH_TO_TEMPLATE_DIR

	# Applying a directory of templates, recursively descending into child directories
	influx apply -R -f $PATH_TO_TEMPLATE_DIR

	# Applying directories from many sources, file and URL
	influx apply -f $PATH_TO_TEMPLATE/template.yml -f $URL_TO_TEMPLATE

	# Applying a template with actions to skip resources applied. The
	# following example skips all buckets and the dashboard whose 
	# metadata.name field matches the provided $DASHBOARD_TMPL_NAME.
	# format for filters:
	#	--filter=kind=Bucket
	#	--filter=resource=Label:$Label_TMPL_NAME
	influx apply \
		-f $PATH_TO_TEMPLATE/template.yml \
		--filter kind=Bucket \
		--filter resource=Dashboard:$DASHBOARD_TMPL_NAME

For information about finding and using InfluxDB templates, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/apply/.

For more templates created by the community, see
https://github.com/influxdata/community-templates.
`,
		Flags: append(
			append(commonFlags(), getOrgFlags(&params.orgParams)...),
			&cli.StringFlag{
				Name:        "stack-id",
				Usage:       "Stack ID to associate with template application",
				Destination: &params.stackId,
			},
			&cli.StringSliceFlag{
				Name:      "file, f",
				Usage:     "Path to template file; Supports file paths or (deprecated) HTTP(S) URLs",
				TakesFile: true,
				Value:     &params.inPaths,
			},
			&cli.StringSliceFlag{
				Name:  "template-url, u",
				Usage: "HTTP(S) URL to template file",
				Value: &params.inUrls,
			},
			&cli.BoolFlag{
				Name:        "recurse, R",
				Usage:       "Process the directory used in -f, --file recursively. Useful when you want to manage related templates organized within the same directory.",
				Destination: &params.recursive,
			},
			&cli.GenericFlag{
				Name:  "encoding, e",
				Usage: "Encoding for the input stream. If a file is provided will gather encoding type from file extension. If extension provided will override.",
				Value: &params.encoding,
			},
			&cli.BoolFlag{
				Name:        "disable-color",
				Usage:       "Disable color in output",
				Destination: &params.noColor,
			},
			&cli.BoolFlag{
				Name:        "disable-table-borders",
				Usage:       "Disable table borders",
				Destination: &params.noTableBorders,
			},
			&cli.BoolFlag{
				Name:        "quiet, q",
				Usage:       "Disable output printing",
				Destination: &params.quiet,
			},
			&cli.StringFlag{
				Name:        "force",
				Usage:       "Set to 'true' to skip confirmation before applying changes. Set to 'conflict' to skip confirmation and overwrite existing resources",
				Destination: &params.force,
			},
			&cli.StringSliceFlag{
				Name:  "secret",
				Usage: "Secrets to provide alongside the template; format --secret SECRET_KEY=SECRET_VALUE --secret SECRET_KEY_2=SECRET_VALUE_2",
				Value: &params.secrets,
			},
			&cli.StringSliceFlag{
				Name:  "env-ref",
				Usage: "Environment references to provide alongside the template; format --env-ref REF_KEY=REF_VALUE --env-ref REF_KEY_2=REF_VALUE_2",
				Value: &params.envVars,
			},
			&cli.StringSliceFlag{
				Name:  "filter",
				Usage: "Resources to skip when applying the template. Filter out by `kind` or by `resource`",
				Value: &params.filters,
			},
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			parsedParams := apply.Params{
				OrgParams:          params.orgParams,
				StackId:            params.stackId,
				Recursive:          params.recursive,
				Quiet:              params.quiet,
				RenderTableBorders: !params.noTableBorders,
				RenderTableColors:  !params.noColor,
				Secrets:            make(map[string]string, len(params.secrets.Value())),
				EnvVars:            make(map[string]string, len(params.envVars.Value())),
				Filters:            make([]apply.ResourceFilter, len(params.filters.Value())),
			}

			// Collect all the sources the CLI needs to read templates from.
			var deprecationShown bool
			for _, in := range params.inPaths.Value() {
				// Heuristic to guess what's a URL and what's a local file.
				// TODO: Remove this once we stop supporting URLs in the --file arg.
				u, err := url.Parse(in)
				if err != nil {
					return fmt.Errorf("failed to parse input path %q: %w", in, err)
				}
				if strings.HasPrefix(u.Scheme, "http") {
					if !deprecationShown {
						log.Println("WARN: Passing URLs via -f/--file is deprecated, please use -u/--template-url instead")
						deprecationShown = true
					}
					parsedParams.Sources = append(parsedParams.Sources, template.SourceFromURL(u, params.encoding))
				} else {
					fileSources, err := template.SourcesFromPath(in, params.recursive, params.encoding)
					if err != nil {
						return err
					}
					parsedParams.Sources = append(parsedParams.Sources, fileSources...)
				}
			}
			for _, in := range params.inUrls.Value() {
				u, err := url.Parse(in)
				if err != nil {
					return fmt.Errorf("failed to parse input URL %q: %w", in, err)
				}
				parsedParams.Sources = append(parsedParams.Sources, template.SourceFromURL(u, params.encoding))
			}
			if !isatty.IsTerminal(os.Stdin.Fd()) {
				parsedParams.Sources = append(parsedParams.Sources, template.SourceFromReader(os.Stdin, params.encoding))
			}

			// Parse env and secret values.
			for _, e := range params.envVars.Value() {
				pieces := strings.Split(e, "=")
				if len(pieces) != 2 {
					return fmt.Errorf("env-ref %q has invalid format, must be `REF_KEY=REF_VALUE`", e)
				}
				parsedParams.EnvVars[pieces[0]] = pieces[1]
			}
			for _, s := range params.secrets.Value() {
				pieces := strings.Split(s, "=")
				if len(pieces) != 2 {
					return fmt.Errorf("secret %q has invalid format, must be `SECRET_KEY=SECRET_VALUE`", s)
				}
				parsedParams.Secrets[pieces[0]] = pieces[1]
			}

			// Parse filters.
			for i, f := range params.filters.Value() {
				pieces := strings.Split(f, "=")
				if len(pieces) != 2 {
					return fmt.Errorf("filter %q has invalid format, expected `resource=KIND:NAME` or `kind=KIND`", f)
				}
				key, val := pieces[0], pieces[1]
				switch strings.ToLower(key) {
				case "kind":
					parsedParams.Filters[i] = apply.ResourceFilter{Kind: val}
				case "resource":
					valPieces := strings.Split(val, ":")
					if len(valPieces) != 2 {
						return fmt.Errorf("resource filter %q has invalid format, expected `resource=KIND:NAME", val)
					}
					kind, name := valPieces[0], valPieces[1]
					parsedParams.Filters[i] = apply.ResourceFilter{Kind: kind, Name: &name}
				default:
					return fmt.Errorf("invalid filter type %q, supported values are `resource` and `kind`", key)
				}
			}

			// Parse our strange way of passing 'force'
			switch params.force {
			case "conflict":
				parsedParams.Force = true
				parsedParams.OverwriteConflicts = true
			case "true":
				parsedParams.Force = true
			default:
			}

			api := getAPI(ctx)
			client := apply.Client{
				CLI:              getCLI(ctx),
				TemplatesApi:     api.TemplatesApi,
				OrganizationsApi: api.OrganizationsApi,
			}

			return client.Apply(getContext(ctx), &parsedParams)
		},
	}
}
