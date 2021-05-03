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

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
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
	tokenFlag       = "token"
	hostFlag        = "host"
	skipVerifyFlag  = "skip-verify"
	traceIdFlag     = "trace-debug-id"
	configPathFlag  = "configs-path"
	configNameFlag  = "active-config"
	httpDebugFlag   = "http-debug"
	printJsonFlag   = "json"
	hideHeadersFlag = "hide-headers"
)

// NOTE: urfave/cli has dedicated support for global flags, but it only parses those flags
// if they're specified before any command names. This is incompatible with the old influx
// CLI, which repeatedly registered common flags on every "leaf" command, forcing the flags
// to be specified after _all_ command names were given.
//
// We replicate the pattern from the old CLI so existing scripts and docs stay valid.

// Flags used by all CLI commands.
var coreFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    hostFlag,
		Usage:   "HTTP address of InfluxDB",
		EnvVars: []string{"INFLUX_HOST"},
	},
	&cli.BoolFlag{
		Name:  skipVerifyFlag,
		Usage: "Skip TLS certificate chain and host name verification",
	},
	&cli.PathFlag{
		Name:    configPathFlag,
		Usage:   "Path to the influx CLI configurations",
		EnvVars: []string{"INFLUX_CLI_CONFIGS_PATH"},
	},
	&cli.StringFlag{
		Name:    configNameFlag,
		Usage:   "Config name to use for command",
		Aliases: []string{"c"},
		EnvVars: []string{"INFLUX_ACTIVE_CONFIG"},
	},
	&cli.StringFlag{
		Name:    traceIdFlag,
		Hidden:  true,
		EnvVars: []string{"INFLUX_TRACE_DEBUG_ID"},
	},
	&cli.BoolFlag{
		Name:   httpDebugFlag,
		Hidden: true,
	},
}

// Flags used by commands that display API resources to the user.
var printFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    printJsonFlag,
		Usage:   "Output data as JSON",
		EnvVars: []string{"INFLUX_OUTPUT_JSON"},
	},
	&cli.BoolFlag{
		Name:    hideHeadersFlag,
		Usage:   "Hide the table headers in output data",
		EnvVars: []string{"INFLUX_HIDE_HEADERS"},
	},
}

// Flag used by commands that hit an authenticated API.
var commonTokenFlag = cli.StringFlag{
	Name:    tokenFlag,
	Usage:   "Authentication token",
	Aliases: []string{"t"},
	EnvVars: []string{"INFLUX_TOKEN"},
}

var commonFlagsNoToken = append(coreFlags, printFlags...)
var commonFlagsNoPrint = append(coreFlags, &commonTokenFlag)
var commonFlags = append(commonFlagsNoToken, &commonTokenFlag)

// newCli builds a CLI core that reads from stdin, writes to stdout/stderr, manages a local config store,
// and optionally tracks a trace ID specified over the CLI.
func newCli(ctx *cli.Context) (*internal.CLI, error) {
	configPath := ctx.String(configPathFlag)
	var err error
	if configPath == "" {
		configPath, err = config.DefaultPath()
		if err != nil {
			return nil, err
		}
	}
	configSvc := config.NewLocalConfigService(configPath)
	var activeConfig config.Config
	if ctx.IsSet(configNameFlag) {
		if activeConfig, err = configSvc.SwitchActive(ctx.String(configNameFlag)); err != nil {
			return nil, err
		}
	} else if activeConfig, err = configSvc.Active(); err != nil {
		return nil, err
	}

	return &internal.CLI{
		StdIO:            stdio.TerminalStdio,
		PrintAsJSON:      ctx.Bool(printJsonFlag),
		HideTableHeaders: ctx.Bool(hideHeadersFlag),
		ActiveConfig:     activeConfig,
		ConfigService:    configSvc,
	}, nil
}

// newApiClient returns an API client configured to communicate with a remote InfluxDB instance over HTTP.
// Client parameters are pulled from the CLI context.
func newApiClient(ctx *cli.Context, cli *internal.CLI, injectToken bool) (*api.APIClient, error) {
	cfg, err := cli.ConfigService.Active()
	if err != nil {
		return nil, err
	}
	if ctx.IsSet(tokenFlag) {
		cfg.Token = ctx.String(tokenFlag)
	}
	if ctx.IsSet(hostFlag) {
		cfg.Host = ctx.String(hostFlag)
	}

	parsedHost, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("host URL %q is invalid: %w", cfg.Host, err)
	}

	clientTransport := http.DefaultTransport.(*http.Transport)
	clientTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: ctx.Bool(skipVerifyFlag)}

	apiConfig := api.NewConfiguration()
	apiConfig.Host = parsedHost.Host
	apiConfig.Scheme = parsedHost.Scheme
	apiConfig.UserAgent = fmt.Sprintf("influx/%s (%s) Sha/%s Date/%s", version, runtime.GOOS, commit, date)
	apiConfig.HTTPClient = &http.Client{Transport: clientTransport}
	if injectToken {
		apiConfig.DefaultHeader["Authorization"] = fmt.Sprintf("Token %s", cfg.Token)
	}
	if ctx.IsSet(traceIdFlag) {
		// NOTE: This is circumventing our codegen. If the header we use for tracing ever changes,
		// we'll need to manually update the string here to match.
		//
		// The alternative is to pass the trace ID to the business logic for every CLI command, and
		// use codegen'd logic to set the header on every HTTP request. Early versions of the CLI
		// used that technique, and we found it to be error-prone and easy to forget during testing.
		apiConfig.DefaultHeader["Zap-Trace-Span"] = ctx.String(traceIdFlag)
	}
	apiConfig.Debug = ctx.Bool(httpDebugFlag)

	return api.NewAPIClient(apiConfig), nil
}

var app = cli.App{
	Name:      "influx",
	Usage:     "Influx Client",
	UsageText: "influx [command]",
	Commands: []*cli.Command{
		newVersionCmd(),
		newPingCmd(),
		newSetupCmd(),
		newWriteCmd(),
		newBucketCmd(),
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

func getCLI(ctx *cli.Context) *internal.CLI {
	i, ok := ctx.App.Metadata["cli"].(*internal.CLI)
	if !ok {
		panic("missing CLI")
	}
	return i
}

func withApi() cli.BeforeFunc {
	makeFn := func(key string, injectToken bool) cli.BeforeFunc {
		return func(ctx *cli.Context) error {
			c := getCLI(ctx)
			apiClient, err := newApiClient(ctx, c, injectToken)
			if err != nil {
				return err
			}
			ctx.App.Metadata[key] = apiClient
			return nil
		}
	}
	return middleware.WithBeforeFns(makeFn("api", true), makeFn("api-no-token", false))
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
