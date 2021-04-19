package internal_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/Netflix/go-expect"
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
		GetHealthExecuteFn: func(req api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			require.Nil(t, req.GetZapTraceSpan())
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_PASS}, nil, nil
		},
	}

	tc, err := expect.NewConsole()
	require.NoError(t, err)
	defer tc.Close()
	cli := &internal.CLI{Stdout: tc.Tty()}

	require.NoError(t, cli.Ping(context.Background(), client))
	_, err = tc.ExpectString("OK")
	require.NoError(t, err)
}

func Test_PingSuccessWithTracing(t *testing.T) {
	t.Parallel()

	traceId := "trace-id"
	client := &testClient{
		GetHealthExecuteFn: func(req api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			require.NotNil(t, req.GetZapTraceSpan())
			require.Equal(t, traceId, *req.GetZapTraceSpan())
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_PASS}, nil, nil
		},
	}

	tc, err := expect.NewConsole()
	require.NoError(t, err)
	defer tc.Close()
	cli := &internal.CLI{Stdout: tc.Tty(), TraceId: traceId}

	require.NoError(t, cli.Ping(context.Background(), client))
	_, err = tc.ExpectString("OK")
	require.NoError(t, err)
}

func Test_PingFailedRequest(t *testing.T) {
	t.Parallel()

	e := "the internet is down"
	client := &testClient{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{}, nil, errors.New(e)
		},
	}

	cli := &internal.CLI{}
	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_PingFailedStatus(t *testing.T) {
	t.Parallel()

	e := "I broke"
	client := &testClient{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_FAIL, Message: &e}, nil, nil
		},
	}

	cli := &internal.CLI{}
	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_PingFailedStatusNoMessage(t *testing.T) {
	t.Parallel()

	name := "foo"
	client := &testClient{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_FAIL, Name: name}, nil, nil
		},
	}

	cli := &internal.CLI{}
	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), name)
}
