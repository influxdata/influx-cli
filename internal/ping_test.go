package internal_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/stretchr/testify/require"
)

type testClient struct {
	GetHealthFn func(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*api.GetHealthResponse, error)
}

func (tc *testClient) GetHealthWithResponse(ctx context.Context, p *api.GetHealthParams, fns ...api.RequestEditorFn) (*api.GetHealthResponse, error) {
	return tc.GetHealthFn(ctx, p, fns...)
}

func Test_PingSuccessNoTracing(t *testing.T) {
	t.Parallel()

	client := &testClient{
		GetHealthFn: func(_ context.Context, p *api.GetHealthParams, _ ...api.RequestEditorFn) (*api.GetHealthResponse, error) {
			require.Nil(t, p.ZapTraceSpan)
			return &api.GetHealthResponse{
				JSON200: &api.HealthCheck{Name: "test", Status: api.HealthCheckStatus_pass},
			}, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	require.NoError(t, cli.Ping(context.Background(), client))
	require.Equal(t, "OK\n", out.String())
}

func Test_PingSuccessWithTracing(t *testing.T) {
	t.Parallel()

	traceId := api.TraceSpan("trace-id")
	client := &testClient{
		GetHealthFn: func(_ context.Context, p *api.GetHealthParams, _ ...api.RequestEditorFn) (*api.GetHealthResponse, error) {
			require.Equal(t, traceId, *p.ZapTraceSpan)
			return &api.GetHealthResponse{
				JSON200: &api.HealthCheck{Name: "test", Status: api.HealthCheckStatus_pass},
			}, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out, TraceId: &traceId}

	require.NoError(t, cli.Ping(context.Background(), client))
	require.Equal(t, "OK\n", out.String())
}

func Test_PingFailedRequest(t *testing.T) {
	t.Parallel()

	e := "the internet is down"
	client := &testClient{
		GetHealthFn: func(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*api.GetHealthResponse, error) {
			return nil, errors.New(e)
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
		GetHealthFn: func(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*api.GetHealthResponse, error) {
			return &api.GetHealthResponse{
				JSON503: &api.HealthCheck{Name: "test", Status: api.HealthCheckStatus_fail, Message: &e},
			}, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
	require.Empty(t, out.String())
}

func Test_PingFailedUnhandledError(t *testing.T) {
	t.Parallel()

	e := "something went boom"
	client := &testClient{
		GetHealthFn: func(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*api.GetHealthResponse, error) {
			return &api.GetHealthResponse{
				JSONDefault: &api.Error{
					Code:    api.ErrorCode_internal_error,
					Message: e,
				},
			}, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
	require.Empty(t, out.String())
}

func Test_PingFailedCheck(t *testing.T) {
	t.Parallel()

	e := "oops, forgot to set the status code"
	client := &testClient{
		GetHealthFn: func(_ context.Context, p *api.GetHealthParams, _ ...api.RequestEditorFn) (*api.GetHealthResponse, error) {
			return &api.GetHealthResponse{
				JSON200: &api.HealthCheck{Name: "test", Status: api.HealthCheckStatus_fail, Message: &e},
			}, nil
		},
	}

	out := &bytes.Buffer{}
	cli := &internal.CLI{Stdout: out}

	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
	require.Empty(t, out.String())
}
