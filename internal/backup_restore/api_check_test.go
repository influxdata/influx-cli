package backup_restore_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/require"
)

func TestServerIsLegacy(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		versionStr *string
		legacy     bool
		wantErr    string
	}{
		{
			name:       "2.0.x",
			versionStr: api.PtrString("2.0.7"),
			legacy:     true,
		},
		{
			name:       "2.1.x",
			versionStr: api.PtrString("2.1.0-RC1"),
		},
		{
			name:       "nightly",
			versionStr: api.PtrString("nightly-2020-01-01"),
		},
		{
			name:       "dev",
			versionStr: api.PtrString("some.custom-version.2"),
		},
		{
			name:       "1.x",
			versionStr: api.PtrString("1.9.3"),
			wantErr:    "InfluxDB v1 does not support the APIs",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			healthApi := mock.NewMockHealthApi(ctrl)
			healthApi.EXPECT().GetHealth(gomock.Any()).Return(api.ApiGetHealthRequest{ApiService: healthApi})
			healthApi.EXPECT().GetHealthExecute(gomock.Any()).Return(api.HealthCheck{Version: tc.versionStr}, nil)

			isLegacy, err := backup_restore.ServerIsLegacy(context.Background(), healthApi)

			if tc.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.wantErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.legacy, isLegacy)
		})
	}
}
