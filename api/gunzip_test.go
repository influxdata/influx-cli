package api_test

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/stretchr/testify/require"
)

var exampleResponse = `result,table,_start,_stop,_time,region,host,_value
mean,0,2018-05-08T20:50:00Z,2018-05-08T20:51:00Z,2018-05-08T20:50:00Z,east,A,15.43
mean,0,2018-05-08T20:50:00Z,2018-05-08T20:51:00Z,2018-05-08T20:50:20Z,east,B,59.25
mean,0,2018-05-08T20:50:00Z,2018-05-08T20:51:00Z,2018-05-08T20:50:40Z,east,C,52.62`

func TestGunzipIfNeeded(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		encoding string
		getBody  func(*testing.T) io.ReadCloser
	}{
		{
			name: "no gzip",
			getBody: func(t *testing.T) io.ReadCloser {
				return ioutil.NopCloser(strings.NewReader(exampleResponse))
			},
		},
		{
			name:     "gzip",
			encoding: "gzip",
			getBody: func(t *testing.T) io.ReadCloser {
				pr, pw := io.Pipe()
				gw := gzip.NewWriter(pw)

				go func() {
					_, err := io.Copy(gw, strings.NewReader(exampleResponse))
					gw.Close()
					pw.Close()
					require.NoError(t, err)
				}()

				return pr
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			headers := http.Header{}
			if tc.encoding != "" {
				headers.Add("Content-Encoding", tc.encoding)
			}
			resp := http.Response{
				Header: headers,
				Body:   tc.getBody(t),
			}

			raw, err := api.GunzipIfNeeded(&resp)
			require.NoError(t, err)
			defer raw.Close()
			body, err := ioutil.ReadAll(raw)
			require.NoError(t, err)
			require.Equal(t, exampleResponse, string(body))
		})
	}
}
