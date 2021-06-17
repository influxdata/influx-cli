package secret_test

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/secret"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSecret_List(t *testing.T) {
	t.Parallel()

	id, err := influxid.IDFromString("1111111111111111")
	require.NoError(t, err)

	fakeResults := "data data data"

	testCases := []struct {
		name string
		params secret.ListParams
		defaultOrgName string
		registerExpectations func(t *testing.T, secretApi *mock.MockSecretsApi)
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
				secretApi.EXPECT().GetOrgsIDSecrets(gomock.Any(), gomock.Any()).Return(api.ApiGetOrgsIDSecretsRequest{ApiService: secretApi})
				secretApi.EXPECT().GetOrgsIDSecretsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsIDSecretsRequest) bool {
					body := in.GetOrgID()
					return assert.NotEmpty(t, body) &&
						assert.Equal(t, body, id.String())
				})).Return(api.SecretKeysResponse{Secrets: &[]string{fakeResults}}, nil)
			},
		},
		{
			name: "org id empty",
			params: secret.ListParams{
				OrgParams: clients.OrgParams{},
			},
			/*registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi) {
				secretApi.EXPECT().GetOrgsIDSecrets(gomock.Any(), gomock.Any()).Return(api.ApiGetOrgsIDSecretsRequest{ApiService: secretApi})
				secretApi.EXPECT().GetOrgsIDSecretsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsIDSecretsRequest) bool {
					body := in.GetOrgID()
					return assert.NotEmpty(t, body) &&
						assert.Equal(t, body, id1.String())
				})).Return(api.SecretKeysResponse{Secrets: &[]string{fakeResults}}, nil)
			},*/
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
			}

			err := cli.List(context.Background(), &tc.params)
			require.NoError(t, err)
			require.Equal(t, fakeResults, writtenBytes.String())
		})
	}
}