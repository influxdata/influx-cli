package gzip

import (
	"compress/gzip"
	"io"
)

var _ io.ReadCloser = (*gzipPipe)(nil)

type gzipPipe struct {
	underlying io.ReadCloser
	pipeOut    io.ReadCloser
}

// NewGzipPipe returns an io.ReadCloser that wraps an input data stream,
// applying gzip compression to the underlying data on Read and closing the
// underlying data on Close.
func NewGzipPipe(in io.ReadCloser) *gzipPipe {
	pr, pw := io.Pipe()
	gw := gzip.NewWriter(pw)

	go func() {
		_, err := io.Copy(gw, in)
		gw.Close()
		if err != nil {
			pw.CloseWithError(err)
		} else {
			pw.Close()
		}
	}()

	return &gzipPipe{underlying: in, pipeOut: pr}
}

func (gzp gzipPipe) Read(p []byte) (int, error) {
	return gzp.pipeOut.Read(p)
}

func (gzp gzipPipe) Close() error {
	if err := gzp.pipeOut.Close(); err != nil {
		return err
	}
	if err := gzp.underlying.Close(); err != nil {
		return err
	}
	return nil
}
