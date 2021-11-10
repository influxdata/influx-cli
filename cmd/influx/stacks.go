package main

import (
	"fmt"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients/stacks"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/template"
	"github.com/urfave/cli"
)

func newStacksCmd() cli.Command {
	var params stacks.ListParams
	return cli.Command{
		Name:  "stacks",
		Usage: "List stack(s) and associated templates. Subcommands manage stacks.",
		Description: `List stack(s) and associated templates. Subcommands manage stacks.

Examples:
	# list all known stacks
	influx stacks

	# list stacks filtered by stack name
	# output here are stacks that have match at least 1 name provided
	influx stacks --stack-name=$STACK_NAME_1 --stack-name=$STACK_NAME_2

	# list stacks filtered by stack id
	# output here are stacks that have match at least 1 ids provided
	influx stacks --stack-id=$STACK_ID_1 --stack-id=$STACK_ID_2
	
	# list stacks filtered by stack id or stack name
	# output here are stacks that have match the id provided or
	# matches of the name provided
	influx stacks --stack-id=$STACK_ID --stack-name=$STACK_NAME

For information about Stacks and how they integrate with InfluxDB templates, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/stacks/`,
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Flags: append(
			append(commonFlags(), getOrgFlags(&params.OrgParams)...),
			&cli.StringSliceFlag{
				Name:  "stack-id",
				Usage: "Stack ID to filter by",
			},
			&cli.StringSliceFlag{
				Name:  "stack-name",
				Usage: "Stack name to filter by",
			},
		),
		Subcommands: []cli.Command{
			newStacksInitCmd(),
			newStacksRemoveCmd(),
			newStacksUpdateCmd(),
		},
		Action: func(ctx *cli.Context) error {
			params.StackIds = ctx.StringSlice("stack-id")
			params.StackNames = ctx.StringSlice("stack-name")
			api := getAPI(ctx)
			client := stacks.Client{
				CLI:              getCLI(ctx),
				StacksApi:        api.StacksApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.List(getContext(ctx), &params)
		},
	}
}

func newStacksInitCmd() cli.Command {
	var params stacks.InitParams
	return cli.Command{
		Name:  "init",
		Usage: "Initialize a stack",
		Description: `The stack init command creates a new stack to associated templates with. A
stack is used to make applying templates idempotent. When you apply a template
and associate it with a stack, the stack can manage the created/updated resources
from the template back to the platform. This enables a multitude of useful features.
Any associated template urls will be applied when applying templates via a stack.

Examples:
	# Initialize a stack with a name and description
	influx stacks init -n $STACK_NAME -d $STACK_DESCRIPTION

	# Initialize a stack with a name and urls to associate with stack.
	influx stacks init -n $STACK_NAME -u $PATH_TO_TEMPLATE

For information about how stacks work with InfluxDB templates, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/stacks/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/stacks/init/`,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			append(commonFlags(), getOrgFlags(&params.OrgParams)...),
			&cli.StringFlag{
				Name:        "stack-name, n",
				Usage:       "Name given to created stack",
				Destination: &params.Name,
			},
			&cli.StringFlag{
				Name:        "stack-description, d",
				Usage:       "Description given to created stack",
				Destination: &params.Description,
			},
			&cli.StringSliceFlag{
				Name:  "template-url, u",
				Usage: "Template urls to associate with new stack",
			},
		),
		Action: func(ctx *cli.Context) error {
			params.URLs = ctx.StringSlice("template-url")
			api := getAPI(ctx)
			client := stacks.Client{
				CLI:              getCLI(ctx),
				StacksApi:        api.StacksApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Init(getContext(ctx), &params)
		},
	}
}

func newStacksRemoveCmd() cli.Command {
	var params stacks.RemoveParams
	return cli.Command{
		Name:    "rm",
		Aliases: []string{"remove", "uninstall"},
		Usage:   "Remove a stack(s) and all associated resources",
		Before:  middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			append(commonFlags(), getOrgFlags(&params.OrgParams)...),
			&cli.StringSliceFlag{
				Name:     "stack-id",
				Usage:    "Stack IDs to be removed",
				Required: true,
			},
			&cli.BoolFlag{
				Name:        "force",
				Usage:       "Remove stack without confirmation prompt",
				Destination: &params.Force,
			},
		),
		Action: func(ctx *cli.Context) error {
			params.Ids = ctx.StringSlice("stack-id")
			api := getAPI(ctx)
			client := stacks.Client{
				CLI:              getCLI(ctx),
				StacksApi:        api.StacksApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Remove(getContext(ctx), &params)
		},
	}
}

func newStacksUpdateCmd() cli.Command {
	var params stacks.UpdateParams
	return cli.Command{
		Name:  "update",
		Usage: "Update a stack",
		Description: `The stack update command updates a stack.

Examples:
	# Update a stack with a name and description
	influx stacks update -i $STACK_ID -n $STACK_NAME -d $STACK_DESCRIPTION

	# Update a stack with a name and urls to associate with stack.
	influx stacks update --stack-id $STACK_ID --stack-name $STACK_NAME --template-url $PATH_TO_TEMPLATE

	# Update stack with new resources to manage
	influx stacks update \
		--stack-id $STACK_ID \
		--addResource=Bucket=$BUCKET_ID \
		--addResource=Dashboard=$DASH_ID

	# Update stack with new resources to manage and export stack
	# as a template
	influx stacks update \
		--stack-id $STACK_ID \
		--addResource=Bucket=$BUCKET_ID \
		--export-file /path/to/file.yml

For information about how stacks work with InfluxDB templates, see
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/stacks/
and
https://docs.influxdata.com/influxdb/latest/reference/cli/influx/stacks/update/
`,
		Before: middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlags(),
			&cli.StringFlag{
				Name:        "stack-id, i",
				Usage:       "ID of stack",
				Destination: &params.Id,
				Required:    true,
			},
			cli.StringFlag{
				Name:  "stack-name, n",
				Usage: "New name for the stack",
			},
			cli.StringFlag{
				Name:  "stack-description, d",
				Usage: "New description for the stack",
			},
			cli.StringSliceFlag{
				Name:  "template-url, u",
				Usage: "New template URLs to associate with the stack",
			},
			cli.StringSliceFlag{
				Name:  "addResource",
				Usage: "Additional resources to associate with the stack",
			},
			cli.StringFlag{
				Name:      "export-file, f",
				Usage:     "Destination for exported template",
				TakesFile: true,
			},
		),
		Action: func(ctx *cli.Context) error {
			if ctx.IsSet("stack-name") {
				name := ctx.String("stack-name")
				params.Name = &name
			}
			if ctx.IsSet("stack-description") {
				desc := ctx.String("stack-description")
				params.Description = &desc
			}
			if ctx.IsSet("template-url") {
				urls := ctx.StringSlice("template-url")
				params.URLs = urls
			}

			rawResources := ctx.StringSlice("addResource")
			for _, res := range rawResources {
				pieces := strings.Split(res, "=")
				if len(pieces) != 2 {
					return fmt.Errorf("invalid resource specification %q, must have format `KIND=ID`", res)
				}
				params.AddedResources = append(params.AddedResources, stacks.AddedResource{
					Kind: pieces[0],
					Id:   pieces[1],
				})
			}

			cli := getCLI(ctx)
			outParams, closer, err := template.ParseOutParams(ctx.String("export-file"), cli.StdIO)
			if closer != nil {
				defer closer()
			}
			if err != nil {
				return err
			}
			params.OutParams = outParams

			api := getAPI(ctx)
			client := stacks.Client{
				CLI:          cli,
				StacksApi:    api.StacksApi,
				TemplatesApi: api.TemplatesApi,
			}
			return client.Update(getContext(ctx), &params)
		},
	}
}
