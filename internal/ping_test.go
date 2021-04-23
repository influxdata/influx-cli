package internal_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/require"
)

func Test_PingSuccess(t *testing.T) {
	t.Parallel()

	client := &mock.HealthApi{
		GetHealthExecuteFn: func(req api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			require.Nil(t, req.GetZapTraceSpan())
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_PASS}, nil, nil
		},
	}

	stdio := mock.NewMockStdio(nil, true)
	cli := &internal.CLI{StdIO: stdio}

	require.NoError(t, cli.Ping(context.Background(), client))
	require.Equal(t, "OK\n", stdio.Stdout())
}

func Test_PingSuccessWithTracing(t *testing.T) {
	t.Parallel()

	traceId := "trace-id"
	client := &mock.HealthApi{
		GetHealthExecuteFn: func(req api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			require.NotNil(t, req.GetZapTraceSpan())
			require.Equal(t, traceId, *req.GetZapTraceSpan())
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_PASS}, nil, nil
		},
	}

	stdio := mock.NewMockStdio(nil, true)
	cli := &internal.CLI{TraceId: traceId, StdIO: stdio}

	require.NoError(t, cli.Ping(context.Background(), client))
	require.Equal(t, "OK\n", stdio.Stdout())
}

func Test_PingFailedRequest(t *testing.T) {
	t.Parallel()

	e := "the internet is down"
	client := &mock.HealthApi{
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
	client := &mock.HealthApi{
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
	client := &mock.HealthApi{
		GetHealthExecuteFn: func(api.ApiGetHealthRequest) (api.HealthCheck, *http.Response, error) {
			return api.HealthCheck{Status: api.HEALTHCHECKSTATUS_FAIL, Name: name}, nil, nil
		},
	}

	cli := &internal.CLI{}
	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), name)
}
