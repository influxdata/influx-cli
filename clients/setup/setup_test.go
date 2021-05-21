package setup_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/influxdata/influx-cli/v2/clients"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/clients/setup"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/duration"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_SetupConfigNameCollision(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{Allowed: api.PtrBool(true)}, nil)

	cfg := "foo"
	configSvc := mock.NewMockConfigService(ctrl)
	configSvc.EXPECT().ListConfigs().Return(map[string]config.Config{cfg: {}}, nil)
	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc},
		SetupApi: client,
	}

	err := cli.Setup(context.Background(), &setup.Params{ConfigName: cfg})
	require.Error(t, err)
	require.Contains(t, err.Error(), cfg)
	require.Contains(t, err.Error(), "already exists")
}

func Test_SetupConfigNameRequired(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{Allowed: api.PtrBool(true)}, nil)

	configSvc := mock.NewMockConfigService(ctrl)
	configSvc.EXPECT().ListConfigs().Return(map[string]config.Config{"foo": {}}, nil)
	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc},
		SetupApi: client,
	}

	err := cli.Setup(context.Background(), &setup.Params{})
	require.Error(t, err)
	require.Equal(t, setup.ErrConfigNameRequired, err)
}

func Test_SetupAlreadySetup(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{Allowed: api.PtrBool(false)}, nil)

	configSvc := mock.NewMockConfigService(ctrl)
	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc},
		SetupApi: client,
	}

	err := cli.Setup(context.Background(), &setup.Params{})
	require.Error(t, err)
	require.Equal(t, setup.ErrAlreadySetUp, err)
}

func Test_SetupCheckFailed(t *testing.T) {
	t.Parallel()

	e := "oh no"
	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{}, errors.New(e))

	configSvc := mock.NewMockConfigService(ctrl)
	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc},
		SetupApi: client,
	}

	err := cli.Setup(context.Background(), &setup.Params{})
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_SetupSuccessNoninteractive(t *testing.T) {
	t.Parallel()

	retentionSecs := int64(duration.Week.Seconds())
	params := setup.Params{
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

	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{Allowed: api.PtrBool(true)}, nil)
	client.EXPECT().PostSetup(gomock.Any()).Return(api.ApiPostSetupRequest{ApiService: client})
	client.EXPECT().PostSetupExecute(tmock.MatchedBy(func(in api.ApiPostSetupRequest) bool {
		body := in.GetOnboardingRequest()
		return assert.NotNil(t, body) &&
			assert.Equal(t, params.Username, body.Username) &&
			assert.Equal(t, params.Password, *body.Password) &&
			assert.Equal(t, params.AuthToken, *body.Token) &&
			assert.Equal(t, params.Org, body.Org) &&
			assert.Equal(t, params.Bucket, body.Bucket) &&
			assert.Equal(t, retentionSecs, *body.RetentionPeriodSeconds)
	})).Return(resp, nil)

	host := "fake-host"
	configSvc := mock.NewMockConfigService(ctrl)
	configSvc.EXPECT().ListConfigs().Return(nil, nil)
	configSvc.EXPECT().CreateConfig(tmock.MatchedBy(func(in config.Config) bool {
		return assert.Equal(t, params.ConfigName, in.Name) &&
			assert.Equal(t, params.AuthToken, in.Token) &&
			assert.Equal(t, host, in.Host) &&
			assert.Equal(t, params.Org, in.Org)
	})).DoAndReturn(func(in config.Config) (config.Config, error) {
		return in, nil
	})

	stdio := mock.NewMockStdIO(ctrl)
	bytesWritten := bytes.Buffer{}
	stdio.EXPECT().Write(gomock.Any()).DoAndReturn(bytesWritten.Write).AnyTimes()
	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio},
		SetupApi: client,
	}
	require.NoError(t, cli.Setup(context.Background(), &params))
	testutils.MatchLines(t, []string{
		`User\s+Organization\s+Bucket`,
		fmt.Sprintf(`%s\s+%s\s+%s`, params.Username, params.Org, params.Bucket),
	}, strings.Split(bytesWritten.String(), "\n"))
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
	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{Allowed: api.PtrBool(true)}, nil)
	client.EXPECT().PostSetup(gomock.Any()).Return(api.ApiPostSetupRequest{ApiService: client})
	client.EXPECT().PostSetupExecute(tmock.MatchedBy(func(in api.ApiPostSetupRequest) bool {
		body := in.GetOnboardingRequest()
		return assert.NotNil(t, body) &&
			assert.Equal(t, username, body.Username) &&
			assert.Equal(t, password, *body.Password) &&
			assert.Nil(t, body.Token) &&
			assert.Equal(t, org, body.Org) &&
			assert.Equal(t, bucket, body.Bucket) &&
			assert.Equal(t, retentionSecs, *body.RetentionPeriodSeconds)
	})).Return(resp, nil)

	host := "fake-host"
	configSvc := mock.NewMockConfigService(ctrl)
	configSvc.EXPECT().ListConfigs().Return(nil, nil)
	configSvc.EXPECT().CreateConfig(tmock.MatchedBy(func(in config.Config) bool {
		return assert.Equal(t, config.DefaultConfig.Name, in.Name) &&
			assert.Equal(t, token, in.Token) &&
			assert.Equal(t, host, in.Host) &&
			assert.Equal(t, org, in.Org)
	})).DoAndReturn(func(in config.Config) (config.Config, error) {
		return in, nil
	})

	stdio := mock.NewMockStdIO(ctrl)
	bytesWritten := bytes.Buffer{}
	stdio.EXPECT().Write(gomock.Any()).DoAndReturn(bytesWritten.Write).AnyTimes()
	stdio.EXPECT().Banner(gomock.Any())
	stdio.EXPECT().GetStringInput(gomock.Eq("Please type your primary username"), gomock.Any()).Return(username, nil)
	stdio.EXPECT().GetPassword(gomock.Eq("Please type your password"), gomock.Any()).Return(password, nil)
	stdio.EXPECT().GetPassword(gomock.Eq("Please type your password again"), gomock.Any()).Return(password, nil)
	stdio.EXPECT().GetStringInput(gomock.Eq("Please type your primary organization name"), gomock.Any()).Return(org, nil)
	stdio.EXPECT().GetStringInput("Please type your primary bucket name", gomock.Any()).Return(bucket, nil)
	stdio.EXPECT().GetStringInput("Please type your retention period in hours, or 0 for infinite", gomock.Any()).Return(strconv.Itoa(retentionHrs), nil)
	stdio.EXPECT().GetConfirm(gomock.Any()).Return(true)
	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio},
		SetupApi: client,
	}
	require.NoError(t, cli.Setup(context.Background(), &setup.Params{}))
	testutils.MatchLines(t, []string{
		`User\s+Organization\s+Bucket`,
		fmt.Sprintf(`%s\s+%s\s+%s`, username, org, bucket),
	}, strings.Split(bytesWritten.String(), "\n"))
}

func Test_SetupPasswordParamToShort(t *testing.T) {
	t.Parallel()

	retentionSecs := int64(duration.Week.Seconds())
	params := setup.Params{
		Username:  "user",
		Password:  "2short",
		AuthToken: "mytoken",
		Org:       "org",
		Bucket:    "bucket",
		Retention: fmt.Sprintf("%ds", retentionSecs),
		Force:     false,
	}

	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{Allowed: api.PtrBool(true)}, nil)

	host := "fake-host"
	configSvc := mock.NewMockConfigService(ctrl)
	configSvc.EXPECT().ListConfigs().Return(nil, nil)

	stdio := mock.NewMockStdIO(ctrl)
	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio},
		SetupApi: client,
	}
	err := cli.Setup(context.Background(), &params)
	require.Equal(t, clients.ErrPasswordIsTooShort, err)
}

func Test_SetupCancelAtConfirmation(t *testing.T) {
	t.Parallel()

	retentionSecs := int64(duration.Week.Seconds())
	params := setup.Params{
		Username:  "user",
		Password:  "mysecretpassword",
		AuthToken: "mytoken",
		Org:       "org",
		Bucket:    "bucket",
		Retention: fmt.Sprintf("%ds", retentionSecs),
		Force:     false,
	}

	ctrl := gomock.NewController(t)
	client := mock.NewMockSetupApi(ctrl)
	client.EXPECT().GetSetup(gomock.Any()).Return(api.ApiGetSetupRequest{ApiService: client})
	client.EXPECT().GetSetupExecute(gomock.Any()).Return(api.InlineResponse200{Allowed: api.PtrBool(true)}, nil)

	host := "fake-host"
	configSvc := mock.NewMockConfigService(ctrl)
	configSvc.EXPECT().ListConfigs().Return(nil, nil)

	stdio := mock.NewMockStdIO(ctrl)
	stdio.EXPECT().Banner(gomock.Any())
	stdio.EXPECT().GetConfirm(gomock.Any()).Return(false)

	cli := setup.Client{
		CLI:      clients.CLI{ConfigService: configSvc, ActiveConfig: config.Config{Host: host}, StdIO: stdio},
		SetupApi: client,
	}
	err := cli.Setup(context.Background(), &params)
	require.Equal(t, setup.ErrSetupCanceled, err)
}
