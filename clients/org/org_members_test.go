package org_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/org"
	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var id1 = "1111111111111111"
var id2 = "2222222222222222"

func TestClient_AddMember(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               org.AddMemberParams
		defaultOrgName       string
		registerExpectations func(*testing.T, *mock.MockOrganizationsApi)
		expectedOut          string
		expectedErr          string
	}{
		{
			name: "org by ID",
			params: org.AddMemberParams{
				OrgParams: clients.OrgParams{OrgID: id1},
				MemberId:  id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1)).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.Equal(t, id1, in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, id2, body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2)}, nil)
			},
			expectedOut: "user \"2222222222222222\" has been added as a member of org \"1111111111111111\"",
		},
		{
			name: "org by name",
			params: org.AddMemberParams{
				OrgParams: clients.OrgParams{OrgName: "org"},
				MemberId:  id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1)}},
				}, nil)

				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1)).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.Equal(t, id1, in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, id2, body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2)}, nil)
			},
			expectedOut: "user \"2222222222222222\" has been added as a member of org \"1111111111111111\"",
		},
		{
			name: "by config org",
			params: org.AddMemberParams{
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1)}},
				}, nil)

				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1)).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.Equal(t, id1, in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, id2, body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2)}, nil)
			},
			expectedOut: "user \"2222222222222222\" has been added as a member of org \"1111111111111111\"",
		},
		{
			name: "no such org",
			params: org.AddMemberParams{
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{}, errors.New("not found"))
			},
			expectedErr: "not found",
		},
		{
			name:        "missing org",
			expectedErr: clients.ErrMustSpecifyOrg.Error(),
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

			cli := org.Client{
				CLI:              clients.CLI{StdIO: stdio, ActiveConfig: config.Config{Org: tc.defaultOrgName}},
				OrganizationsApi: api,
			}
			err := cli.AddMember(context.Background(), &tc.params)
			if tc.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedOut, strings.TrimSpace(stdout.String()))
		})
	}
}

func TestClient_ListMembers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                    string
		params                  org.ListMemberParams
		defaultOrgName          string
		registerOrgExpectations func(*testing.T, *mock.MockOrganizationsApi)
		expectedOut             []string
		expectedErr             string
	}{
		{
			name:           "no members",
			defaultOrgName: "my-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1)}},
				}, nil)

				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1)).Return(req)
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).Return(api.ResourceMembers{}, nil)

				ownerReq := api.ApiGetOrgsIDOwnersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDOwners(gomock.Any(), gomock.Eq(id1)).Return(ownerReq)
				orgApi.EXPECT().GetOrgsIDOwnersExecute(gomock.Eq(ownerReq)).Return(api.ResourceOwners{}, nil)
			},
		},
		{
			name: "one member",
			params: org.ListMemberParams{
				OrgParams: clients.OrgParams{OrgName: "org"},
			},
			defaultOrgName: "my-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1)}},
				}, nil)

				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1)).Return(req)
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).
					Return(api.ResourceMembers{Users: &[]api.ResourceMember{{
						Id:     api.PtrString(id2),
						Name:   "user1",
						Role:   api.PtrString("member"),
						Status: api.PtrString("active"),
					}}}, nil)

				ownerReq := api.ApiGetOrgsIDOwnersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDOwners(gomock.Any(), gomock.Eq(id1)).Return(ownerReq)
				orgApi.EXPECT().GetOrgsIDOwnersExecute(gomock.Eq(ownerReq)).Return(api.ResourceOwners{}, nil)
			},
			expectedOut: []string{`2222222222222222\s+user1\s+member\s+active`},
		},
		{
			name: "one owner",
			params: org.ListMemberParams{
				OrgParams: clients.OrgParams{OrgName: "org"},
			},
			defaultOrgName: "my-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1)}},
				}, nil)

				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1)).Return(req)
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).
					Return(api.ResourceMembers{}, nil)

				ownerReq := api.ApiGetOrgsIDOwnersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDOwners(gomock.Any(), gomock.Eq(id1)).Return(ownerReq)
				orgApi.EXPECT().GetOrgsIDOwnersExecute(gomock.Eq(ownerReq)).Return(api.ResourceOwners{Users: &[]api.ResourceOwner{{
					Id:     api.PtrString(id2),
					Name:   "user1",
					Role:   api.PtrString("owner"),
					Status: api.PtrString("active"),
				}}}, nil)
			},
			expectedOut: []string{`2222222222222222\s+user1\s+owner\s+active`},
		},
		{
			name: "many users/members",
			params: org.ListMemberParams{
				OrgParams: clients.OrgParams{OrgID: id1},
			},
			// NOTE: We previously saw a deadlock when # members was > 10, so test that here.
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1)).Return(req)
				members := make([]api.ResourceMember, 11)
				for i := 0; i < 11; i++ {
					status := "active"
					if i%2 == 0 {
						status = "inactive"
					}
					members[i] = api.ResourceMember{
						Id:     api.PtrString(fmt.Sprintf("%016d", i)),
						Name:   fmt.Sprintf("user%d", i),
						Status: &status,
						Role:   api.PtrString("member"),
					}
				}
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).Return(api.ResourceMembers{Users: &members}, nil)

				owners := make([]api.ResourceOwner, 5)
				offset := 11
				for i := offset; i < 5+offset; i++ {
					status := "active"
					if i%2 == 0 {
						status = "inactive"
					}
					owners[i-offset] = api.ResourceOwner{
						Id:     api.PtrString(fmt.Sprintf("%016d", i)),
						Name:   fmt.Sprintf("user%d", i),
						Status: &status,
						Role:   api.PtrString("owner"),
					}
				}
				ownerReq := api.ApiGetOrgsIDOwnersRequest{ApiService: orgApi}.OrgID(id1)
				orgApi.EXPECT().GetOrgsIDOwners(gomock.Any(), gomock.Eq(id1)).Return(ownerReq)
				orgApi.EXPECT().GetOrgsIDOwnersExecute(gomock.Eq(ownerReq)).Return(api.ResourceOwners{Users: &owners}, nil)
			},
			expectedOut: []string{
				`0000000000000011\s+user11\s+owner\s+active`,
				`0000000000000012\s+user12\s+owner\s+inactive`,
				`0000000000000013\s+user13\s+owner\s+active`,
				`0000000000000014\s+user14\s+owner\s+inactive`,
				`0000000000000015\s+user15\s+owner\s+active`,

				`0000000000000000\s+user0\s+member\s+inactive`,
				`0000000000000001\s+user1\s+member\s+active`,
				`0000000000000002\s+user2\s+member\s+inactive`,
				`0000000000000003\s+user3\s+member\s+active`,
				`0000000000000004\s+user4\s+member\s+inactive`,
				`0000000000000005\s+user5\s+member\s+active`,
				`0000000000000006\s+user6\s+member\s+inactive`,
				`0000000000000007\s+user7\s+member\s+active`,
				`0000000000000008\s+user8\s+member\s+inactive`,
				`0000000000000009\s+user9\s+member\s+active`,
				`0000000000000010\s+user10\s+member\s+inactive`,
			},
		},
		{
			name:           "no such org",
			defaultOrgName: "my-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{}, errors.New("not found"))
			},
			expectedErr: "not found",
		},
		{
			name:        "missing org",
			expectedErr: clients.ErrMustSpecifyOrg.Error(),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			orgApi := mock.NewMockOrganizationsApi(ctrl)
			userApi := mock.NewMockUsersApi(ctrl)
			if tc.registerOrgExpectations != nil {
				tc.registerOrgExpectations(t, orgApi)
			}

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := org.Client{
				CLI:              clients.CLI{StdIO: stdio, ActiveConfig: config.Config{Org: tc.defaultOrgName}},
				OrganizationsApi: orgApi,
				UsersApi:         userApi,
			}
			err := cli.ListMembers(context.Background(), &tc.params)
			if tc.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
				return
			}
			require.NoError(t, err)
			testutils.MatchLines(t, append([]string{`ID\s+Name\s+User Type\s+Status`}, tc.expectedOut...),
				strings.Split(stdout.String(), "\n"))
		})
	}
}

func TestClient_RemoveMembers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               org.RemoveMemberParams
		defaultOrgName       string
		registerExpectations func(*testing.T, *mock.MockOrganizationsApi)
		expectedOut          string
		expectedErr          string
	}{
		{
			name: "org by ID",
			params: org.RemoveMemberParams{
				OrgParams: clients.OrgParams{OrgID: id1},
				MemberId:  id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				req := api.ApiDeleteOrgsIDMembersIDRequest{ApiService: orgApi}.OrgID(id1).UserID(id2)
				orgApi.EXPECT().
					DeleteOrgsIDMembersID(gomock.Any(), gomock.Eq(id2), gomock.Eq(id1)).Return(req)
				orgApi.EXPECT().DeleteOrgsIDMembersIDExecute(gomock.Eq(req)).Return(nil)
			},
			expectedOut: "user \"2222222222222222\" has been removed from org \"1111111111111111\"",
		},
		{
			name: "org by name",
			params: org.RemoveMemberParams{
				OrgParams: clients.OrgParams{OrgName: "org"},
				MemberId:  id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1)}},
				}, nil)

				req := api.ApiDeleteOrgsIDMembersIDRequest{ApiService: orgApi}.OrgID(id1).UserID(id2)
				orgApi.EXPECT().
					DeleteOrgsIDMembersID(gomock.Any(), gomock.Eq(id2), gomock.Eq(id1)).Return(req)
				orgApi.EXPECT().DeleteOrgsIDMembersIDExecute(gomock.Eq(req)).Return(nil)
			},
			expectedOut: "user \"2222222222222222\" has been removed from org \"1111111111111111\"",
		},
		{
			name: "by config org",
			params: org.RemoveMemberParams{
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1)}},
				}, nil)

				req := api.ApiDeleteOrgsIDMembersIDRequest{ApiService: orgApi}.OrgID(id1).UserID(id2)
				orgApi.EXPECT().
					DeleteOrgsIDMembersID(gomock.Any(), gomock.Eq(id2), gomock.Eq(id1)).Return(req)
				orgApi.EXPECT().DeleteOrgsIDMembersIDExecute(gomock.Eq(req)).Return(nil)
			},
			expectedOut: "user \"2222222222222222\" has been removed from org \"1111111111111111\"",
		},
		{
			name: "no such org",
			params: org.RemoveMemberParams{
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{}, errors.New("not found"))
			},
			expectedErr: "not found",
		},
		{
			name:        "missing org",
			expectedErr: clients.ErrMustSpecifyOrg.Error(),
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

			cli := org.Client{
				CLI:              clients.CLI{StdIO: stdio, ActiveConfig: config.Config{Org: tc.defaultOrgName}},
				OrganizationsApi: api,
			}
			err := cli.RemoveMember(context.Background(), &tc.params)
			if tc.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedOut, strings.TrimSpace(stdout.String()))
		})
	}
}
