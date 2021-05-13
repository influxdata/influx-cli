package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"

	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/cmd"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/stdio"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/influxdata/influx-cli/v2/pkg/signals"
	"github.com/urfave/cli/v2"
)

// Fields set via ldflags at build time.
var (
	version = "dev"
	commit  = "none"
	date    = ""
)

func init() {
	if len(date) == 0 {
		date = time.Now().UTC().Format(time.RFC3339)
	}

	cli.VersionFlag = nil
}

var (
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

// NOTE: urfave/cli has dedicated support for global flags, but it only parses those flags
// if they're specified before any command names. This is incompatible with the old influx
// CLI, which repeatedly registered common flags on every "leaf" command, forcing the flags
// to be specified after _all_ command names were given.
//
// We replicate the pattern from the old CLI so existing scripts and docs stay valid.

var configPathFlag = cli.PathFlag{
	Name:    configPathFlagName,
	Usage:   "Path to the influx CLI configurations",
	EnvVars: []string{"INFLUX_CLI_CONFIGS_PATH"},
}

// Flags used by all CLI commands that make HTTP requests.
var coreFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    hostFlagName,
		Usage:   "HTTP address of InfluxDB",
		EnvVars: []string{"INFLUX_HOST"},
	},
	&cli.BoolFlag{
		Name:  skipVerifyFlagName,
		Usage: "Skip TLS certificate chain and host name verification",
	},
	&configPathFlag,
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

// Flags used by commands that display API resources to the user.
var printFlags = []cli.Flag{
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

// Flag used by commands that hit an authenticated API.
var commonTokenFlag = cli.StringFlag{
	Name:    tokenFlagName,
	Usage:   "Authentication token",
	Aliases: []string{"t"},
	EnvVars: []string{"INFLUX_TOKEN"},
}

var commonFlagsNoToken = append(coreFlags, printFlags...)
var commonFlagsNoPrint = append(coreFlags, &commonTokenFlag)
var commonFlags = append(commonFlagsNoToken, &commonTokenFlag)

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

var app = cli.App{
	Name:                 "influx",
	Usage:                "Influx Client",
	UsageText:            "influx [command]",
	EnableBashCompletion: true,
	Commands: []*cli.Command{
		newVersionCmd(),
		newPingCmd(),
		newSetupCmd(),
		newWriteCmd(),
		newBucketCmd(),
		newCompletionCmd(),
		newBucketSchemaCmd(),
		newQueryCmd(),
		newConfigCmd(),
	},
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

func main() {
	ctx := signals.WithStandardSignals(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
