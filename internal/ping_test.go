package internal_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/internal"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/require"
)

func Test_PingSuccess(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := mock.NewMockHealthApi(ctrl)
	client.EXPECT().GetHealth(gomock.Any()).Return(api.ApiGetHealthRequest{ApiService: client})
	client.EXPECT().GetHealthExecute(gomock.Any()).Return(api.HealthCheck{Status: api.HEALTHCHECKSTATUS_PASS}, nil)

	stdio := mock.NewMockStdIO(ctrl)
	bytesWritten := bytes.Buffer{}
	stdio.EXPECT().Write(gomock.Any()).DoAndReturn(bytesWritten.Write).AnyTimes()
	cli := &internal.CLI{StdIO: stdio}

	require.NoError(t, cli.Ping(context.Background(), client))
	require.Equal(t, "OK\n", bytesWritten.String())
}

func Test_PingFailedRequest(t *testing.T) {
	t.Parallel()

	e := "the internet is down"
	ctrl := gomock.NewController(t)
	client := mock.NewMockHealthApi(ctrl)
	client.EXPECT().GetHealth(gomock.Any()).Return(api.ApiGetHealthRequest{ApiService: client})
	client.EXPECT().GetHealthExecute(gomock.Any()).Return(api.HealthCheck{}, errors.New(e))

	cli := &internal.CLI{}
	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_PingFailedStatus(t *testing.T) {
	t.Parallel()

	e := "I broke"
	ctrl := gomock.NewController(t)
	client := mock.NewMockHealthApi(ctrl)
	client.EXPECT().GetHealth(gomock.Any()).Return(api.ApiGetHealthRequest{ApiService: client})
	client.EXPECT().GetHealthExecute(gomock.Any()).
		Return(api.HealthCheck{}, &api.HealthCheck{Status: api.HEALTHCHECKSTATUS_FAIL, Message: &e})

	cli := &internal.CLI{}
	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), e)
}

func Test_PingFailedStatusNoMessage(t *testing.T) {
	t.Parallel()

	name := "foo"
	ctrl := gomock.NewController(t)
	client := mock.NewMockHealthApi(ctrl)
	client.EXPECT().GetHealth(gomock.Any()).Return(api.ApiGetHealthRequest{ApiService: client})
	client.EXPECT().GetHealthExecute(gomock.Any()).
		Return(api.HealthCheck{}, &api.HealthCheck{Status: api.HEALTHCHECKSTATUS_FAIL, Name: name})

	cli := &internal.CLI{}
	err := cli.Ping(context.Background(), client)
	require.Error(t, err)
	require.Contains(t, err.Error(), name)
}
