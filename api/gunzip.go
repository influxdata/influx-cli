package api

import (
	"io"
	"net/http"

	"github.com/influxdata/influx-cli/v2/pkg/gzip"
)

func GunzipIfNeeded(resp *http.Response) (io.ReadCloser, error) {
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewGunzipReadCloser(resp.Body)
		if err != nil {
			resp.Body.Close()
			return nil, err
		}
		return reader, nil
	}
	return resp.Body, nil
}
