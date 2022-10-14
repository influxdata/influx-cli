package config_test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	cmd "github.com/influxdata/influx-cli/v2/clients/config"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/require"
)

func TestClient_SwitchActive(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	stdio := mock.NewMockStdIO(ctrl)
	writtenBytes := bytes.Buffer{}
	stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

	name := "foo"
	cfg := config.Config{
		Name:   name,
		Active: true,
		Host:   "http://localhost:8086",
		Token:  "supersecret",
		Org:    "me",
	}
	svc := mock.NewMockConfigService(ctrl)
	svc.EXPECT().SwitchActive(gomock.Eq(name)).Return(cfg, nil)

	cli := cmd.Client{CLI: clients.CLI{ConfigService: svc, StdIO: stdio}}
	require.NoError(t, cli.SwitchActive(name))
	testutils.MatchLines(t, []string{
		`Active\s+Name\s+URL\s+Org`,
		fmt.Sprintf(`\*\s+%s\s+%s\s+%s`, cfg.Name, cfg.Host, cfg.Org),
	}, strings.Split(writtenBytes.String(), "\n"))
}

func TestClient_PrintActive(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	stdio := mock.NewMockStdIO(ctrl)
	writtenBytes := bytes.Buffer{}
	stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

	cfg := config.Config{
		Name:   "foo",
		Active: true,
		Host:   "http://localhost:8086",
		Token:  "supersecret",
		Org:    "me",
	}
	svc := mock.NewMockConfigService(ctrl)
	svc.EXPECT().Active().Return(cfg, nil)

	cli := cmd.Client{CLI: clients.CLI{ConfigService: svc, StdIO: stdio}}
	require.NoError(t, cli.PrintActive())
	testutils.MatchLines(t, []string{
		`Active\s+Name\s+URL\s+Org`,
		fmt.Sprintf(`\*\s+%s\s+%s\s+%s`, cfg.Name, cfg.Host, cfg.Org),
	}, strings.Split(writtenBytes.String(), "\n"))
}

func TestClient_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		cfg  config.Config
		err  error
	}{
		{
			name: "token",
			cfg: config.Config{
				Name:   "foo",
				Active: true,
				Host:   "http://localhost:8086",
				Token:  "supersecret",
				Org:    "me",
			},
			err: nil,
		},
		{
			name: "userpass",
			cfg: config.Config{
				Name:   "foo",
				Active: true,
				Host:   "http://localhost:8086",
				Cookie: base64.StdEncoding.EncodeToString([]byte("user:pass")),
				Org:    "me",
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

			cfg := tc.cfg
			svc := mock.NewMockConfigService(ctrl)
			svc.EXPECT().CreateConfig(cfg).Return(cfg, nil)

			cli := cmd.Client{CLI: clients.CLI{ConfigService: svc, StdIO: stdio}}
			err := cli.Create(cfg)
			require.NoError(t, err)
			testutils.MatchLines(t, []string{
				`Active\s+Name\s+URL\s+Org`,
				fmt.Sprintf(`\*\s+%s\s+%s\s+%s`, cfg.Name, cfg.Host, cfg.Org),
			}, strings.Split(writtenBytes.String(), "\n"))
		})
	}

}

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		in                   []string
		registerExpectations func(service *mock.MockConfigService)
		out                  []string
	}{
		{
			name: "empty",
		},
		{
			name: "one",
			in:   []string{"foo"},
			registerExpectations: func(svc *mock.MockConfigService) {
				svc.EXPECT().DeleteConfig(gomock.Eq("foo")).
					Return(config.Config{Name: "foo", Host: "bar", Org: "baz"}, nil)
			},
			out: []string{`\s+foo\s+bar\s+baz\s+true`},
		},
		{
			name: "many",
			in:   []string{"foo", "qux", "wibble"},
			registerExpectations: func(svc *mock.MockConfigService) {
				svc.EXPECT().DeleteConfig(gomock.Eq("foo")).
					Return(config.Config{Name: "foo", Host: "bar", Org: "baz"}, nil)
				svc.EXPECT().DeleteConfig(gomock.Eq("qux")).
					Return(config.Config{}, &api.Error{Code: api.ERRORCODE_NOT_FOUND})
				svc.EXPECT().DeleteConfig(gomock.Eq("wibble")).
					Return(config.Config{Name: "wibble", Host: "bar", Active: true}, nil)
			},
			out: []string{
				`\s+foo\s+bar\s+baz\s+true`,
				`\*\s+wibble\s+bar\s+true`,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

			svc := mock.NewMockConfigService(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(svc)
			}

			cli := cmd.Client{CLI: clients.CLI{ConfigService: svc, StdIO: stdio}}
			require.NoError(t, cli.Delete(tc.in))

			// Can't use our usual 'MatchLines' because list output depends on map iteration,
			// so the order isn't well-defined.
			out := writtenBytes.String()
			for _, l := range append([]string{`Active\s+Name\s+URL\s+Org\s+Deleted`}, tc.out...) {
				require.Regexp(t, l, out)
			}
		})
	}
}

func TestClient_Update(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		updates config.Config
		cfg     config.Config
		err     error
	}{
		{
			name: "token",
			updates: config.Config{
				Name:   "foo",
				Active: true,
				Token:  "doublesecret",
			},
			cfg: config.Config{
				Name:   "foo",
				Active: true,
				Host:   "http://localhost:8086",
				Token:  "doublesecret",
				Org:    "me",
			},
			err: nil,
		},
		{
			name: "userpass",
			updates: config.Config{
				Name:   "foo",
				Active: true,
				Cookie: base64.StdEncoding.EncodeToString([]byte("user:pass")),
			},
			cfg: config.Config{
				Name:   "foo",
				Active: true,
				Host:   "http://localhost:8086",
				Cookie: base64.StdEncoding.EncodeToString([]byte("user:pass")),
				Org:    "me",
			},
			err: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

			svc := mock.NewMockConfigService(ctrl)
			svc.EXPECT().UpdateConfig(tc.updates).Return(tc.cfg, nil)

			cli := cmd.Client{CLI: clients.CLI{ConfigService: svc, StdIO: stdio}}
			require.NoError(t, cli.Update(tc.updates))
			testutils.MatchLines(t, []string{
				`Active\s+Name\s+URL\s+Org`,
				fmt.Sprintf(`\*\s+%s\s+%s\s+%s`, tc.cfg.Name, tc.cfg.Host, tc.cfg.Org),
			}, strings.Split(writtenBytes.String(), "\n"))
		})
	}

}

func TestClient_List(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		cfgs     config.Configs
		expected []string
	}{
		{
			name: "empty",
		},
		{
			name: "one",
			cfgs: config.Configs{
				"foo": config.Config{Name: "foo", Host: "bar", Org: "baz"},
			},
			expected: []string{`\s+foo\s+bar\s+baz`},
		},
		{
			name: "many",
			cfgs: config.Configs{
				"foo":    config.Config{Name: "foo", Host: "bar", Org: "baz"},
				"wibble": config.Config{Name: "wibble", Host: "bar", Active: true},
			},
			expected: []string{
				`\s+foo\s+bar\s+baz`,
				`\*\s+wibble\s+bar`,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

			svc := mock.NewMockConfigService(ctrl)
			svc.EXPECT().ListConfigs().Return(tc.cfgs, nil)

			cli := cmd.Client{CLI: clients.CLI{ConfigService: svc, StdIO: stdio}}
			require.NoError(t, cli.List())

			// Can't use our usual 'MatchLines' because list output depends on map iteration,
			// so the order isn't well-defined.
			out := writtenBytes.String()
			for _, l := range append([]string{`Active\s+Name\s+URL\s+Org`}, tc.expected...) {
				require.Regexp(t, l, out)
			}
		})
	}
}
