package gzip

import (
	"compress/gzip"
	"io"
)

func NewGunzipReadCloser(in io.ReadCloser) (*gunzipReadCloser, error) {
	gzr, err := gzip.NewReader(in)
	if err != nil {
		return nil, err
	}
	return &gunzipReadCloser{underlying: in, gunzip: gzr}, nil
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
