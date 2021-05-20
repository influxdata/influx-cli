package user_test

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/cmd"
	"github.com/influxdata/influx-cli/v2/internal/cmd/user"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var id1, _ = influxid.IDFromString("1111111111111111")
var id2, _ = influxid.IDFromString("2222222222222222")

func TestClient_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                     string
		params                   user.CreateParams
		defaultOrgName           string
		registerUserExpectations func(*testing.T, *mock.MockUsersApi)
		registerOrgExpectations  func(*testing.T, *mock.MockOrganizationsApi)

		expectedOut    string
		expectedStderr string
		expectedErr    string
	}{
		{
			name: "in org by ID",
			params: user.CreateParams{
				Name:     "my-user",
				Password: "my-password",
				OrgParams: cmd.OrgParams{
					OrgID: id1,
				},
			},
			defaultOrgName: "my-default-org",
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				userApi.EXPECT().PostUsers(gomock.Any()).Return(api.ApiPostUsersRequest{ApiService: userApi})
				userApi.EXPECT().PostUsersExecute(tmock.MatchedBy(func(in api.ApiPostUsersRequest) bool {
					body := in.GetUser()
					return assert.NotNil(t, body) && assert.Equal(t, "my-user", body.GetName())
				})).Return(api.UserResponse{Id: api.PtrString(id2.String()), Name: "my-user"}, nil)

				userApi.EXPECT().PostUsersIDPassword(gomock.Any(), gomock.Eq(id2.String())).
					Return(api.ApiPostUsersIDPasswordRequest{ApiService: userApi}.UserID(id2.String()))
				userApi.EXPECT().
					PostUsersIDPasswordExecute(tmock.MatchedBy(func(in api.ApiPostUsersIDPasswordRequest) bool {
						body := in.GetPasswordResetBody()
						return assert.NotNil(t, body) &&
							assert.Equal(t, id2.String(), in.GetUserID()) &&
							assert.Equal(t, "my-password", body.GetPassword())
					})).Return(nil)
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
			},
			expectedOut: `2222222222222222\s+my-user`,
		},
		{
			name: "in org by name",
			params: user.CreateParams{
				Name:     "my-user",
				Password: "my-password",
				OrgParams: cmd.OrgParams{
					OrgName: "my-org",
				},
			},
			defaultOrgName: "my-default-org",
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				userApi.EXPECT().PostUsers(gomock.Any()).Return(api.ApiPostUsersRequest{ApiService: userApi})
				userApi.EXPECT().PostUsersExecute(tmock.MatchedBy(func(in api.ApiPostUsersRequest) bool {
					body := in.GetUser()
					return assert.NotNil(t, body) && assert.Equal(t, "my-user", body.GetName())
				})).Return(api.UserResponse{Id: api.PtrString(id2.String()), Name: "my-user"}, nil)

				userApi.EXPECT().PostUsersIDPassword(gomock.Any(), gomock.Eq(id2.String())).
					Return(api.ApiPostUsersIDPasswordRequest{ApiService: userApi}.UserID(id2.String()))
				userApi.EXPECT().
					PostUsersIDPasswordExecute(tmock.MatchedBy(func(in api.ApiPostUsersIDPasswordRequest) bool {
						body := in.GetPasswordResetBody()
						return assert.NotNil(t, body) &&
							assert.Equal(t, id2.String(), in.GetUserID()) &&
							assert.Equal(t, "my-password", body.GetPassword())
					})).Return(nil)
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-org", *in.GetOrg())
				})).Return(api.Organizations{Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}}}, nil)

				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
			},
			expectedOut: `2222222222222222\s+my-user`,
		},
		{
			name: "in default org",
			params: user.CreateParams{
				Name:     "my-user",
				Password: "my-password",
			},
			defaultOrgName: "my-default-org",
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				userApi.EXPECT().PostUsers(gomock.Any()).Return(api.ApiPostUsersRequest{ApiService: userApi})
				userApi.EXPECT().PostUsersExecute(tmock.MatchedBy(func(in api.ApiPostUsersRequest) bool {
					body := in.GetUser()
					return assert.NotNil(t, body) && assert.Equal(t, "my-user", body.GetName())
				})).Return(api.UserResponse{Id: api.PtrString(id2.String()), Name: "my-user"}, nil)

				userApi.EXPECT().PostUsersIDPassword(gomock.Any(), gomock.Eq(id2.String())).
					Return(api.ApiPostUsersIDPasswordRequest{ApiService: userApi}.UserID(id2.String()))
				userApi.EXPECT().
					PostUsersIDPasswordExecute(tmock.MatchedBy(func(in api.ApiPostUsersIDPasswordRequest) bool {
						body := in.GetPasswordResetBody()
						return assert.NotNil(t, body) &&
							assert.Equal(t, id2.String(), in.GetUserID()) &&
							assert.Equal(t, "my-password", body.GetPassword())
					})).Return(nil)
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-default-org", *in.GetOrg())
				})).Return(api.Organizations{Orgs: &[]api.Organization{{Id: api.PtrString(id1.String())}}}, nil)

				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
			},
			expectedOut: `2222222222222222\s+my-user`,
		},
		{
			name: "no password",
			params: user.CreateParams{
				Name: "my-user",
				OrgParams: cmd.OrgParams{
					OrgID: id1,
				},
			},
			defaultOrgName: "my-default-org",
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				userApi.EXPECT().PostUsers(gomock.Any()).Return(api.ApiPostUsersRequest{ApiService: userApi})
				userApi.EXPECT().PostUsersExecute(tmock.MatchedBy(func(in api.ApiPostUsersRequest) bool {
					body := in.GetUser()
					return assert.NotNil(t, body) && assert.Equal(t, "my-user", body.GetName())
				})).Return(api.UserResponse{Id: api.PtrString(id2.String()), Name: "my-user"}, nil)
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
			},
			expectedOut:    `2222222222222222\s+my-user`,
			expectedStderr: `initial password not set`,
		},
		{
			name: "no org",
			params: user.CreateParams{
				Name:     "my-user",
				Password: "my-password",
			},
			expectedErr: user.ErrMustSpecifyOrg.Error(),
		},
		{
			name: "org not found",
			params: user.CreateParams{
				Name:     "my-user",
				Password: "my-password",
			},
			defaultOrgName: "my-default-org",
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().GetOrgs(gomock.Any()).Return(api.ApiGetOrgsRequest{ApiService: orgApi})
				orgApi.EXPECT().GetOrgsExecute(tmock.MatchedBy(func(in api.ApiGetOrgsRequest) bool {
					return assert.Equal(t, "my-default-org", *in.GetOrg())
				})).Return(api.Organizations{Orgs: &[]api.Organization{}}, nil)
			},
			expectedErr: "no org found",
		},
		{
			name: "assigning membership failed",
			params: user.CreateParams{
				Name:     "my-user",
				Password: "my-password",
				OrgParams: cmd.OrgParams{
					OrgID: id1,
				},
			},
			defaultOrgName: "my-default-org",
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				userApi.EXPECT().PostUsers(gomock.Any()).Return(api.ApiPostUsersRequest{ApiService: userApi})
				userApi.EXPECT().PostUsersExecute(tmock.MatchedBy(func(in api.ApiPostUsersRequest) bool {
					body := in.GetUser()
					return assert.NotNil(t, body) && assert.Equal(t, "my-user", body.GetName())
				})).Return(api.UserResponse{Id: api.PtrString(id2.String()), Name: "my-user"}, nil)
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{}, errors.New("I broke"))
			},
			expectedOut:    `2222222222222222\s+my-user`,
			expectedErr:    "I broke",
			expectedStderr: "initial password not set",
		},
		{
			name: "setting password failed",
			params: user.CreateParams{
				Name:     "my-user",
				Password: "my-password",
				OrgParams: cmd.OrgParams{
					OrgID: id1,
				},
			},
			defaultOrgName: "my-default-org",
			registerUserExpectations: func(t *testing.T, userApi *mock.MockUsersApi) {
				userApi.EXPECT().PostUsers(gomock.Any()).Return(api.ApiPostUsersRequest{ApiService: userApi})
				userApi.EXPECT().PostUsersExecute(tmock.MatchedBy(func(in api.ApiPostUsersRequest) bool {
					body := in.GetUser()
					return assert.NotNil(t, body) && assert.Equal(t, "my-user", body.GetName())
				})).Return(api.UserResponse{Id: api.PtrString(id2.String()), Name: "my-user"}, nil)

				userApi.EXPECT().PostUsersIDPassword(gomock.Any(), gomock.Eq(id2.String())).
					Return(api.ApiPostUsersIDPasswordRequest{ApiService: userApi}.UserID(id2.String()))
				userApi.EXPECT().
					PostUsersIDPasswordExecute(tmock.MatchedBy(func(in api.ApiPostUsersIDPasswordRequest) bool {
						body := in.GetPasswordResetBody()
						return assert.NotNil(t, body) &&
							assert.Equal(t, id2.String(), in.GetUserID()) &&
							assert.Equal(t, "my-password", body.GetPassword())
					})).Return(errors.New("I broke"))
			},
			registerOrgExpectations: func(t *testing.T, orgApi *mock.MockOrganizationsApi) {
				orgApi.EXPECT().PostOrgsIDMembers(gomock.Any(), gomock.Eq(id1.String())).
					Return(api.ApiPostOrgsIDMembersRequest{ApiService: orgApi}.OrgID(id1.String()))
				orgApi.EXPECT().PostOrgsIDMembersExecute(tmock.MatchedBy(func(in api.ApiPostOrgsIDMembersRequest) bool {
					body := in.GetAddResourceMemberRequestBody()
					return assert.NotNil(t, body) &&
						assert.Equal(t, id1.String(), in.GetOrgID()) &&
						assert.Equal(t, id2.String(), body.GetId())
				})).Return(api.ResourceMember{Id: api.PtrString(id2.String())}, nil)
			},
			expectedOut: `2222222222222222\s+my-user`,
			expectedErr: "I broke",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			orgApi := mock.NewMockOrganizationsApi(ctrl)
			if tc.registerOrgExpectations != nil {
				tc.registerOrgExpectations(t, orgApi)
			}
			userApi := mock.NewMockUsersApi(ctrl)
			if tc.registerUserExpectations != nil {
				tc.registerUserExpectations(t, userApi)
			}

			stdout := bytes.Buffer{}
			stderr := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()
			stdio.EXPECT().WriteErr(gomock.Any()).DoAndReturn(stderr.Write).AnyTimes()

			cli := user.Client{
				CLI:              cmd.CLI{StdIO: stdio, ActiveConfig: config.Config{Org: tc.defaultOrgName}},
				OrganizationsApi: orgApi,
				UsersApi:         userApi,
			}
			err := cli.Create(context.Background(), &tc.params)
			require.Contains(t, stderr.String(), tc.expectedStderr)
			if tc.expectedOut != "" {
				testutils.MatchLines(t, []string{`ID\s+Name`, tc.expectedOut}, strings.Split(stdout.String(), "\n"))
			}

			if tc.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErr)
				return
			}
			require.NoError(t, err)
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
			userApi := mock.NewMockUsersApi(ctrl)

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := user.Client{CLI: cmd.CLI{StdIO: stdio}, UsersApi: userApi}

			getReq := api.ApiGetUsersIDRequest{ApiService: userApi}.UserID(id2.String())
			userApi.EXPECT().GetUsersID(gomock.Any(), gomock.Eq(id2.String())).Return(getReq)
			userApi.EXPECT().GetUsersIDExecute(gomock.Eq(getReq)).
				DoAndReturn(func(api.ApiGetUsersIDRequest) (api.UserResponse, error) {
					if tc.notFound {
						return api.UserResponse{}, &api.Error{Code: api.ERRORCODE_NOT_FOUND}
					}
					return api.UserResponse{Id: api.PtrString(id2.String()), Name: "my-user"}, nil
				})

			if tc.notFound {
				require.Error(t, cli.Delete(context.Background(), id2))
				require.Empty(t, stdout.String())
				return
			}

			delReq := api.ApiDeleteUsersIDRequest{ApiService: userApi}.UserID(id2.String())
			userApi.EXPECT().DeleteUsersID(gomock.Any(), gomock.Eq(id2.String())).Return(delReq)
			userApi.EXPECT().DeleteUsersIDExecute(delReq).Return(nil)

			err := cli.Delete(context.Background(), id2)
			require.NoError(t, err)
			testutils.MatchLines(t, []string{
				`ID\s+Name\s+Deleted`,
				`2222222222222222\s+my-user\s+true`,
			}, strings.Split(stdout.String(), "\n"))
		})
	}
}

func TestClient_List(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               user.ListParams
		registerExpectations func(*testing.T, *mock.MockUsersApi)
		outLines             []string
	}{
		{
			name: "no results",
			registerExpectations: func(t *testing.T, usersApi *mock.MockUsersApi) {
				usersApi.EXPECT().GetUsers(gomock.Any()).Return(api.ApiGetUsersRequest{ApiService: usersApi})
				usersApi.EXPECT().GetUsersExecute(gomock.Any()).Return(api.Users{}, nil)
			},
		},
		{
			name: "many results",
			registerExpectations: func(t *testing.T, usersApi *mock.MockUsersApi) {
				usersApi.EXPECT().GetUsers(gomock.Any()).Return(api.ApiGetUsersRequest{ApiService: usersApi})
				usersApi.EXPECT().GetUsersExecute(gomock.Any()).Return(api.Users{
					Users: &[]api.UserResponse{
						{Id: api.PtrString("123"), Name: "user1"},
						{Id: api.PtrString("456"), Name: "user2"},
					},
				}, nil)
			},
			outLines: []string{`123\s+user1`, `456\s+user2`},
		},
		{
			name: "by name",
			params: user.ListParams{Name: "user1"},
			registerExpectations: func(t *testing.T, usersApi *mock.MockUsersApi) {
				usersApi.EXPECT().GetUsers(gomock.Any()).Return(api.ApiGetUsersRequest{ApiService: usersApi})
				usersApi.EXPECT().GetUsersExecute(tmock.MatchedBy(func(in api.ApiGetUsersRequest) bool {
					return assert.Equal(t, "user1", *in.GetName()) && assert.Nil(t, in.GetId())
				})).Return(api.Users{
					Users: &[]api.UserResponse{
						{Id: api.PtrString("123"), Name: "user1"},
					},
				}, nil)
			},
			outLines: []string{`123\s+user1`},
		},
		{
			name: "by ID",
			params: user.ListParams{Id: id2},
			registerExpectations: func(t *testing.T, usersApi *mock.MockUsersApi) {
				usersApi.EXPECT().GetUsers(gomock.Any()).Return(api.ApiGetUsersRequest{ApiService: usersApi})
				usersApi.EXPECT().GetUsersExecute(tmock.MatchedBy(func(in api.ApiGetUsersRequest) bool {
					return assert.Equal(t, id2.String(), *in.GetId()) && assert.Nil(t, in.GetName())
				})).Return(api.Users{
					Users: &[]api.UserResponse{
						{Id: api.PtrString(id2.String()), Name: "user11"},
					},
				}, nil)
			},
			outLines: []string{`2222222222222222\s+user11`},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			userApi := mock.NewMockUsersApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, userApi)
			}
			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := user.Client{CLI: cmd.CLI{StdIO: stdio}, UsersApi: userApi}
			require.NoError(t, cli.List(context.Background(), &tc.params))
			testutils.MatchLines(t, append([]string{`ID\s+Name`}, tc.outLines...), strings.Split(stdout.String(), "\n"))
		})
	}
}

func TestClient_Update(t *testing.T) {
	t.Parallel()
}

func TestClient_SetPassword(t *testing.T) {
	t.Parallel()
}
