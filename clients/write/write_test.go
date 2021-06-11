package write_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/write"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
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
	t.Parallel()

	inLines := []string{"fake line protocol 1", "fake line protocol 2", "fake line protocol 3"}
	mockReader := bufferReader{}
	for _, l := range inLines {
		_, err := io.Copy(&mockReader.buf, strings.NewReader(l+"\n"))
		require.NoError(t, err)
	}
	mockThrottler := noopThrottler{}
	mockBatcher := lineBatcher{}

	params := write.Params{
		OrgID:     "12345",
		BucketID:  "98765",
		Precision: api.WRITEPRECISION_S,
	}

	ctrl := gomock.NewController(t)
	client := mock.NewMockWriteApi(ctrl)
	var writtenLines []string
	client.EXPECT().PostWrite(gomock.Any()).Return(api.ApiPostWriteRequest{ApiService: client}).Times(len(inLines))
	client.EXPECT().PostWriteExecute(tmock.MatchedBy(func(in api.ApiPostWriteRequest) bool {
		return assert.Equal(t, params.OrgID, *in.GetOrg()) &&
			assert.Equal(t, params.BucketID, *in.GetBucket()) &&
			assert.Equal(t, params.Precision, *in.GetPrecision()) &&
			assert.Equal(t, "gzip", *in.GetContentEncoding())
	})).DoAndReturn(func(in api.ApiPostWriteRequest) error {
		bodyBytes := bytes.NewReader(in.GetBody())
		gzr, err := gzip.NewReader(bodyBytes)
		require.NoError(t, err)
		defer gzr.Close()
		buf := bytes.Buffer{}
		_, err = buf.ReadFrom(gzr)
		require.NoError(t, err)
		writtenLines = append(writtenLines, buf.String())
		return nil
	}).Times(len(inLines))

	cli := write.Client{
		CLI:         clients.CLI{ActiveConfig: config.Config{Org: "my-default-org"}},
		LineReader:  &mockReader,
		RateLimiter: &mockThrottler,
		BatchWriter: &mockBatcher,
		WriteApi:    client,
	}

	require.NoError(t, cli.Write(context.Background(), &params))
	require.Equal(t, inLines, writtenLines)
	require.True(t, mockThrottler.used)
}

func TestWriteByNames(t *testing.T) {
	t.Parallel()

	inLines := []string{"fake line protocol 1", "fake line protocol 2", "fake line protocol 3"}
	mockReader := bufferReader{}
	for _, l := range inLines {
		_, err := io.Copy(&mockReader.buf, strings.NewReader(l+"\n"))
		require.NoError(t, err)
	}
	mockThrottler := noopThrottler{}
	mockBatcher := lineBatcher{}

	params := write.Params{
		OrgName:    "my-org",
		BucketName: "my-bucket",
		Precision:  api.WRITEPRECISION_US,
	}

	ctrl := gomock.NewController(t)
	client := mock.NewMockWriteApi(ctrl)
	var writtenLines []string
	client.EXPECT().PostWrite(gomock.Any()).Return(api.ApiPostWriteRequest{ApiService: client}).Times(len(inLines))
	client.EXPECT().PostWriteExecute(tmock.MatchedBy(func(in api.ApiPostWriteRequest) bool {
		return assert.Equal(t, params.OrgName, *in.GetOrg()) &&
			assert.Equal(t, params.BucketName, *in.GetBucket()) &&
			assert.Equal(t, params.Precision, *in.GetPrecision()) &&
			assert.Equal(t, "gzip", *in.GetContentEncoding())
	})).DoAndReturn(func(in api.ApiPostWriteRequest) error {
		bodyBytes := bytes.NewReader(in.GetBody())
		gzr, err := gzip.NewReader(bodyBytes)
		require.NoError(t, err)
		defer gzr.Close()
		buf := bytes.Buffer{}
		_, err = buf.ReadFrom(gzr)
		require.NoError(t, err)
		writtenLines = append(writtenLines, buf.String())
		return nil
	}).Times(len(inLines))

	cli := write.Client{
		CLI:         clients.CLI{ActiveConfig: config.Config{Org: "my-default-org"}},
		LineReader:  &mockReader,
		RateLimiter: &mockThrottler,
		BatchWriter: &mockBatcher,
		WriteApi:    client,
	}

	require.NoError(t, cli.Write(context.Background(), &params))
	require.Equal(t, inLines, writtenLines)
	require.True(t, mockThrottler.used)
}

func TestWriteOrgFromConfig(t *testing.T) {
	t.Parallel()

	inLines := []string{"fake line protocol 1", "fake line protocol 2", "fake line protocol 3"}
	mockReader := bufferReader{}
	for _, l := range inLines {
		_, err := io.Copy(&mockReader.buf, strings.NewReader(l+"\n"))
		require.NoError(t, err)
	}
	mockThrottler := noopThrottler{}
	mockBatcher := lineBatcher{}

	params := write.Params{
		BucketName: "my-bucket",
		Precision:  api.WRITEPRECISION_US,
	}

	defaultOrg := "my-default-org"
	ctrl := gomock.NewController(t)
	client := mock.NewMockWriteApi(ctrl)
	var writtenLines []string
	client.EXPECT().PostWrite(gomock.Any()).Return(api.ApiPostWriteRequest{ApiService: client}).Times(len(inLines))
	client.EXPECT().PostWriteExecute(tmock.MatchedBy(func(in api.ApiPostWriteRequest) bool {
		return assert.Equal(t, defaultOrg, *in.GetOrg()) &&
			assert.Equal(t, params.BucketName, *in.GetBucket()) &&
			assert.Equal(t, params.Precision, *in.GetPrecision()) &&
			assert.Equal(t, "gzip", *in.GetContentEncoding()) // Make sure the body is properly marked for compression.
	})).DoAndReturn(func(in api.ApiPostWriteRequest) error {
		bodyBytes := bytes.NewReader(in.GetBody())
		gzr, err := gzip.NewReader(bodyBytes)
		require.NoError(t, err)
		defer gzr.Close()
		buf := bytes.Buffer{}
		_, err = buf.ReadFrom(gzr)
		require.NoError(t, err)
		writtenLines = append(writtenLines, buf.String())
		return nil
	}).Times(len(inLines))

	cli := write.Client{
		CLI:         clients.CLI{ActiveConfig: config.Config{Org: defaultOrg}},
		LineReader:  &mockReader,
		RateLimiter: &mockThrottler,
		BatchWriter: &mockBatcher,
		WriteApi:    client,
	}

	require.NoError(t, cli.Write(context.Background(), &params))
	require.Equal(t, inLines, writtenLines)
	require.True(t, mockThrottler.used)
}
