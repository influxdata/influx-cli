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
	defaultOrgName = "default org"
	fakeResults = "data data data"
	fakeKey = "key1"
)

func TestSecret_List(t *testing.T) {
	t.Parallel()

	printHeader := "Key\t\tOrganization ID\n"
	id, err := influxid.IDFromString("1111111111111111")
	require.NoError(t, err)

	testCases := []struct {
		name string
		params secret.ListParams
		defaultOrgName string
		registerExpectations func(t *testing.T, secretApi *mock.MockSecretsApi, orgApi *mock.MockOrganizationsApi)
		expectMatcher string
		expectError string
	}{
		{
			name: "org id",
			params: secret.ListParams{
				OrgParams: clients.OrgParams{
					OrgID: id,
				},
			},
			defaultOrgName: defaultOrgName,
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi, orgApi *mock.MockOrganizationsApi) {
				req := api.ApiGetOrgsIDSecretsRequest{ApiService: secretApi}.OrgID(id.String())
				secretApi.EXPECT().GetOrgsIDSecrets(gomock.Any(), gomock.Eq(id.String())).Return(req)
				secretApi.EXPECT().GetOrgsIDSecretsExecute(gomock.Eq(req)).
					Return(api.SecretKeysResponse{Secrets: &[]string{fakeResults}}, nil)
			},
			expectMatcher: printHeader+fakeResults+"\t"+id.String()+"\n",
		},
		{
			name: "default org",
			params: secret.ListParams{
				OrgParams: clients.OrgParams{},
			},
			defaultOrgName: defaultOrgName,
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi, orgApi *mock.MockOrganizationsApi) {
				orgReq := api.ApiGetOrgsRequest{ApiService: orgApi}.Org(defaultOrgName)
				orgObj := api.Organization{Name: defaultOrgName}
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(orgReq)
				orgApi.EXPECT().GetOrgsExecute(orgReq).Return(api.Organizations{Orgs: &[]api.Organization{orgObj}}, nil)

				secReq := api.ApiGetOrgsIDSecretsRequest{ApiService: secretApi}.OrgID(orgObj.GetId())
				secObj := api.SecretKeysResponse{}
				secretApi.EXPECT().GetOrgsIDSecrets(gomock.Any(), orgObj.GetId()).Return(secReq)
				secretApi.EXPECT().GetOrgsIDSecretsExecute(secReq).Return(secObj, nil)
			},
			expectMatcher: "Key\tOrganization ID\n",
		},
		{
			name: "no org provided",
			params: secret.ListParams{
				OrgParams: clients.OrgParams{},
			},
			expectError: "must specify org ID or org name",
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
			organizationsApi := mock.NewMockOrganizationsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, secretsApi, organizationsApi)
			}
			cli := secret.Client{
				CLI: clients.CLI{
					ActiveConfig: config.Config{
						Org: tc.defaultOrgName,
					},
					StdIO: stdio,
				},
				SecretsApi: secretsApi,
				OrganizationsApi: organizationsApi,
			}

			err := cli.List(context.Background(), &tc.params)
			if tc.expectError != "" {
				require.Error(t, err)
				require.Equal(t, tc.expectError, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectMatcher, writtenBytes.String())
		})
	}
}

func TestSecret_Delete(t *testing.T) {
	t.Parallel()

	printHeader := "Key\tOrganization ID\t\tDeleted\n"
	id, err := influxid.IDFromString("1111111111111111")
	require.NoError(t, err)

	testCases := []struct {
		name string
		params secret.DeleteParams
		defaultOrgName string
		registerExpectations func(t *testing.T, secretApi *mock.MockSecretsApi)
		expectMatcher string
		expectError string
	}{
		{
			name: "delete",
			params: secret.DeleteParams{
				OrgParams: clients.OrgParams{
					OrgID: id,
				},
				Key:       fakeKey,
			},
			defaultOrgName: defaultOrgName,
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi) {
				req := api.ApiPostOrgsIDSecretsRequest{ApiService: secretApi}.
					OrgID(id.String()).
					SecretKeys(api.SecretKeys{Secrets: &[]string{"key1"}})
				secretApi.EXPECT().PostOrgsIDSecrets(gomock.Any(), gomock.Eq(id.String())).Return(req)
				secretApi.EXPECT().PostOrgsIDSecretsExecute(gomock.Eq(req)).Return(nil)
			},
			expectMatcher: printHeader+fakeKey+"\t"+id.String()+"\ttrue\n",
		},
		{
			// This situation cannot happen since the CLI will stop it.
			// Still worth testing though
			name: "delete no key",
			params: secret.DeleteParams{
				OrgParams: clients.OrgParams{
					OrgID: id,
				},
			},
			defaultOrgName: defaultOrgName,
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi) {
				req := api.ApiPostOrgsIDSecretsRequest{ApiService: secretApi}.
					OrgID(id.String()).
					SecretKeys(api.SecretKeys{Secrets: &[]string{""}})
				secretApi.EXPECT().PostOrgsIDSecrets(gomock.Any(), gomock.Eq(id.String())).Return(req)
				secretApi.EXPECT().PostOrgsIDSecretsExecute(gomock.Eq(req)).Return(nil)
			},
			expectMatcher: printHeader+"\t"+id.String()+"\ttrue\n",
		},
		{
			name: "delete no org",
			params: secret.DeleteParams{
				Key: fakeKey,
			},
			expectError: "must specify org ID or org name",
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

			err := cli.Delete(context.Background(), &tc.params)
			if tc.expectError != "" {
				require.Error(t, err)
				require.Equal(t, tc.expectError, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectMatcher, writtenBytes.String())
		})
	}
}

func TestSecret_Update(t *testing.T) {
	t.Parallel()

	printHeader := "Key\tOrganization ID\n"
	fakeValue := "someValue"
	id, err := influxid.IDFromString("1111111111111111")
	require.NoError(t, err)

	testCases := []struct {
		name string
		params secret.UpdateParams
		defaultOrgName string
		registerExpectations func(t *testing.T, secretApi *mock.MockSecretsApi, stdio *mock.MockStdIO)
		expectMatcher string
		expectError string
	}{
		{
			name: "update",
			params: secret.UpdateParams{
				OrgParams: clients.OrgParams{
					OrgID: id,
				},
				Key:       fakeKey,
				Value:     fakeValue,
			},
			defaultOrgName: defaultOrgName,
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi, stdio *mock.MockStdIO) {
				req := api.ApiPatchOrgsIDSecretsRequest{ApiService: secretApi}.
					OrgID(id.String()).
					RequestBody(map[string]string{fakeKey: fakeValue})
				secretApi.EXPECT().PatchOrgsIDSecrets(gomock.Any(), gomock.Eq(id.String())).Return(req)
				secretApi.EXPECT().PatchOrgsIDSecretsExecute(gomock.Eq(req)).Return(nil)
			},
			expectMatcher: printHeader+fakeKey+"\t"+id.String()+"\n",
		},
		{
			name: "update no key",
			params: secret.UpdateParams{
				OrgParams: clients.OrgParams{
					OrgID: id,
				},
				Value:     fakeValue,
			},
			defaultOrgName: defaultOrgName,
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi, stdio *mock.MockStdIO) {
				req := api.ApiPatchOrgsIDSecretsRequest{ApiService: secretApi}.
					OrgID(id.String()).
					RequestBody(map[string]string{"": fakeValue})
				secretApi.EXPECT().PatchOrgsIDSecrets(gomock.Any(), gomock.Eq(id.String())).Return(req)
				secretApi.EXPECT().PatchOrgsIDSecretsExecute(gomock.Eq(req)).Return(nil)
			},
			expectMatcher: printHeader+"\t"+id.String()+"\n",
		},
		{
			name: "update no value",
			params: secret.UpdateParams{
				OrgParams: clients.OrgParams{
					OrgID: id,
				},
				Key:       fakeKey,
			},
			defaultOrgName: defaultOrgName,
			registerExpectations: func(t *testing.T, secretApi *mock.MockSecretsApi, stdio *mock.MockStdIO) {
				stdio.EXPECT().GetSecret(gomock.Eq("Please type your secret"), gomock.Eq(0)).Return(fakeValue, nil)

				req := api.ApiPatchOrgsIDSecretsRequest{ApiService: secretApi}.
					OrgID(id.String()).
					RequestBody(map[string]string{fakeKey: fakeValue})
				secretApi.EXPECT().PatchOrgsIDSecrets(gomock.Any(), gomock.Eq(id.String())).Return(req)
				secretApi.EXPECT().PatchOrgsIDSecretsExecute(gomock.Eq(req)).Return(nil)
			},
			expectMatcher: printHeader+fakeKey+"\t"+id.String()+"\n",
		},
		{
			name: "update no org",
			params: secret.UpdateParams{
				Key:   fakeKey,
				Value: fakeValue,
			},
			expectError: "must specify org ID or org name",
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
				tc.registerExpectations(t, secretsApi, stdio)
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

			err := cli.Update(context.Background(), &tc.params)
			if tc.expectError != "" {
				require.Error(t, err)
				require.Equal(t, tc.expectError, err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectMatcher, writtenBytes.String())
		})
	}
}
