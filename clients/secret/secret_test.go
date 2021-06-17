package secret_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/secret"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/stretchr/testify/require"
)

const (
	printHeader = "Key\t\tOrganization ID\n"
	fakeResults = "data data data"
)

func TestSecret_List(t *testing.T) {
	t.Parallel()

	id, err := influxid.IDFromString("1111111111111111")
	require.NoError(t, err)

	testCases := []struct {
		name string
		params secret.ListParams
		defaultOrgName string
		registerExpectations func(t *testing.T, secretApi *mock.MockSecretsApi)
		expectError string
	}{
		{
			name: "org id",
			params: secret.ListParams{
				OrgParams: clients.OrgParams{
					OrgID: id,
				},
			},
			defaultOrgName: "default-org",
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi) {
				req := api.ApiGetOrgsIDSecretsRequest{ApiService: secretApi}.OrgID(id.String())
				secretApi.EXPECT().GetOrgsIDSecrets(gomock.Any(), gomock.Eq(id.String())).Return(req)
				secretApi.EXPECT().GetOrgsIDSecretsExecute(gomock.Eq(req)).Return(api.SecretKeysResponse{Secrets: &[]string{fakeResults}}, nil)
			},
		},
		{
			name: "no org provided",
			params: secret.ListParams{
				OrgParams: clients.OrgParams{},
			},
			expectError: "org or org-id must be provided",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()

			secretsApi := mock.NewMockSecretsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, secretsApi)
			}
			cli := secret.Client{
				CLI: clients.CLI{
					ActiveConfig: config.Config{
						Org: tc.defaultOrgName,
					},
					StdIO: stdio,
				},
				SecretsApi: secretsApi,
			}

			err := cli.List(context.Background(), &tc.params)
			if tc.expectError != "" {
				require.Error(t, err)
				require.Equal(t, tc.expectError, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t,
				printHeader+fakeResults+"\t"+id.String()+"\n",
				writtenBytes.String())
		})
	}
}
