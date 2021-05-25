package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/cmd"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/stdio"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli/v2"
)

const (
	tokenFlagName       = "token"
	hostFlagName        = "host"
	skipVerifyFlagName  = "skip-verify"
	traceIdFlagName     = "trace-debug-id"
	configPathFlagName  = "configs-path"
	configNameFlagName  = "active-config"
	httpDebugFlagName   = "http-debug"
	printJsonFlagName   = "json"
	hideHeadersFlagName = "hide-headers"
)

// newCli builds a CLI core that reads from stdin, writes to stdout/stderr, manages a local config store,
// and optionally tracks a trace ID specified over the CLI.
func newCli(ctx *cli.Context) (cmd.CLI, error) {
	configPath := ctx.String(configPathFlagName)
	var err error
	if configPath == "" {
		configPath, err = config.DefaultPath()
		if err != nil {
			return cmd.CLI{}, err
		}
	}
	configSvc := config.NewLocalConfigService(configPath)
	var activeConfig config.Config
	if ctx.IsSet(configNameFlagName) {
		if activeConfig, err = configSvc.SwitchActive(ctx.String(configNameFlagName)); err != nil {
			return cmd.CLI{}, err
		}
	} else if activeConfig, err = configSvc.Active(); err != nil {
		return cmd.CLI{}, err
	}

	return cmd.CLI{
		StdIO:            stdio.TerminalStdio,
		PrintAsJSON:      ctx.Bool(printJsonFlagName),
		HideTableHeaders: ctx.Bool(hideHeadersFlagName),
		ActiveConfig:     activeConfig,
		ConfigService:    configSvc,
	}, nil
}

// newApiClient returns an API client configured to communicate with a remote InfluxDB instance over HTTP.
// Client parameters are pulled from the CLI context.
func newApiClient(ctx *cli.Context, configSvc config.Service, injectToken bool) (*api.APIClient, error) {
	cfg, err := configSvc.Active()
	if err != nil {
		return nil, err
	}
	if ctx.IsSet(tokenFlagName) {
		cfg.Token = ctx.String(tokenFlagName)
	}
	if ctx.IsSet(hostFlagName) {
		cfg.Host = ctx.String(hostFlagName)
	}

	parsedHost, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("host URL %q is invalid: %w", cfg.Host, err)
	}

	clientTransport := http.DefaultTransport.(*http.Transport)
	clientTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: ctx.Bool(skipVerifyFlagName)}

	apiConfig := api.NewConfiguration()
	apiConfig.Host = parsedHost.Host
	apiConfig.Scheme = parsedHost.Scheme
	apiConfig.UserAgent = fmt.Sprintf("influx/%s (%s) Sha/%s Date/%s", version, runtime.GOOS, commit, date)
	apiConfig.HTTPClient = &http.Client{Transport: clientTransport}
	if injectToken {
		apiConfig.DefaultHeader["Authorization"] = fmt.Sprintf("Token %s", cfg.Token)
	}
	if ctx.IsSet(traceIdFlagName) {
		// NOTE: This is circumventing our codegen. If the header we use for tracing ever changes,
		// we'll need to manually update the string here to match.
		//
		// The alternative is to pass the trace ID to the business logic for every CLI command, and
		// use codegen'd logic to set the header on every HTTP request. Early versions of the CLI
		// used that technique, and we found it to be error-prone and easy to forget during testing.
		apiConfig.DefaultHeader["Zap-Trace-Span"] = ctx.String(traceIdFlagName)
	}
	apiConfig.Debug = ctx.Bool(httpDebugFlagName)

	return api.NewAPIClient(apiConfig), nil
}

func withCli() cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		c, err := newCli(ctx)
		if err != nil {
			return err
		}
		ctx.App.Metadata["cli"] = c
		return nil
	}
}

func getCLI(ctx *cli.Context) cmd.CLI {
	i, ok := ctx.App.Metadata["cli"].(cmd.CLI)
	if !ok {
		panic("missing CLI")
	}
	return i
}

func withApi(injectToken bool) cli.BeforeFunc {
	key := "api-no-token"
	if injectToken {
		key = "api"
	}

	makeFn := func(ctx *cli.Context) error {
		c := getCLI(ctx)
		apiClient, err := newApiClient(ctx, c.ConfigService, injectToken)
		if err != nil {
			return err
		}
		ctx.App.Metadata[key] = apiClient
		return nil
	}
	return middleware.WithBeforeFns(makeFn)
}

func getAPI(ctx *cli.Context) *api.APIClient {
	i, ok := ctx.App.Metadata["api"].(*api.APIClient)
	if !ok {
		panic("missing APIClient with token")
	}
	return i
}

func getAPINoToken(ctx *cli.Context) *api.APIClient {
	i, ok := ctx.App.Metadata["api-no-token"].(*api.APIClient)
	if !ok {
		panic("missing APIClient without token")
	}
	return i
}

// NOTE: urfave/cli has dedicated support for global flags, but it only parses those flags
// if they're specified before any command names. This is incompatible with the old influx
// CLI, which repeatedly registered common flags on every "leaf" command, forcing the flags
// to be specified after _all_ command names were given.
//
// We replicate the pattern from the old CLI so existing scripts and docs stay valid.

// configPathFlag returns the flag used by commands that access the CLI's local config store.
func configPathFlag() cli.Flag {
	return &cli.PathFlag{
		Name:    configPathFlagName,
		Usage:   "Path to the influx CLI configurations",
		EnvVars: []string{"INFLUX_CLI_CONFIGS_PATH"},
	}
}

// coreFlags returns flags used by all CLI commands that make HTTP requests.
func coreFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    hostFlagName,
			Usage:   "HTTP address of InfluxDB",
			EnvVars: []string{"INFLUX_HOST"},
		},
		&cli.BoolFlag{
			Name:  skipVerifyFlagName,
			Usage: "Skip TLS certificate chain and host name verification",
		},
		configPathFlag(),
		&cli.StringFlag{
			Name:    configNameFlagName,
			Usage:   "Config name to use for command",
			Aliases: []string{"c"},
			EnvVars: []string{"INFLUX_ACTIVE_CONFIG"},
		},
		&cli.StringFlag{
			Name:    traceIdFlagName,
			Hidden:  true,
			EnvVars: []string{"INFLUX_TRACE_DEBUG_ID"},
		},
		&cli.BoolFlag{
			Name:   httpDebugFlagName,
			Hidden: true,
		},
	}
}

// printFlags returns flags used by commands that display API resources to the user.
func printFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    printJsonFlagName,
			Usage:   "Output data as JSON",
			EnvVars: []string{"INFLUX_OUTPUT_JSON"},
		},
		&cli.BoolFlag{
			Name:    hideHeadersFlagName,
			Usage:   "Hide the table headers in output data",
			EnvVars: []string{"INFLUX_HIDE_HEADERS"},
		},
	}
}

// commonTokenFlag returns the flag used by commands that hit an authenticated API.
func commonTokenFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    tokenFlagName,
		Usage:   "Authentication token",
		Aliases: []string{"t"},
		EnvVars: []string{"INFLUX_TOKEN"},
	}
}

// commonFlagsNoToken returns flags used by commands that display API resources to the user, but don't need auth.
func commonFlagsNoToken() []cli.Flag {
	return append(coreFlags(), printFlags()...)
}

// commonFlagsNoPrint returns flags used by commands that need auth, but don't display API resources to the user.
func commonFlagsNoPrint() []cli.Flag {
	return append(coreFlags(), commonTokenFlag())
}

// commonFlags returns flags used by commands that need auth and display API resources to the user.
func commonFlags() []cli.Flag {
	return append(commonFlagsNoToken(), commonTokenFlag())
}

// getOrgFlags returns flags used by commands that are scoped to a single org, binding
// the flags to the given params container.
func getOrgFlags(params *cmd.OrgParams) []cli.Flag {
	return []cli.Flag{
		&cli.GenericFlag{
			Name:    "org-id",
			Usage:   "The ID of the organization",
			EnvVars: []string{"INFLUX_ORG_ID"},
			Value:   &params.OrgID,
		},
		&cli.StringFlag{
			Name:        "org",
			Usage:       "The name of the organization",
			Aliases:     []string{"o"},
			EnvVars:     []string{"INFLUX_ORG"},
			Destination: &params.OrgName,
		},
	}
}

// getBucketFlags returns flags used by commands that are scoped to a single bucket, binding
// the flags to the given params container.
func getBucketFlags(params *cmd.BucketParams) []cli.Flag {
	return []cli.Flag{
		&cli.GenericFlag{
			Name:    "bucket-id",
			Usage:   "The bucket ID, required if name isn't provided",
			Aliases: []string{"i"},
			Value:   &params.BucketID,
		},
		&cli.StringFlag{
			Name:        "bucket",
			Usage:       "The bucket name, org or org-id will be required by choosing this",
			Aliases:     []string{"n"},
			Destination: &params.BucketName,
		},
	}
}

// getOrgBucketFlags returns flags used by commands that are scoped to a single org/bucket, binding
// the flags to the given params container.
func getOrgBucketFlags(c *cmd.OrgBucketParams) []cli.Flag {
	return append(getBucketFlags(&c.BucketParams), getOrgFlags(&c.OrgParams)...)
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
