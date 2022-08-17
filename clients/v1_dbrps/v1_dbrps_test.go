package v1dbrps_test

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
	v1dbrps "github.com/influxdata/influx-cli/v2/clients/v1_dbrps"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/internal/testutils"
	"github.com/stretchr/testify/require"
)

var (
	id1        = "1111111111111111"
	errApiTest = errors.New("api error for testing")
)

func TestClient_List(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               v1dbrps.ListParams
		registerExpectations func(*testing.T, *mock.MockDBRPsApi)
		expectedError        error
		outLines             []string
	}{
		{
			name:          "no org id or org name",
			params:        v1dbrps.ListParams{},
			expectedError: clients.ErrMustSpecifyOrg,
		},
		{
			name: "no results",
			params: v1dbrps.ListParams{
				OrgParams: clients.OrgParams{
					OrgID: id1,
				},
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().GetDBRPs(gomock.Any()).Return(api.ApiGetDBRPsRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().GetDBRPsExecute(gomock.Any()).Return(api.DBRPs{}, nil)
			},
		},
		{
			name: "many results",
			params: v1dbrps.ListParams{
				OrgParams: clients.OrgParams{
					OrgName: "example-org",
				},
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().GetDBRPs(gomock.Any()).Return(api.ApiGetDBRPsRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().GetDBRPsExecute(gomock.Any()).Return(api.DBRPs{
					Content: &[]api.DBRP{
						{
							Id:              "123",
							Database:        "someDB",
							BucketID:        "456",
							RetentionPolicy: "someRP",
							Default:         false,
							OrgID:           "1234123412341234",
						},
						{
							Id:              "234",
							Database:        "someDB",
							BucketID:        "456",
							RetentionPolicy: "someRP",
							Default:         true,
							OrgID:           "1234123412341234",
						},
					},
				}, nil)
			},
			outLines: []string{
				`123\s+someDB\s+456\s+someRP\s+false\s+1234123412341234`,
				`234\s+someDB\s+456\s+someRP\s+true\s+1234123412341234`,
			},
		},
		{
			name: "api error",
			params: v1dbrps.ListParams{
				OrgParams: clients.OrgParams{
					OrgName: "someOrg",
				},
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().GetDBRPs(gomock.Any()).Return(api.ApiGetDBRPsRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().GetDBRPsExecute(gomock.Any()).Return(api.DBRPs{}, errApiTest)
			},
			expectedError: fmt.Errorf("failed to list dbrps: %w", errApiTest),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			DBRPsApi := mock.NewMockDBRPsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, DBRPsApi)
			}

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := v1dbrps.Client{CLI: clients.CLI{StdIO: stdio}, DBRPsApi: DBRPsApi}

			err := cli.List(context.Background(), &tc.params)
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				header := `ID\s+Database\s+Bucket\s+ID\s+Retention Policy\s+Default\s+Organization ID`
				testutils.MatchLines(t,
					append([]string{header},
						append(tc.outLines,
							[]string{
								`VIRTUAL DBRP MAPPINGS \(READ-ONLY\)`,
								"----------------------------------",
								header}...)...),
					strings.Split(stdout.String(), "\n"))
			}
		})
	}
}

func TestClient_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               v1dbrps.CreateParams
		registerExpectations func(*testing.T, *mock.MockDBRPsApi)
		expectedError        error
		outLines             []string
	}{
		{
			name:          "no org id or org name",
			params:        v1dbrps.CreateParams{},
			expectedError: clients.ErrMustSpecifyOrg,
		},
		{
			name: "create with org id",
			params: v1dbrps.CreateParams{
				OrgParams: clients.OrgParams{
					OrgID: id1,
				},
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().PostDBRP(gomock.Any()).Return(api.ApiPostDBRPRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().PostDBRPExecute(gomock.Any()).Return(api.DBRP{
					Id:              "123",
					Database:        "someDB",
					BucketID:        "456",
					RetentionPolicy: "someRP",
					Default:         false,
					OrgID:           "1234123412341234",
				}, nil)
			},
			outLines: []string{
				`123\s+someDB\s+456\s+someRP\s+false\s+1234123412341234`,
			},
		},
		{
			name: "api error",
			params: v1dbrps.CreateParams{
				OrgParams: clients.OrgParams{
					OrgID: id1,
				},
				BucketID: "1234",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().PostDBRP(gomock.Any()).Return(api.ApiPostDBRPRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().PostDBRPExecute(gomock.Any()).Return(api.DBRP{}, errApiTest)
			},
			expectedError: fmt.Errorf("failed to create dbrp for bucket %q: %w", "1234", errApiTest),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			DBRPsApi := mock.NewMockDBRPsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, DBRPsApi)
			}

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := v1dbrps.Client{CLI: clients.CLI{StdIO: stdio}, DBRPsApi: DBRPsApi}

			err := cli.Create(context.Background(), &tc.params)
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				testutils.MatchLines(t, append([]string{`ID\s+Database\s+Bucket\s+ID\s+Retention Policy\s+Default\s+Organization ID`}, tc.outLines...), strings.Split(stdout.String(), "\n"))
			}
		})
	}
}

func TestClient_Update(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               v1dbrps.UpdateParams
		registerExpectations func(*testing.T, *mock.MockDBRPsApi)
		expectedError        error
		outLines             []string
	}{
		{
			name:          "no org id or org name",
			params:        v1dbrps.UpdateParams{},
			expectedError: clients.ErrMustSpecifyOrg,
		},
		{
			name: "update with org id",
			params: v1dbrps.UpdateParams{
				OrgParams: clients.OrgParams{
					OrgID: id1,
				},
				ID: "123",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().PatchDBRPID(gomock.Any(), "123").Return(api.ApiPatchDBRPIDRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().PatchDBRPIDExecute(gomock.Any()).Return(api.DBRPGet{
					Content: &api.DBRP{
						Id:              "123",
						Database:        "someDB",
						BucketID:        "456",
						RetentionPolicy: "someRP",
						Default:         false,
						OrgID:           "1234123412341234",
					},
				}, nil)
			},
			outLines: []string{
				`123\s+someDB\s+456\s+someRP\s+false\s+1234123412341234`,
			},
		},
		{
			name: "update with org name",
			params: v1dbrps.UpdateParams{
				OrgParams: clients.OrgParams{
					OrgName: "someOrg",
				},
				ID: "123",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().PatchDBRPID(gomock.Any(), "123").Return(api.ApiPatchDBRPIDRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().PatchDBRPIDExecute(gomock.Any()).Return(api.DBRPGet{
					Content: &api.DBRP{
						Id:              "123",
						Database:        "someDB",
						BucketID:        "456",
						RetentionPolicy: "someRP",
						Default:         false,
						OrgID:           "1234123412341234",
					},
				}, nil)
			},
			outLines: []string{
				`123\s+someDB\s+456\s+someRP\s+false\s+1234123412341234`,
			},
		},
		{
			name: "api error",
			params: v1dbrps.UpdateParams{
				OrgParams: clients.OrgParams{
					OrgName: "someOrg",
				},
				ID: "123",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().PatchDBRPID(gomock.Any(), "123").Return(api.ApiPatchDBRPIDRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().PatchDBRPIDExecute(gomock.Any()).Return(api.DBRPGet{}, errApiTest)
			},
			expectedError: fmt.Errorf("failed to update DBRP mapping %q: %w", "123", errApiTest),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			DBRPsApi := mock.NewMockDBRPsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, DBRPsApi)
			}

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := v1dbrps.Client{CLI: clients.CLI{StdIO: stdio}, DBRPsApi: DBRPsApi}

			err := cli.Update(context.Background(), &tc.params)
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				testutils.MatchLines(t, append([]string{`ID\s+Database\s+Bucket\s+ID\s+Retention Policy\s+Default\s+Organization ID`}, tc.outLines...), strings.Split(stdout.String(), "\n"))
			}
		})
	}
}

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		params               v1dbrps.DeleteParams
		registerExpectations func(*testing.T, *mock.MockDBRPsApi)
		expectedError        error
		outLines             []string
	}{
		{
			name:          "no org id or org name",
			params:        v1dbrps.DeleteParams{},
			expectedError: clients.ErrMustSpecifyOrg,
		},
		{
			name: "delete with org id",
			params: v1dbrps.DeleteParams{
				OrgParams: clients.OrgParams{
					OrgID: id1,
				},
				ID: "123",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().GetDBRPsID(gomock.Any(), "123").Return(api.ApiGetDBRPsIDRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().DeleteDBRPID(gomock.Any(), "123").Return(api.ApiDeleteDBRPIDRequest{ApiService: DBRPsApi})

				DBRPsApi.EXPECT().GetDBRPsIDExecute(gomock.Any()).Return(api.DBRPGet{
					Content: &api.DBRP{
						Id:              "123",
						Database:        "someDB",
						BucketID:        "456",
						RetentionPolicy: "someRP",
						Default:         false,
						OrgID:           "1234123412341234",
					},
				}, nil)

				DBRPsApi.EXPECT().DeleteDBRPIDExecute(gomock.Any()).Return(nil)
			},
			outLines: []string{
				`123\s+someDB\s+456\s+someRP\s+false\s+1234123412341234`,
			},
		},
		{
			name: "delete with org name",
			params: v1dbrps.DeleteParams{
				OrgParams: clients.OrgParams{
					OrgName: "someOrg",
				},
				ID: "123",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().GetDBRPsID(gomock.Any(), "123").Return(api.ApiGetDBRPsIDRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().DeleteDBRPID(gomock.Any(), "123").Return(api.ApiDeleteDBRPIDRequest{ApiService: DBRPsApi})

				DBRPsApi.EXPECT().GetDBRPsIDExecute(gomock.Any()).Return(api.DBRPGet{
					Content: &api.DBRP{
						Id:              "123",
						Database:        "someDB",
						BucketID:        "456",
						RetentionPolicy: "someRP",
						Default:         false,
						OrgID:           "1234123412341234",
					},
				}, nil)

				DBRPsApi.EXPECT().DeleteDBRPIDExecute(gomock.Any()).Return(nil)
			},
			outLines: []string{
				`123\s+someDB\s+456\s+someRP\s+false\s+1234123412341234`,
			},
		},
		{
			name: "api error with get request",
			params: v1dbrps.DeleteParams{
				OrgParams: clients.OrgParams{
					OrgName: "someOrg",
				},
				ID: "123",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().GetDBRPsID(gomock.Any(), "123").Return(api.ApiGetDBRPsIDRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().DeleteDBRPID(gomock.Any(), "123").Return(api.ApiDeleteDBRPIDRequest{ApiService: DBRPsApi})

				DBRPsApi.EXPECT().GetDBRPsIDExecute(gomock.Any()).Return(api.DBRPGet{}, errApiTest)
			},
			expectedError: fmt.Errorf("failed to find DBRP mapping %q: %w", "123", errApiTest),
		},
		{
			name: "api error with delete request",
			params: v1dbrps.DeleteParams{
				OrgParams: clients.OrgParams{
					OrgName: "someOrg",
				},
				ID: "123",
			},
			registerExpectations: func(t *testing.T, DBRPsApi *mock.MockDBRPsApi) {
				DBRPsApi.EXPECT().GetDBRPsID(gomock.Any(), "123").Return(api.ApiGetDBRPsIDRequest{ApiService: DBRPsApi})
				DBRPsApi.EXPECT().DeleteDBRPID(gomock.Any(), "123").Return(api.ApiDeleteDBRPIDRequest{ApiService: DBRPsApi})

				DBRPsApi.EXPECT().GetDBRPsIDExecute(gomock.Any()).Return(api.DBRPGet{
					Content: &api.DBRP{
						Id:              "123",
						Database:        "someDB",
						BucketID:        "456",
						RetentionPolicy: "someRP",
						Default:         false,
						OrgID:           "1234123412341234",
					},
				}, nil)

				DBRPsApi.EXPECT().DeleteDBRPIDExecute(gomock.Any()).Return(errApiTest)
			},
			expectedError: fmt.Errorf("failed to delete DBRP mapping %q: %w", "123", errApiTest),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			DBRPsApi := mock.NewMockDBRPsApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, DBRPsApi)
			}

			stdout := bytes.Buffer{}
			stdio := mock.NewMockStdIO(ctrl)
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(stdout.Write).AnyTimes()

			cli := v1dbrps.Client{CLI: clients.CLI{StdIO: stdio}, DBRPsApi: DBRPsApi}

			err := cli.Delete(context.Background(), &tc.params)
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				testutils.MatchLines(t, append([]string{`ID\s+Database\s+Bucket\s+ID\s+Retention Policy\s+Default\s+Organization ID`}, tc.outLines...), strings.Split(stdout.String(), "\n"))
			}
		})
	}
}
