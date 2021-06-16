package write_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/write"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/require"
)

func TestWriteDryRun(t *testing.T) {
	t.Parallel()

	inLines := `
fake line protocol 1
fake line protocol 2
fake line protocol 3
`
	mockReader := bufferReader{}
	_, err := io.Copy(&mockReader.buf, strings.NewReader(inLines))
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	stdio := mock.NewMockStdIO(ctrl)
	bytesWritten := bytes.Buffer{}
	stdio.EXPECT().Write(gomock.Any()).DoAndReturn(bytesWritten.Write).AnyTimes()

	cli := write.DryRunClient{
		CLI:        clients.CLI{ActiveConfig: config.Config{Org: "my-default-org"}, StdIO: stdio},
		LineReader: &mockReader,
	}

	require.NoError(t, cli.WriteDryRun(context.Background()))
	require.Equal(t, inLines, bytesWritten.String())
}
