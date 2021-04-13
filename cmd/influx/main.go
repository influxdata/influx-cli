package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/internal/ping"
	"github.com/influxdata/influx-cli/v2/kit/tracing"
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
)

// newApiClient returns an API client configured to communicate with a remote InfluxDB instance over HTTP.
// Client parameters are pulled from the CLI context.
func newApiClient(ctx *cli.Context, injectToken bool) (api.ClientWithResponsesInterface, error) {
	clientTransport := http.DefaultTransport.(*http.Transport)
	clientTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: ctx.Bool(skipVerifyFlag)}

	client := &http.Client{Transport: clientTransport}
	userAgent := fmt.Sprintf("influx/%s (%s) Sha/%s Date/%s", version, runtime.GOOS, commit, date)

	opts := []api.ClientOption{
		api.WithHTTPClient(client),
		api.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
			req.Header.Set("User-Agent", userAgent)
			return nil
		}),
	}
	if injectToken {
		authHeader := fmt.Sprintf("Token %s", ctx.String(tokenFlag))
		opts = append(opts, api.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", authHeader)
			return nil
		}))
	}

	return api.NewClientWithResponses(ctx.String(hostFlag), opts...)
}

// tracingCtx bundles the Jaeger trace ID given on the CLI (if any) with
// the underlying CLI context.
func tracingCtx(ctx *cli.Context) tracing.Context {
	return tracing.WrapContext(ctx.Context, ctx.String(traceIdFlag))
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
			Name:    "configs-path",
			Usage:   "Path to the influx CLI configurations",
			EnvVars: []string{"INFLUX_CLI_CONFIGS_PATH"},
		},
		&cli.StringFlag{
			Name:    "active-config",
			Usage:   "Config name to use for command",
			Aliases: []string{"c"},
			EnvVars: []string{"INFLUX_ACTIVE_CONFIG"},
		},
		&cli.StringFlag{
			Name:    traceIdFlag,
			Hidden:  true,
			EnvVars: []string{"INFLUX_TRACE_DEBUG_ID"},
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
					return ping.Ping(tracingCtx(ctx), client)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
