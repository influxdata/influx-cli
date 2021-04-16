package main

import (
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
	"github.com/urfave/cli/v2"
)

// Fields set via ldflags at build time.
var (
	version = "dev"
	commit  = "none"
	date    = ""
)

var (
	tokenFlag      = "token"
	hostFlag       = "host"
	skipVerifyFlag = "skip-verify"
	traceIdFlag    = "trace-debug-id"
	configPathFlag = "config-path"
	configNameFlag = "active-config"
	httpDebugFlag  = "http-debug"
)

// loadConfig reads CLI configs from disk, returning the config with the
// name specified over the CLI (or default if no name was given).
func loadConfig(ctx *cli.Context) (config.Config, error) {
	configs := config.GetConfigsOrDefault(ctx.String(configPathFlag))
	configName := ctx.String(configNameFlag)
	if configName != "" {
		if err := configs.Switch(configName); err != nil {
			return config.Config{}, err
		}
	}
	return configs.Active(), nil
}

// newApiClient returns an API client configured to communicate with a remote InfluxDB instance over HTTP.
// Client parameters are pulled from the CLI context.
func newApiClient(ctx *cli.Context, injectToken bool) (*api.APIClient, error) {
	cfg, err := loadConfig(ctx)
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
	apiConfig.Debug = ctx.Bool(httpDebugFlag)

	return api.NewAPIClient(apiConfig), nil
}

// newCli builds a CLI core that reads from stdin, writes to stdout/stderr, and
// optionally tracks a trace ID specified over the CLI.
func newCli(ctx *cli.Context) *internal.CLI {
	return &internal.CLI{
		Stdin:   ctx.App.Reader,
		Stdout:  ctx.App.Writer,
		Stderr:  ctx.App.ErrWriter,
		TraceId: ctx.String(traceIdFlag),
	}
}

func main() {
	if len(date) == 0 {
		date = time.Now().UTC().Format(time.RFC3339)
	}

	cli.VersionFlag = nil

	// NOTE: urfave/cli has dedicated support for global flags, but it only parses those flags
	// if they're specified before any command names. This is incompatible with the old influx
	// CLI, which repeatedly registered common flags on every "leaf" command, forcing the flags
	// to be specified after _all_ command names were given.
	//
	// We replicate the pattern from the old CLI so existing scripts and docs stay valid.
	commonFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    tokenFlag,
			Usage:   "Authentication token",
			Aliases: []string{"t"},
			EnvVars: []string{"INFLUX_TOKEN"},
		},
		&cli.StringFlag{
			Name:    hostFlag,
			Usage:   "HTTP address of InfluxDB",
			EnvVars: []string{"INFLUX_HOST"},
		},
		&cli.BoolFlag{
			Name:  skipVerifyFlag,
			Usage: "Skip TLS certificate chain and host name verification",
		},
		&cli.StringFlag{
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

	app := cli.App{
		Name:      "influx",
		Usage:     "Influx Client",
		UsageText: "influx [command]",
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Print the influx CLI version",
				Action: func(*cli.Context) error {
					fmt.Printf("Influx CLI %s (git: %s) build_date: %s\n", version, commit, date)
					return nil
				},
			},
			{
				Name:  "ping",
				Usage: "Check the InfluxDB /health endpoint",
				Flags: commonFlags,
				Action: func(ctx *cli.Context) error {
					client, err := newApiClient(ctx, false)
					if err != nil {
						return err
					}
					return newCli(ctx).Ping(ctx.Context, client.HealthApi)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
