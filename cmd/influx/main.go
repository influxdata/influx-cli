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
	"github.com/influxdata/influx-cli/v2/internal/stdio"
	"github.com/urfave/cli/v2"
)

// Fields set via ldflags at build time.
var (
	version = "dev"
	commit  = "none"
	date    = ""
)

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
		TraceId:          ctx.String(traceIdFlag),
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
	apiConfig.Debug = ctx.Bool(httpDebugFlag)

	return api.NewAPIClient(apiConfig), nil
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

	// Some commands (i.e. `setup` use custom help-text for the token flag).
	commonFlagsNoToken := []cli.Flag{
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

	// Most commands use this form of the token flag.
	//commonFlags := append(commonFlagsNoToken, &cli.StringFlag{
	//	Name:    tokenFlag,
	//	Usage:   "Authentication token",
	//	Aliases: []string{"t"},
	//	EnvVars: []string{"INFLUX_TOKEN"},
	//})

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
				Flags: commonFlagsNoToken,
				Action: func(ctx *cli.Context) error {
					cli, err := newCli(ctx)
					if err != nil {
						return err
					}
					client, err := newApiClient(ctx, cli, false)
					if err != nil {
						return err
					}
					return cli.Ping(ctx.Context, client.HealthApi)
				},
			},
			{
				Name:  "setup",
				Usage: "Setup instance with initial user, org, bucket",
				Flags: append(
					commonFlagsNoToken,
					&cli.StringFlag{
						Name:    "username",
						Usage:   "Name of initial user to create",
						Aliases: []string{"u"},
					},
					&cli.StringFlag{
						Name:    "password",
						Usage:   "Password to set on initial user",
						Aliases: []string{"p"},
					},
					&cli.StringFlag{
						Name:        tokenFlag,
						Usage:       "Auth token to set on the initial user",
						Aliases:     []string{"t"},
						EnvVars:     []string{"INFLUX_TOKEN"},
						DefaultText: "auto-generated",
					},
					&cli.StringFlag{
						Name:    "org",
						Usage:   "Name of initial organization to create",
						Aliases: []string{"o"},
					},
					&cli.StringFlag{
						Name:    "bucket",
						Usage:   "Name of initial bucket to create",
						Aliases: []string{"b"},
					},
					&cli.StringFlag{
						Name:        "retention",
						Usage:       "Duration initial bucket will retain data, or 0 for infinite",
						Aliases:     []string{"r"},
						DefaultText: "infinite",
					},
					&cli.BoolFlag{
						Name:    "force",
						Usage:   "Skip confirmation prompt",
						Aliases: []string{"f"},
					},
					&cli.StringFlag{
						Name:    "name",
						Usage:   "Name to set on CLI config generated for the InfluxDB instance, required if other configs exist",
						Aliases: []string{"n"},
					},
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
				),
				Action: func(ctx *cli.Context) error {
					cli, err := newCli(ctx)
					if err != nil {
						return err
					}
					client, err := newApiClient(ctx, cli, false)
					if err != nil {
						return err
					}
					return cli.Setup(ctx.Context, client.SetupApi, &internal.SetupParams{
						Username:   ctx.String("username"),
						Password:   ctx.String("password"),
						AuthToken:  ctx.String(tokenFlag),
						Org:        ctx.String("org"),
						Bucket:     ctx.String("bucket"),
						Retention:  ctx.String("retention"),
						Force:      ctx.Bool("force"),
						ConfigName: ctx.String("name"),
					})
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
