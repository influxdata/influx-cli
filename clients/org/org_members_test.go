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
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var id1, _ = influxid.IDFromString("1111111111111111")
var id2, _ = influxid.IDFromString("2222222222222222")

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
				OrgID:    id1,
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
			},
			expectedOut: "user \"2222222222222222\" has been added as a member of org \"1111111111111111\"",
		},
		{
			name: "org by name",
			params: org.AddMemberParams{
				OrgName:  "org",
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}},
				}, nil)

				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
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
					Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}},
				}, nil)

				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.NotNil(t, body) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
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
		name                     string
		params                   org.ListMemberParams
		defaultOrgName           string
		registerOrgExpectations  func(*testing.T, *mock.MockOrganizationsApi)
		registerUserExpectations func(*testing.T, *mock.MockUsersApi)
		expectedOut              []string
		expectedErr              string
	}{
		{
			name:           "no members",
			defaultOrgName: "my-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}},
				}, nil)

				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String())
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).Return(req)
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).Return(api.ResourceMembers{}, nil)
			},
		},
		{
			name: "one member",
			params: org.ListMemberParams{
				OrgName: "org",
			},
			defaultOrgName: "my-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}},
				}, nil)

				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String())
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).Return(req)
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).
					Return(api.ResourceMembers{Users: &[]api.ResourceMember{{Id: api.PtrString(id2.String())}}}, nil)
			},
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				req := api.ApiGetUsersIDRequest{ApiService: userApi}.UserID(id2.String())
				userApi.EXPECT().GetUsersID(gomock.Any(), gomock.Eq(id2.String())).Return(req)
				userApi.EXPECT().GetUsersIDExecute(gomock.Eq(req)).Return(api.UserResponse{
					Id:     api.PtrString(id2.String()),
					Name:   "user1",
					Status: api.PtrString("active"),
				}, nil)
			},
			expectedOut: []string{`2222222222222222\s+user1\s+member\s+active`},
		},
		{
			name: "many members",
			params: org.ListMemberParams{
				OrgID: id1,
			},
			// NOTE: We previously saw a deadlock when # members was > 10, so test that here.
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String())
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).Return(req)
				members := make([]api.ResourceMember, 11)
				for i := 0; i < 11; i++ {
					members[i] = api.ResourceMember{Id: api.PtrString(fmt.Sprintf("%016d", i))}
				}
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).Return(api.ResourceMembers{Users: &members}, nil)
			},
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				for i := 0; i < 11; i++ {
					id := fmt.Sprintf("%016d", i)
					status := "active"
					if i%2 == 0 {
						status = "inactive"
					}
					req := api.ApiGetUsersIDRequest{ApiService: userApi}.UserID(id)
					userApi.EXPECT().GetUsersID(gomock.Any(), gomock.Eq(id)).Return(req)
					userApi.EXPECT().GetUsersIDExecute(gomock.Eq(req)).Return(api.UserResponse{
						Id:     &id,
						Name:   fmt.Sprintf("user%d", i),
						Status: &status,
					}, nil)
				}
			},
			expectedOut: []string{
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
			name: "no such user",
			params: org.ListMemberParams{
				OrgID: id1,
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				req := api.ApiGetOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String())
				orgApi.EXPECT().GetOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).Return(req)
				orgApi.EXPECT().GetOrgsIDMembersExecute(gomock.Eq(req)).
					Return(api.ResourceMembers{Users: &[]api.ResourceMember{{Id: api.PtrString(id2.String())}}}, nil)
			},
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				req := api.ApiGetUsersIDRequest{ApiService: userApi}.UserID(id2.String())
				userApi.EXPECT().GetUsersID(gomock.Any(), gomock.Eq(id2.String())).Return(req)
				userApi.EXPECT().GetUsersIDExecute(gomock.Eq(req)).Return(api.UserResponse{}, errors.New("not found"))
			},
			expectedErr: "user \"2222222222222222\": not found",
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
			if tc.registerUserExpectations != nil {
				tc.registerUserExpectations(t, userApi)
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
				OrgID:    id1,
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				req := api.ApiDeleteOrgsIDMembersIDRequest{ApiService: orgApi}.OrgID(id1.String()).UserID(id2.String())
				orgApi.EXPECT().
					DeleteOrgsIDMembersID(gomock.Any(), gomock.Eq(id1.String()), gomock.Eq(id2.String())).Return(req)
				orgApi.EXPECT().DeleteOrgsIDMembersIDExecute(gomock.Eq(req)).Return(nil)
			},
			expectedOut: "user \"2222222222222222\" has been removed from org \"1111111111111111\"",
		},
		{
			name: "org by name",
			params: org.RemoveMemberParams{
				OrgName:  "org",
				MemberId: id2,
			},
			defaultOrgName: "my-org",
			registerExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "org", *in.GetOrg())
				})).Return(api.Organizations{
					Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}},
				}, nil)

				req := api.ApiDeleteOrgsIDMembersIDRequest{ApiService: orgApi}.OrgID(id1.String()).UserID(id2.String())
				orgApi.EXPECT().
					DeleteOrgsIDMembersID(gomock.Any(), gomock.Eq(id1.String()), gomock.Eq(id2.String())).Return(req)
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
					Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}},
				}, nil)

				req := api.ApiDeleteOrgsIDMembersIDRequest{ApiService: orgApi}.OrgID(id1.String()).UserID(id2.String())
				orgApi.EXPECT().
					DeleteOrgsIDMembersID(gomock.Any(), gomock.Eq(id1.String()), gomock.Eq(id2.String())).Return(req)
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
