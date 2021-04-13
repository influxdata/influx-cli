package ping_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/internal/ping"
	"github.com/influxdata/influx-cli/v2/kit/tracing"
	"github.com/stretchr/testify/require"
)

type testClient struct {
	GetHealthFn func(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*http.Response, error)
}

func (tc *testClient) GetHealth(ctx context.Context, p *api.GetHealthParams, fns ...api.RequestEditorFn) (*http.Response, error) {
	return tc.GetHealthFn(ctx, p, fns...)
}

// Default context, no trace ID
var ctx = tracing.WrapContext(context.Background(), "")

func Test_PingSuccessNoTracing(t *testing.T) {
	t.Parallel()

	client := &testClient{
		GetHealthFn: func(_ context.Context, p *api.GetHealthParams, _ ...api.RequestEditorFn) (*http.Response, error) {
			require.Nil(t, p.ZapTraceSpan)
			respJson := `{"name":"test","status":"pass"}`
			return &http.Response{
				StatusCode: 200,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       ioutil.NopCloser(bytes.NewBufferString(respJson)),
			}, nil
		},
	}

	require.NoError(t, ping.Ping(ctx, client))
}

func Test_PingSuccessWithTracing(t *testing.T) {
	t.Parallel()

	traceId := "trace-id"
	client := &testClient{
		GetHealthFn: func(_ context.Context, p *api.GetHealthParams, _ ...api.RequestEditorFn) (*http.Response, error) {
			require.Equal(t, api.TraceSpan(traceId), *p.ZapTraceSpan)
			respJson := `{"name":"test","status":"pass"}`
			return &http.Response{
				StatusCode: 200,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       ioutil.NopCloser(bytes.NewBufferString(respJson)),
			}, nil
		},
	}

	ctx := tracing.WrapContext(context.Background(), traceId)
	require.NoError(t, ping.Ping(ctx, client))
}

func Test_PingFailedRequest(t *testing.T) {
	t.Parallel()

	e := "the internet is down"
	client := &testClient{
		GetHealthFn: func(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*http.Response, error) {
			return nil, errors.New(e)
		},
	}

	err := ping.Ping(ctx, client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_PingFailedStatus(t *testing.T) {
	t.Parallel()

	client := &testClient{
		GetHealthFn: func(_ context.Context, p *api.GetHealthParams, _ ...api.RequestEditorFn) (*http.Response, error) {
			require.Nil(t, p.ZapTraceSpan)
			respJson := `{"name":"test","status":"fail","message":"I broke"}`
			return &http.Response{
				StatusCode: 503,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       ioutil.NopCloser(bytes.NewBufferString(respJson)),
			}, nil
		},
	}

	err := ping.Ping(ctx, client)
	require.Error(t, err)
	require.Contains(t, err.Error(), "I broke")
}

func Test_PingBadBody(t *testing.T) {
	t.Parallel()

	e := "server doesn't know it's supposed to return JSON"
	client := &testClient{
		GetHealthFn: func(context.Context, *api.GetHealthParams, ...api.RequestEditorFn) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(e)),
			}, nil
		},
	}

	err := ping.Ping(ctx, client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_PingFailedCheck(t *testing.T) {
	t.Parallel()

	client := &testClient{
		GetHealthFn: func(_ context.Context, p *api.GetHealthParams, _ ...api.RequestEditorFn) (*http.Response, error) {
			require.Nil(t, p.ZapTraceSpan)
			respJson := `{"name":"test","status":"fail","message":"I broke"}`
			return &http.Response{
				StatusCode: 200,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       ioutil.NopCloser(bytes.NewBufferString(respJson)),
			}, nil
		},
	}

	err := ping.Ping(ctx, client)
	require.Error(t, err)
	require.Contains(t, err.Error(), "I broke")
}
