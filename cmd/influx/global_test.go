package main

import (
	"flag"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

func TestNewAPIClient(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config.Config
		flags       [][2]string
		injectToken bool
		assertions  func(*testing.T, *api.APIClient)
	}{
		{
			name: "no token specified, without injectToken",
			cfg: config.Config{
				Name:   "test",
				Active: true,
				Host:   "http://localhost:8086",
				Org:    "test",
			},
			assertions: func(t *testing.T, c *api.APIClient) {
				assert.Equal(t, "localhost:8086", c.GetConfig().Host)
				assert.Equal(t, "http", c.GetConfig().Scheme)
				assert.True(t, strings.HasPrefix(c.GetConfig().UserAgent, "influx/"))
				assert.NotContains(t, c.GetConfig().DefaultHeader, "Authorization")
				assert.Equal(t, false, c.GetConfig().Debug)
				transport := c.GetConfig().HTTPClient.Transport.(*http.Transport)
				assert.Equal(t, false, transport.TLSClientConfig.InsecureSkipVerify)
			},
		},
		{
			name: "no flags specified",
			cfg: config.Config{
				Name:   "test",
				Active: true,
				Host:   "http://localhost:8086",
				Token:  "test",
				Org:    "test",
			},
			injectToken: true,
			assertions: func(t *testing.T, c *api.APIClient) {
				assert.Equal(t, "localhost:8086", c.GetConfig().Host)
				assert.Equal(t, "http", c.GetConfig().Scheme)
				assert.True(t, strings.HasPrefix(c.GetConfig().UserAgent, "influx/"))
				assert.Equal(t, "Token test", c.GetConfig().DefaultHeader["Authorization"])
				assert.Equal(t, false, c.GetConfig().Debug)
				transport := c.GetConfig().HTTPClient.Transport.(*http.Transport)
				assert.Equal(t, false, transport.TLSClientConfig.InsecureSkipVerify)
			},
		},
		{
			name: "flags specified",
			cfg: config.Config{
				Name:   "test",
				Active: true,
				Host:   "http://localhost:8086",
				Token:  "test",
				Org:    "test",
			},
			flags: [][2]string{
				{"token", "token-from-flag"},
				{"host", "http://localhost:9999"},
				{"skip-verify", "true"},
				{"http-debug", "true"},
				{"trace-debug-id", "trace-id-from-flag"},
			},
			injectToken: true,
			assertions: func(t *testing.T, c *api.APIClient) {
				assert.Equal(t, "localhost:9999", c.GetConfig().Host)
				assert.Equal(t, "http", c.GetConfig().Scheme)
				assert.True(t, strings.HasPrefix(c.GetConfig().UserAgent, "influx/"))
				assert.Equal(t, "Token token-from-flag", c.GetConfig().DefaultHeader["Authorization"])
				assert.Equal(t, "trace-id-from-flag", c.GetConfig().DefaultHeader["Zap-Trace-Span"])
				assert.Equal(t, true, c.GetConfig().Debug)
				transport := c.GetConfig().HTTPClient.Transport.(*http.Transport)
				assert.Equal(t, true, transport.TLSClientConfig.InsecureSkipVerify)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			cmd := cli.Command{Name: "TEST", Flags: commonFlagsNoPrint()}
			flagSet := flag.NewFlagSet("", flag.ContinueOnError)
			for _, f := range cmd.Flags {
				f.Apply(flagSet)
			}
			for _, f := range tc.flags {
				require.NoError(t, flagSet.Set(f[0], f[1]))
			}
			ctx := cli.NewContext(nil, flagSet, nil)
			ctx.Command = cmd

			svc := mock.NewMockConfigService(ctrl)
			svc.EXPECT().Active().Return(tc.cfg, nil)

			c, err := newApiClient(ctx, svc, tc.injectToken)
			require.NoError(t, err)

			tc.assertions(t, c)
		})
	}
}

func TestNewAPIClientErrors(t *testing.T) {
	testCases := []struct {
		name                  string
		cfg                   config.Config
		injectToken           bool
		expectedErrorContains string
	}{
		{
			name: "invalid host",
			cfg: config.Config{
				Name:   "test",
				Active: true,
				Host:   ":/:invalid:host:value",
				Token:  "test",
				Org:    "test",
			},
			expectedErrorContains: "is invalid",
		},
		{
			name: "missing token",
			cfg: config.Config{
				Name:   "test",
				Active: true,
				Host:   "http://localhost:8086",
				Org:    "test",
			},
			injectToken:           true,
			expectedErrorContains: "influx token required",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			flagSet := flag.NewFlagSet("", flag.ContinueOnError)
			ctx := cli.NewContext(nil, flagSet, nil)

			svc := mock.NewMockConfigService(ctrl)
			svc.EXPECT().Active().Return(tc.cfg, nil)

			_, err := newApiClient(ctx, svc, tc.injectToken)
			require.Error(t, err)

			assert.Contains(t, err.Error(), tc.expectedErrorContains)
		})
	}
}
