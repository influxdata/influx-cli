package internal_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/duration"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/require"
)

func Test_SetupConfigNameCollision(t *testing.T) {
	t.Parallel()

	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{Allowed: api.PtrBool(true)}, nil, nil
		},
	}

	cfg := "foo"
	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return map[string]config.Config{cfg: {}}, nil
		},
	}
	cli := &internal.CLI{ConfigService: configSvc}

	err := cli.Setup(context.Background(), client, &internal.SetupParams{ConfigName: cfg})
	require.Error(t, err)
	require.Contains(t, err.Error(), cfg)
	require.Contains(t, err.Error(), "already exists")
}

func Test_SetupConfigNameRequired(t *testing.T) {
	t.Parallel()

	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{Allowed: api.PtrBool(true)}, nil, nil
		},
	}

	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return map[string]config.Config{"foo": {}}, nil
		},
	}
	cli := &internal.CLI{ConfigService: configSvc}

	err := cli.Setup(context.Background(), client, &internal.SetupParams{})
	require.Error(t, err)
	require.Equal(t, internal.ErrConfigNameRequired, err)
}

func Test_SetupAlreadySetup(t *testing.T) {
	t.Parallel()

	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{Allowed: api.PtrBool(false)}, nil, nil
		},
	}

	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return map[string]config.Config{"foo": {}}, nil
		},
	}
	cli := &internal.CLI{ConfigService: configSvc}

	err := cli.Setup(context.Background(), client, &internal.SetupParams{})
	require.Error(t, err)
	require.Equal(t, internal.ErrAlreadySetUp, err)
}

func Test_SetupCheckFailed(t *testing.T) {
	t.Parallel()

	e := "oh no"
	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{}, nil, errors.New(e)
		},
	}

	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return nil, nil
		},
	}
	cli := &internal.CLI{ConfigService: configSvc}

	err := cli.Setup(context.Background(), client, &internal.SetupParams{})
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_SetupSuccessNoninteractive(t *testing.T) {
	t.Parallel()

	retentionSecs := int64(duration.Week.Seconds())
	params := internal.SetupParams{
		Username:   "user",
		Password:   "mysecretpassword",
		AuthToken:  "mytoken",
		Org:        "org",
		Bucket:     "bucket",
		Retention:  fmt.Sprintf("%ds", retentionSecs),
		Force:      true,
		ConfigName: "my-config",
	}
	resp := api.OnboardingResponse{
		Auth:   &api.Authorization{Token: &params.AuthToken},
		Org:    &api.Organization{Name: params.Org},
		User:   &api.UserResponse{Name: params.Username},
		Bucket: &api.Bucket{Name: params.Bucket},
	}
	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{Allowed: api.PtrBool(true)}, nil, nil
		},
		PostSetupExecuteFn: func(req api.ApiPostSetupRequest) (api.OnboardingResponse, *http.Response, error) {
			body := req.GetOnboardingRequest()
			require.Equal(t, params.Username, body.Username)
			require.Equal(t, params.Password, *body.Password)
			require.Equal(t, params.AuthToken, *body.Token)
			require.Equal(t, params.Org, body.Org)
			require.Equal(t, params.Bucket, body.Bucket)
			require.Equal(t, retentionSecs, *body.RetentionPeriodSeconds)
			return resp, nil, nil
		},
	}

	host := "fake-host"
	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return nil, nil
		},
		CreateConfigFn: func(cfg config.Config) (config.Config, error) {
			require.Equal(t, params.ConfigName, cfg.Name)
			require.Equal(t, params.AuthToken, cfg.Token)
			require.Equal(t, host, cfg.Host)
			require.Equal(t, params.Org, cfg.Org)
			return cfg, nil
		},
	}
	stdio := mock.NewMockStdio(nil, true)
	cli := &internal.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio}
	require.NoError(t, cli.Setup(context.Background(), client, &params))

	outLines := strings.Split(strings.TrimSpace(stdio.Stdout()), "\n")
	require.Len(t, outLines, 2)
	header, data := outLines[0], outLines[1]
	require.Regexp(t, "User\\s+Organization\\s+Bucket", header)
	require.Regexp(t, fmt.Sprintf("%s\\s+%s\\s+%s", params.Username, params.Org, params.Bucket), data)
}

func Test_SetupSuccessInteractive(t *testing.T) {
	t.Parallel()

	retentionSecs := int64(duration.Week.Seconds())
	retentionHrs := int(duration.Week.Hours())
	username := "user"
	password := "mysecretpassword"
	token := "mytoken"
	org := "org"
	bucket := "bucket"

	resp := api.OnboardingResponse{
		Auth:   &api.Authorization{Token: &token},
		Org:    &api.Organization{Name: org},
		User:   &api.UserResponse{Name: username},
		Bucket: &api.Bucket{Name: bucket},
	}
	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{Allowed: api.PtrBool(true)}, nil, nil
		},
		PostSetupExecuteFn: func(req api.ApiPostSetupRequest) (api.OnboardingResponse, *http.Response, error) {
			body := req.GetOnboardingRequest()
			require.Equal(t, username, body.Username)
			require.Equal(t, password, *body.Password)
			require.Nil(t, body.Token)
			require.Equal(t, org, body.Org)
			require.Equal(t, bucket, body.Bucket)
			require.Equal(t, retentionSecs, *body.RetentionPeriodSeconds)
			return resp, nil, nil
		},
	}

	host := "fake-host"
	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return nil, nil
		},
		CreateConfigFn: func(cfg config.Config) (config.Config, error) {
			require.Equal(t, config.DefaultConfig.Name, cfg.Name)
			require.Equal(t, token, cfg.Token)
			require.Equal(t, host, cfg.Host)
			require.Equal(t, org, cfg.Org)
			return cfg, nil
		},
	}
	stdio := mock.NewMockStdio(map[string]string{
		"Please type your primary username":                             username,
		"Please type your password":                                     password,
		"Please type your password again":                               password,
		"Please type your primary organization name":                    org,
		"Please type your primary bucket name":                          bucket,
		"Please type your retention period in hours, or 0 for infinite": strconv.Itoa(retentionHrs),
	}, true)
	cli := &internal.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio}
	require.NoError(t, cli.Setup(context.Background(), client, &internal.SetupParams{}))

	outLines := strings.Split(strings.TrimSpace(stdio.Stdout()), "\n")
	require.Len(t, outLines, 2)
	header, data := outLines[0], outLines[1]
	require.Regexp(t, "User\\s+Organization\\s+Bucket", header)
	require.Regexp(t, fmt.Sprintf("%s\\s+%s\\s+%s", username, org, bucket), data)
}

func Test_SetupPasswordParamToShort(t *testing.T) {
	t.Parallel()

	retentionSecs := int64(duration.Week.Seconds())
	params := internal.SetupParams{
		Username:  "user",
		Password:  "2short",
		AuthToken: "mytoken",
		Org:       "org",
		Bucket:    "bucket",
		Retention: fmt.Sprintf("%ds", retentionSecs),
		Force:     false,
	}
	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{Allowed: api.PtrBool(true)}, nil, nil
		},
	}

	host := "fake-host"
	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return nil, nil
		},
	}
	stdio := mock.NewMockStdio(nil, false)
	cli := &internal.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio}
	err := cli.Setup(context.Background(), client, &params)
	require.Equal(t, internal.ErrPasswordIsTooShort, err)
}

func Test_SetupCancelAtConfirmation(t *testing.T) {
	t.Parallel()

	retentionSecs := int64(duration.Week.Seconds())
	params := internal.SetupParams{
		Username:  "user",
		Password:  "mysecretpassword",
		AuthToken: "mytoken",
		Org:       "org",
		Bucket:    "bucket",
		Retention: fmt.Sprintf("%ds", retentionSecs),
		Force:     false,
	}
	client := &mock.SetupApi{
		GetSetupExecuteFn: func(api.ApiGetSetupRequest) (api.InlineResponse200, *http.Response, error) {
			return api.InlineResponse200{Allowed: api.PtrBool(true)}, nil, nil
		},
	}

	host := "fake-host"
	configSvc := &mock.ConfigService{
		ListConfigsFn: func() (config.Configs, error) {
			return nil, nil
		},
	}
	stdio := mock.NewMockStdio(nil, false)
	cli := &internal.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio}
	err := cli.Setup(context.Background(), client, &params)
	require.Equal(t, internal.ErrSetupCanceled, err)
}
