package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoGzipRequest(t *testing.T) {
	client := APIClient{cfg: NewConfiguration()}
	body := []byte("This should get gzipped")
	req, err := client.prepareRequest(
		context.Background(),
		"/foo", "POST", body,
		map[string]string{},
		nil, nil, "", "", nil,
	)
	require.NoError(t, err)
	defer req.Body.Close()

	out := bytes.Buffer{}
	_, err = io.Copy(&out, req.Body)
	require.NoError(t, err)

	require.Equal(t, string(body), out.String())
}

func TestGzipRequest(t *testing.T) {
	client := APIClient{cfg: NewConfiguration()}
	body := []byte("This should get gzipped")
	req, err := client.prepareRequest(
		context.Background(),
		"/foo", "POST", body,
		map[string]string{"Content-Encoding": "gzip"},
		nil, nil, "", "", nil,
	)
	require.NoError(t, err)
	defer req.Body.Close()

	out := bytes.Buffer{}
	gzr, err := gzip.NewReader(req.Body)
	require.NoError(t, err)
	defer gzr.Close()
	_, err = io.Copy(&out, gzr)
	require.NoError(t, err)

	require.Equal(t, string(body), out.String())
}
