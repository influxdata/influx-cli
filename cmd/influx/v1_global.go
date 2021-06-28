package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"runtime"

	v1 "github.com/influxdata/influx-cli/v2/api/v1"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

// newApiV1Client returns an API clients configured to communicate with a remote InfluxDB instance over HTTP.
// Client parameters are pulled from the CLI context.
func newApiV1Client(ctx *cli.Context, configSvc config.Service, injectToken bool) (*v1.APIClient, error) {
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

	apiConfig := v1.NewConfiguration()
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

	return v1.NewAPIClient(apiConfig), nil
}

func withApiV1(injectToken bool) cli.BeforeFunc {
	key := "api-v1-no-token"
	if injectToken {
		key = "api-v1"
	}

	makeFn := func(ctx *cli.Context) error {
		c := getCLI(ctx)
		apiClient, err := newApiV1Client(ctx, c.ConfigService, injectToken)
		if err != nil {
			return err
		}
		ctx.App.Metadata[key] = apiClient
		return nil
	}
	return middleware.WithBeforeFns(makeFn)
}

func getAPIV1(ctx *cli.Context) *v1.APIClient {
	i, ok := ctx.App.Metadata["api-v1"].(*v1.APIClient)
	if !ok {
		panic("missing API V1 Client with token")
	}
	return i
}

func getAPIV1NoToken(ctx *cli.Context) *v1.APIClient {
	i, ok := ctx.App.Metadata["api-no-token"].(*v1.APIClient)
	if !ok {
		panic("missing APIClient without token")
	}
	return i
}
