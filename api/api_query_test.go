package api_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/stretchr/testify/require"
)

var exampleResponse = `result,table,_start,_stop,_time,region,host,_value
mean,0,2018-05-08T20:50:00Z,2018-05-08T20:51:00Z,2018-05-08T20:50:00Z,east,A,15.43
mean,0,2018-05-08T20:50:00Z,2018-05-08T20:51:00Z,2018-05-08T20:50:20Z,east,B,59.25
mean,0,2018-05-08T20:50:00Z,2018-05-08T20:51:00Z,2018-05-08T20:50:40Z,east,C,52.62`

func setupServer(t *testing.T) (*httptest.Server, api.QueryApi) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		isGzip := req.Header.Get("Accept-Encoding") == "gzip"
		if isGzip {
			rw.Header().Set("Content-Encoding", "gzip")
		}
		rw.WriteHeader(200)

		if isGzip {
			buf := bytes.Buffer{}
			gzw := gzip.NewWriter(&buf)
			if _, err := gzw.Write([]byte(exampleResponse)); err != nil {
				_ = gzw.Close()
				t.Fatalf("unexpected error: %v", err)
			}
			require.NoError(t, gzw.Close())
			_, err := io.Copy(rw, &buf)
			require.NoError(t, err)
		} else {
			_, err := rw.Write([]byte(exampleResponse))
			require.NoError(t, err)
		}
	}))

	config := api.NewConfiguration()
	config.Scheme = "http"
	config.Host = server.Listener.Addr().String()
	client := api.NewAPIClient(config)

	return server, client.QueryApi
}

func TestQuery_NoGzip(t *testing.T) {
	t.Parallel()

	server, client := setupServer(t)
	defer server.Close()

	resp, err := client.PostQuery(context.Background()).
		AcceptEncoding("identity").
		Query(api.Query{}).
		Execute()
	require.NoError(t, err)
	defer resp.Close()

	out := bytes.Buffer{}
	_, err = io.Copy(&out, resp)
	require.NoError(t, err)

	require.Equal(t, exampleResponse, out.String())
}

func TestQuery_Gzip(t *testing.T) {
	t.Parallel()

	server, client := setupServer(t)
	defer server.Close()

	resp, err := client.PostQuery(context.Background()).
		AcceptEncoding("gzip").
		Query(api.Query{}).
		Execute()
	require.NoError(t, err)
	defer resp.Close()

	out := bytes.Buffer{}
	_, err = io.Copy(&out, resp)
	require.NoError(t, err)

	require.Equal(t, exampleResponse, out.String())
}
