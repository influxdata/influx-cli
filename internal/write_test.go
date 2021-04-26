package internal_test

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/require"
)

type bufferReader struct {
	buf bytes.Buffer
}

func (pr *bufferReader) Open(context.Context) (io.Reader, io.Closer, error) {
	return &pr.buf, ioutil.NopCloser(nil), nil
}

type noopThrottler struct {
	used bool
}

func (nt *noopThrottler) Throttle(_ context.Context, in io.Reader) io.Reader {
	nt.used = true
	return in
}

type lineBatcher struct{}

func (pb *lineBatcher) WriteBatches(_ context.Context, r io.Reader, writeFn func(batch []byte) error) error {
	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, r); err != nil {
		return err
	}
	for _, l := range strings.Split(buf.String(), "\n") {
		if l != "" {
			if err := writeFn([]byte(l)); err != nil {
				return err
			}
		}
	}
	return nil
}

func TestWriteByIDs(t *testing.T) {
	inLines := []string{"fake line protocol 1", "fake line protocol 2", "fake line protocol 3"}
	mockReader := bufferReader{}
	for _, l := range inLines {
		_, err := io.Copy(&mockReader.buf, strings.NewReader(l+"\n"))
		require.NoError(t, err)
	}
	mockThrottler := noopThrottler{}
	mockBatcher := lineBatcher{}

	params := internal.WriteParams{
		OrgID:     "12345",
		BucketID:  "98765",
		Precision: api.WRITEPRECISION_S,
	}
	cli := internal.CLI{ActiveConfig: config.Config{Org: "my-default-org"}}

	var writtenLines []string
	client := mock.WriteApi{
		PostWriteExecuteFn: func(req api.ApiPostWriteRequest) (*http.Response, error) {
			// Make sure query params are set.
			require.Equal(t, params.OrgID, *req.GetOrg())
			require.Equal(t, params.BucketID, *req.GetBucket())
			require.Equal(t, params.Precision, *req.GetPrecision())

			// Make sure the body is properly marked for compression, and record what was sent.
			require.Equal(t, "gzip", *req.GetContentEncoding())
			writtenLines = append(writtenLines, string(req.GetBody()))
			return nil, nil
		},
	}

	clients := internal.WriteClients{
		Reader:    &mockReader,
		Throttler: &mockThrottler,
		Writer:    &mockBatcher,
		Client:    &client,
	}

	require.NoError(t, cli.Write(context.Background(), &clients, &params))
	require.Equal(t, inLines, writtenLines)
	require.True(t, mockThrottler.used)
}

func TestWriteByNames(t *testing.T) {
	inLines := []string{"fake line protocol 1", "fake line protocol 2", "fake line protocol 3"}
	mockReader := bufferReader{}
	for _, l := range inLines {
		_, err := io.Copy(&mockReader.buf, strings.NewReader(l+"\n"))
		require.NoError(t, err)
	}
	mockThrottler := noopThrottler{}
	mockBatcher := lineBatcher{}

	params := internal.WriteParams{
		OrgName:    "my-org",
		BucketName: "my-bucket",
		Precision:  api.WRITEPRECISION_US,
	}
	cli := internal.CLI{TraceId: "my-trace-id", ActiveConfig: config.Config{Org: "my-default-org"}}

	var writtenLines []string
	client := mock.WriteApi{
		PostWriteExecuteFn: func(req api.ApiPostWriteRequest) (*http.Response, error) {
			// Make sure query params are set.
			require.Equal(t, params.OrgName, *req.GetOrg())
			require.Equal(t, params.BucketName, *req.GetBucket())
			require.Equal(t, params.Precision, *req.GetPrecision())
			require.Equal(t, cli.TraceId, *req.GetZapTraceSpan())

			// Make sure the body is properly marked for compression, and record what was sent.
			require.Equal(t, "gzip", *req.GetContentEncoding())
			writtenLines = append(writtenLines, string(req.GetBody()))
			return nil, nil
		},
	}

	clients := internal.WriteClients{
		Reader:    &mockReader,
		Throttler: &mockThrottler,
		Writer:    &mockBatcher,
		Client:    &client,
	}

	require.NoError(t, cli.Write(context.Background(), &clients, &params))
	require.Equal(t, inLines, writtenLines)
	require.True(t, mockThrottler.used)
}

func TestWriteOrgFromConfig(t *testing.T) {
	inLines := []string{"fake line protocol 1", "fake line protocol 2", "fake line protocol 3"}
	mockReader := bufferReader{}
	for _, l := range inLines {
		_, err := io.Copy(&mockReader.buf, strings.NewReader(l+"\n"))
		require.NoError(t, err)
	}
	mockThrottler := noopThrottler{}
	mockBatcher := lineBatcher{}

	params := internal.WriteParams{
		BucketName: "my-bucket",
		Precision:  api.WRITEPRECISION_US,
	}
	cli := internal.CLI{ActiveConfig: config.Config{Org: "my-default-org"}}

	var writtenLines []string
	client := mock.WriteApi{
		PostWriteExecuteFn: func(req api.ApiPostWriteRequest) (*http.Response, error) {
			// Make sure query params are set.
			require.Equal(t, cli.ActiveConfig.Org, *req.GetOrg())
			require.Equal(t, params.BucketName, *req.GetBucket())
			require.Equal(t, params.Precision, *req.GetPrecision())

			// Make sure the body is properly marked for compression, and record what was sent.
			require.Equal(t, "gzip", *req.GetContentEncoding())
			writtenLines = append(writtenLines, string(req.GetBody()))
			return nil, nil
		},
	}

	clients := internal.WriteClients{
		Reader:    &mockReader,
		Throttler: &mockThrottler,
		Writer:    &mockBatcher,
		Client:    &client,
	}

	require.NoError(t, cli.Write(context.Background(), &clients, &params))
	require.Equal(t, inLines, writtenLines)
	require.True(t, mockThrottler.used)
}

func TestWriteDryRun(t *testing.T) {
	inLines := `
fake line protocol 1
fake line protocol 2
fake line protocol 3
`
	mockReader := bufferReader{}
	_, err := io.Copy(&mockReader.buf, strings.NewReader(inLines))
	require.NoError(t, err)
	stdio := mock.NewMockStdio(nil, true)
	cli := internal.CLI{ActiveConfig: config.Config{Org: "my-default-org"}, StdIO: stdio}

	require.NoError(t, cli.WriteDryRun(context.Background(), &mockReader))
	require.Equal(t, inLines, stdio.Stdout())
}
