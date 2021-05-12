package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/influxdata/influx-cli/v2/internal/cmd"
	"github.com/influxdata/influx-cli/v2/internal/cmd/query"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

func newQueryCmd() *cli.Command {
	var orgParams cmd.OrgParams
	return &cli.Command{
		Name:        "query",
		Usage:       "Execute a Flux query",
		Description: "Execute a Flux query provided via the first argument, a file, or stdin",
		ArgsUsage:   "[query literal or '-' for stdin]",
		Before:      middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlagsNoPrint,
			&cli.GenericFlag{
				Name:    "org-id",
				Usage:   "The ID of the organization",
				EnvVars: []string{"INFLUX_ORG_ID"},
				Value:   &orgParams.OrgID,
			},
			&cli.StringFlag{
				Name:        "org",
				Usage:       "The name of the organization",
				Aliases:     []string{"o"},
				EnvVars:     []string{"INFLUX_ORG"},
				Destination: &orgParams.OrgName,
			},
			&cli.StringFlag{
				Name:    "file",
				Usage:   "Path to Flux query file",
				Aliases: []string{"f"},
			},
			&cli.BoolFlag{
				Name:    "raw",
				Usage:   "Display raw query results",
				Aliases: []string{"r"},
			},
			&cli.StringSliceFlag{
				Name:    "profilers",
				Usage:   "Names of Flux profilers to enable",
				Aliases: []string{"p"},
			},
		),
		Action: func(ctx *cli.Context) error {
			queryString, err := readQuery(ctx)
			if err != nil {
				return err
			}
			queryString = strings.TrimSpace(queryString)
			if queryString == "" {
				return errors.New("no query provided")
			}

			// The old CLI allowed specifying this either via repeated flags or
			// via a single flag w/ a comma-separated value.
			rawProfilers := ctx.StringSlice("profilers")
			var profilers []string
			for _, p := range rawProfilers {
				profilers = append(profilers, strings.Split(p, ",")...)
			}

			params := query.Params{
				OrgParams: orgParams,
				Query:     queryString,
				Profilers: profilers,
			}

			var printer query.ResultPrinter
			if ctx.Bool("raw") {
				printer = query.RawResultPrinter
			} else {
				printer = query.NewFormattingPrinter()
			}

			client := query.Client{
				CLI:           getCLI(ctx),
				QueryApi:      getAPI(ctx).QueryApi,
				ResultPrinter: printer,
			}
			return client.Query(ctx.Context, &params)
		},
	}
}

// readQuery reads a Flux query into memory from a file, args, or stdin based on CLI parameters.
func readQuery(ctx *cli.Context) (string, error) {
	nargs := ctx.NArg()
	file := ctx.String("file")

	if nargs > 1 {
		return "", fmt.Errorf("at most 1 query string can be specified over the CLI, got %d", ctx.NArg())
	}
	if nargs == 1 && file != "" {
		return "", errors.New("query can be specified via --file or over the CLI, not both")
	}

	readFile := func(path string) (string, error) {
		queryBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("failed to read query from %q: %w", path, err)
		}
		return string(queryBytes), nil
	}

	readStdin := func() (string, error) {
		queryBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read query from stdin: %w", err)
		}
		return string(queryBytes), err
	}

	if file != "" {
		return readFile(file)
	}
	if nargs == 0 {
		return readStdin()
	}

	arg := ctx.Args().Get(0)
	// Backwards compatibility.
	if strings.HasPrefix(arg, "@") {
		return readFile(arg[1:])
	} else if arg == "-" {
		return readStdin()
	} else {
		return arg, nil
	}
}
