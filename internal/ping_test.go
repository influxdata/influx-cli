package internal_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/stretchr/testify/require"
)

type testClient struct {
	GetHealthExecuteFn func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error)
}

func (tc *testClient) GetHealth(context.Context) api.ApiGetHealthRequest {
	return api.ApiGetHealthRequest{
		ApiService: tc,
	}
}

func (tc *testClient) GetHealthExecute(req api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
	return tc.GetHealthExecuteFn(req)
}

func Test_PingSuccess(t *testing.T) {
	t.Parallel()

	client := &testClient{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_PASS}, nil, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	require.NoError(t, cli.Ping(context.Background(), client))
	require.Equal(t, "OK\n", out.String())
}

func Test_PingFailedRequest(t *testing.T) {
	t.Parallel()

	e := "the internet is down"
	client := &testClient{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{}, nil, errors.New(e)
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
	require.Empty(t, out.String())
}

func Test_PingFailedStatus(t *testing.T) {
	t.Parallel()

	e := "I broke"
	client := &testClient{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_FAIL, Message: &e}, nil, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
	require.Empty(t, out.String())
}

func Test_PingFailedStatusNoMessage(t *testing.T) {
	t.Parallel()

	name := "foo"
	client := &testClient{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_FAIL, Name: name}, nil, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), name)
	require.Empty(t, out.String())
}
