package gzip_test

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	pgzip "github.com/influxdata/influx-cli/v2/pkg/gzip"
	"github.com/stretchr/testify/require"
)

func TestGzipPipe(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		data := strings.Repeat("Data data I'm some data\n", 1024)
		reader := strings.NewReader(data)
		pipe := pgzip.NewGzipPipe(ioutil.NopCloser(reader))
		defer pipe.Close()
		gunzip, err := gzip.NewReader(pipe)
		require.NoError(t, err)
		defer gunzip.Close()

		out := bytes.Buffer{}
		_, err = io.Copy(&out, gunzip)
		require.NoError(t, err)

		require.Equal(t, data, out.String())
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		reader := &failingReader{n: 3, err: errors.New("I BROKE")}
		pipe := pgzip.NewGzipPipe(ioutil.NopCloser(reader))
		defer pipe.Close()
		gunzip, err := gzip.NewReader(pipe)
		require.NoError(t, err)
		defer gunzip.Close()

		out := bytes.Buffer{}
		_, err = io.Copy(&out, gunzip)
		require.Error(t, err)
		require.Equal(t, reader.err, err)
	})
}

type failingReader struct {
	n int
	err error
}

func (frc *failingReader) Read(p []byte) (int, error) {
	if frc.n <= 0 {
		return 0, frc.err
	}
	frc.n--
	p[0] = 'a'
	return 1, nil
}
