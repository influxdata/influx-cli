package api

import (
	"compress/gzip"
	"io"
	"net/http"
)

func GunzipIfNeeded(resp *http.Response) (io.ReadCloser, error) {
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return resp.Body, err
		}
		return &gunzipReadCloser{underlying: resp.Body, gunzip: gzr}, nil
	}
	return resp.Body, nil
}

type gunzipReadCloser struct {
	underlying io.ReadCloser
	gunzip     io.ReadCloser
}

func (gzrc *gunzipReadCloser) Read(p []byte) (int, error) {
	return gzrc.gunzip.Read(p)
}

func (gzrc *gunzipReadCloser) Close() error {
	if err := gzrc.gunzip.Close(); err != nil {
		return err
	}
	return gzrc.underlying.Close()
}
