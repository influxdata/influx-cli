package org_test

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/org"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var id, _ = influxid.IDFromString("1111111111111111")

func TestClient_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               org.CreateParams
		registerExpectations func(*testing.T, *mock.MockOrganizationsApi)
		outLine              string
	}{
		{
			name: "name only",
			params: org.CreateParams{
				Name: "my-org",
			},
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgs(gomock.Any()).Return(api.ApiPostOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().PostOrgsExecute(tmock.MatchedBy(func(in api.ApiPostOrgsRequest) bool {
					body := in.GetPostOrganizationRequest()
					return assert.NotNil(t, body) &&
						assert.Equal(t, "my-org", body.GetName()) &&
						assert.Nil(t, body.Description)
				})).Return(api.Organization{Name: "my-org", Id: api.PtrString("123")}, nil)
			},
			outLine: `123\s+my-org`,
		},
		{
			name: "with description",
			params: org.CreateParams{
				Name:        "my-org",
				Description: "my cool new org",
			},
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgs(gomock.Any()).Return(api.ApiPostOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().PostOrgsExecute(tmock.MatchedBy(func(in api.ApiPostOrgsRequest) bool {
					body := in.GetPostOrganizationRequest()
					return assert.NotNil(t, body) &&
						assert.Equal(t, "my-org", body.GetName()) &&
						assert.Equal(t, "my cool new org", body.GetDescription())
				})).Return(api.Organization{
					Name:        "my-org",
					Id:          api.PtrString("123"),
					Description: api.PtrString("my cool new org"),
				}, nil)
			},
			outLine: `123\s+my-org`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			api := mock.NewMockOrganizationsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, api)
			}

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := org.Client{CLI: clients.CLI{StdIO: stdio}, OrganizationsApi: api}
			require.NoError(t, cli.Create(context.Background(), &tc.params))
			testutils.MatchLines(t, []string{`ID\s+Name`, tc.outLine}, strings.Split(stdout.String(), "\n"))
		})
	}
}

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		notFound bool
	}{
		{
			name: "delete existing",
		},
		{
			name:     "delete non-existing",
			notFound: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			orgApi := mock.NewMockOrganizationsApi(ctrl)

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := org.Client{CLI: clients.CLI{StdIO: stdio}, OrganizationsApi: orgApi}

			getReq := api.ApiGetOrgsIDRequest{ApiService: orgApi}.OrgID(id.String())
			orgApi.EXPECT().GetOrgsID(gomock.Any(), gomock.Eq(id.String())).Return(getReq)
			orgApi.EXPECT().GetOrgsIDExecute(gomock.Eq(getReq)).
				DoAndReturn(func(request api.ApiGetOrgsIDRequest) (api.Organization, error) {
					if tc.notFound {
						return api.Organization{}, &api.Error{Code: api.ERRORCODE_NOT_FOUND}
					}
					return api.Organization{Id: api.PtrString(id.String()), Name: "my-org"}, nil
				})

			if tc.notFound {
				require.Error(t, cli.Delete(context.Background(), id))
				require.Empty(t, stdout.String())
				return
			}

			delReq := api.ApiDeleteOrgsIDRequest{ApiService: orgApi}.OrgID(id.String())
			orgApi.EXPECT().DeleteOrgsID(gomock.Any(), gomock.Eq(id.String())).Return(delReq)
			orgApi.EXPECT().DeleteOrgsIDExecute(gomock.Eq(delReq)).Return(nil)

			require.NoError(t, cli.Delete(context.Background(), id))
			testutils.MatchLines(t, []string{
				`ID\s+Name\s+Deleted`,
				`1111111111111111\s+my-org\s+true`,
			}, strings.Split(stdout.String(), "\n"))
		})
	}
}

func TestClient_List(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               org.ListParams
		registerExpectations func(*testing.T, *mock.MockOrganizationsApi)
		outLines             []string
	}{
		{
			name: "no results",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(gomock.Any()).Return(api.Organizations{}, nil)
			},
		},
		{
			name: "many results",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(gomock.Any()).Return(api.Organizations{
					Orgs: &[]api.Organization{
						{Id: api.PtrString("123"), Name: "org1"},
						{Id: api.PtrString("456"), Name: "org2"},
					},
				}, nil)
			},
			outLines: []string{`123\s+org1`, `456\s+org2`},
		},
		{
			name:   "by name",
			params: org.ListParams{Name: "org1"},
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org1", *in.GetOrg()) && assert.Nil(t, in.GetOrgID())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{
						{Id: api.PtrString("123"), Name: "org1"},
					},
				}, nil)
			},
			outLines: []string{`123\s+org1`},
		},
		{
			name:   "by ID",
			params: org.ListParams{ID: id},
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Nil(t, in.GetOrg()) && assert.Equal(t, id.String(), *in.GetOrgID())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{
						{Id: api.PtrString(id.String()), Name: "org3"},
					},
				}, nil)
			},
			outLines: []string{fmt.Sprintf(`%s\s+org3`, id)},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			orgApi := mock.NewMockOrganizationsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, orgApi)
			}
			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := org.Client{CLI: clients.CLI{StdIO: stdio}, OrganizationsApi: orgApi}
			require.NoError(t, cli.List(context.Background(), &tc.params))
			testutils.MatchLines(t, append([]string{`ID\s+Name`}, tc.outLines...), strings.Split(stdout.String(), "\n"))
		})
	}
}

func TestClient_Update(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               org.UpdateParams
		registerExpectations func(*testing.T, *mock.MockOrganizationsApi)
		outLine              string
	}{
		{
			name:   "name",
			params: org.UpdateParams{ID: id, Name: "my-org"},
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PatchOrgsID(gomock.Any(), gomock.Eq(id.String())).
					Return(api.ApiPatchOrgsIDRequest{ApiService: orgApi}.OrgID(id.String()))
				orgApi.EXPECT().PatchOrgsIDExecute(tmock.MatchedBy(func(in api.ApiPatchOrgsIDRequest) bool {
					body := in.GetPatchOrganizationRequest()
					return assert.Equal(t, id.String(), in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, "my-org", body.GetName()) &&
						assert.Nil(t, body.Description)
				})).Return(api.Organization{Id: api.PtrString(id.String()), Name: "my-org"}, nil)
			},
			outLine: fmt.Sprintf(`%s\s+my-org`, id.String()),
		},
		{
			name:   "description",
			params: org.UpdateParams{ID: id, Description: "my cool org"},
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PatchOrgsID(gomock.Any(), gomock.Eq(id.String())).
					Return(api.ApiPatchOrgsIDRequest{ApiService: orgApi}.OrgID(id.String()))
				orgApi.EXPECT().PatchOrgsIDExecute(tmock.MatchedBy(func(in api.ApiPatchOrgsIDRequest) bool {
					body := in.GetPatchOrganizationRequest()
					return assert.Equal(t, id.String(), in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Nil(t, body.Name) &&
						assert.Equal(t, "my cool org", body.GetDescription())
				})).Return(api.Organization{Id: api.PtrString(id.String()), Name: "my-org"}, nil)
			},
			outLine: fmt.Sprintf(`%s\s+my-org`, id.String()),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			orgApi := mock.NewMockOrganizationsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, orgApi)
			}
			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := org.Client{CLI: clients.CLI{StdIO: stdio}, OrganizationsApi: orgApi}
			require.NoError(t, cli.Update(context.Background(), &tc.params))
			testutils.MatchLines(t, []string{`ID\s+Name`, tc.outLine}, strings.Split(stdout.String(), "\n"))
		})
	}
}
