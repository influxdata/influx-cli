package main

import (
	"fmt"
	"net/url"
	"runtime"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/stdio"
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
func newCli(ctx *cli.Context) (clients.CLI, error) {
	configPath := ctx.String(configPathFlagName)
	var err error
	if configPath == "" {
		configPath, err = config.DefaultPath()
		if err != nil {
			return clients.CLI{}, err
		}
	}
	configSvc := config.NewLocalConfigService(configPath)
	var activeConfig config.Config
	if ctx.IsSet(configNameFlagName) {
		if activeConfig, err = configSvc.SwitchActive(ctx.String(configNameFlagName)); err != nil {
			return clients.CLI{}, err
		}
	} else if activeConfig, err = configSvc.Active(); err != nil {
		return clients.CLI{}, err
	}

	return clients.CLI{
		StdIO:            stdio.TerminalStdio,
		PrintAsJSON:      ctx.Bool(printJsonFlagName),
		HideTableHeaders: ctx.Bool(hideHeadersFlagName),
		ActiveConfig:     activeConfig,
		ConfigService:    configSvc,
	}, nil
}

// newApiClient returns an API clients configured to communicate with a remote InfluxDB instance over HTTP.
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

	clientParams := api.ClientParams{
		UserAgent:        fmt.Sprintf("influx/%s (%s) Sha/%s Date/%s", version, runtime.GOOS, commit, date),
		AllowInsecureTLS: ctx.Bool(skipVerifyFlagName),
		Debug:            ctx.Bool(httpDebugFlagName),
	}

	parsedHost, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("host URL %q is invalid: %w", cfg.Host, err)
	}
	clientParams.Host = parsedHost

	if injectToken {
		clientParams.Token = &cfg.Token
	}
	if ctx.IsSet(traceIdFlagName) {
		clientParams.TraceId = api.PtrString(ctx.String(traceIdFlagName))
	}

	return api.NewApiClient(clientParams), nil
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

func getCLI(ctx *cli.Context) clients.CLI {
	i, ok := ctx.App.Metadata["cli"].(clients.CLI)
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
func getOrgFlags(params *clients.OrgParams) []cli.Flag {
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
func getBucketFlags(params *clients.BucketParams) []cli.Flag {
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
func getOrgBucketFlags(c *clients.OrgBucketParams) []cli.Flag {
	return append(getBucketFlags(&c.BucketParams), getOrgFlags(&c.OrgParams)...)
}
