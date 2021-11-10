package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/template"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	pkgtmpl "github.com/influxdata/influx-cli/v2/pkg/template"
	"github.com/urfave/cli"
)

type templateParams struct {
	orgParams clients.OrgParams
	files     cli.StringSlice
	urls      cli.StringSlice
	recurse   bool
	encoding  pkgtmpl.Encoding
}

func templateFlags(params *templateParams) []cli.Flag {
	return append(getOrgFlags(&params.orgParams), []cli.Flag{
		&cli.StringSliceFlag{
			Name:      "file, f",
			Usage:     "Path to template file; Supports file paths or (deprecated) HTTP(S) URLs",
			TakesFile: true,
			Value:     &params.files,
		},
		&cli.StringSliceFlag{
			Name:  "template-url, u",
			Usage: "HTTP(S) URL to template file",
			Value: &params.urls,
		},
		&cli.BoolFlag{
			Name:        "recurse, R",
			Usage:       "Process the directory used in -f, --file recursively. Useful when you want to manage related templates organized within the same directory.",
			Destination: &params.recurse,
		},
		&cli.GenericFlag{
			Name:  "encoding, e",
			Usage: "Encoding for the input stream. If a file is provided will gather encoding type from file extension. If extension provided will override.",
			Value: &params.encoding,
		},
	}...)
}

func (params templateParams) parseSources() ([]pkgtmpl.Source, error) {
	var deprecationShown bool
	var sources []pkgtmpl.Source
	for _, in := range params.files.Value() {
		// Heuristic to guess what's a URL and what's a local file.
		// TODO: Remove this once we stop supporting URLs in the --file arg.
		u, err := url.Parse(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse input path %q: %w", in, err)
		}
		if strings.HasPrefix(u.Scheme, "http") {
			if !deprecationShown {
				log.Println("WARN: Passing URLs via -f/--file is deprecated, please use -u/--template-url instead")
				deprecationShown = true
			}
			sources = append(sources, pkgtmpl.SourceFromURL(u, params.encoding))
		} else {
			fileSources, err := pkgtmpl.SourcesFromPath(in, params.recurse, params.encoding)
			if err != nil {
				return nil, err
			}
			sources = append(sources, fileSources...)
		}
	}
	for _, in := range params.urls.Value() {
		u, err := url.Parse(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse input URL: %q: %w", in, err)
		}
		sources = append(sources, pkgtmpl.SourceFromURL(u, params.encoding))
	}
	return sources, nil
}

func newTemplateCmd() cli.Command {
	var params struct {
		templateParams
		noColor        bool
		noTableBorders bool
	}
	return cli.Command{
		Name:  "template",
		Usage: "Summarize the provided template",
		Subcommands: []cli.Command{
			newTemplateValidateCmd(),
		},
		Flags: append(
			append(commonFlags(), templateFlags(&params.templateParams)...),
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
		),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			parsedParams := template.SummarizeParams{
				OrgParams:          params.orgParams,
				RenderTableColors:  !params.noColor,
				RenderTableBorders: !params.noTableBorders,
			}
			sources, err := params.parseSources()
			if err != nil {
				return err
			}
			parsedParams.Sources = sources

			api := getAPI(ctx)
			client := template.Client{
				CLI:              getCLI(ctx),
				TemplatesApi:     api.TemplatesApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Summarize(getContext(ctx), &parsedParams)
		},
	}
}

func newTemplateValidateCmd() cli.Command {
	var params templateParams
	return cli.Command{
		Name:   "validate",
		Usage:  "Validate the provided template",
		Flags:  append(commonFlagsNoPrint(), templateFlags(&params)...),
		Before: middleware.WithBeforeFns(withCli(), withApi(true), middleware.NoArgs),
		Action: func(ctx *cli.Context) error {
			parsedParams := template.ValidateParams{
				OrgParams: params.orgParams,
			}
			sources, err := params.parseSources()
			if err != nil {
				return err
			}
			parsedParams.Sources = sources

			api := getAPI(ctx)
			client := template.Client{
				CLI:              getCLI(ctx),
				TemplatesApi:     api.TemplatesApi,
				OrganizationsApi: api.OrganizationsApi,
			}
			return client.Validate(getContext(ctx), &parsedParams)
		},
	}
}
