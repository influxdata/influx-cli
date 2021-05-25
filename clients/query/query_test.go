package query_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/clients/query"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/influxdata/influx-cli/v2/pkg/influxid"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRawResultPrinter_PrintQueryResults(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		rawTable string
	}{
		{
			name:     "empty",
			rawTable: "",
		},
		{
			name: "single table",
			rawTable: `#group,false,false,true,true,false,false,true,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string
#default,_result,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,bar
,,0,1921-05-08T15:29:32.475078Z,2021-05-08T15:29:32.475078Z,2021-05-04T18:29:52.764702Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:29:32.475078Z,2021-05-08T15:29:32.475078Z,2021-05-04T19:30:59.67555Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:29:32.475078Z,2021-05-08T15:29:32.475078Z,2021-05-04T19:31:01.876079Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:29:32.475078Z,2021-05-08T15:29:32.475078Z,2021-05-04T19:31:02.499461Z,12345,qux,foo,"""baz"""
`,
		},
		{
			name: "multi table",
			rawTable: `#group,false,false,true,true,false,false,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string
#default,_result,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,bar
,,0,1921-05-08T15:42:58.218436Z,2021-05-08T15:42:58.218436Z,2021-05-04T18:29:52.764702Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:42:58.218436Z,2021-05-08T15:42:58.218436Z,2021-05-04T19:30:59.67555Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:42:58.218436Z,2021-05-08T15:42:58.218436Z,2021-05-04T19:31:01.876079Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:42:58.218436Z,2021-05-08T15:42:58.218436Z,2021-05-04T19:31:02.499461Z,12345,qux,foo,"""baz"""

#group,false,false,true,true,false,false,true,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,bar,is_foo
,,1,1921-05-08T15:42:58.218436Z,2021-05-08T15:42:58.218436Z,2021-05-08T15:42:19.567667Z,12345,qux,foo,"""baz""",t

#group,false,false,true,false,false,false,false,false,false,false
#datatype,string,long,string,string,string,long,long,long,long,double
#default,_profiler,,,,,,,,,
,result,table,_measurement,Type,Label,Count,MinDuration,MaxDuration,DurationSum,MeanDuration
,,0,profiler/operator,*influxdb.readFilterSource,ReadRange2,1,367331,367331,367331,367331
`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			in := ioutil.NopCloser(strings.NewReader(tc.rawTable))
			out := bytes.Buffer{}
			require.NoError(t, query.RawResultPrinter.PrintQueryResults(in, &out))
			require.Equal(t, tc.rawTable, out.String())
		})
	}
}

func TestQuery(t *testing.T) {
	t.Parallel()

	// Use dummy data here + the raw output printer to keep things
	// focused on the business logic of the Query command; details
	// of how results are formatted are tested elsewhere.
	fakeQuery := query.BuildDefaultAST("I'm a query!")
	fakeResults := "data data data"

	orgID, err := influxid.IDFromString("1111111111111111")
	require.NoError(t, err)

	testCases := []struct {
		name                 string
		params               query.Params
		configOrgName        string
		registerExpectations func(t *testing.T, queryApi *mock.MockQueryApi)
		expectInErr          string
	}{
		{
			name: "by org ID",
			params: query.Params{
				OrgParams: clients.OrgParams{
					OrgID: orgID,
				},
				Query: fakeQuery.Query,
			},
			configOrgName: "default-org",
			registerExpectations: func(t *testing.T, queryApi *mock.MockQueryApi) {
				queryApi.EXPECT().PostQuery(gomock.Any()).Return(api.ApiPostQueryRequest{ApiService: queryApi})
				queryApi.EXPECT().PostQueryExecute(tmock.MatchedBy(func(in api.ApiPostQueryRequest) bool {
					body := in.GetQuery()
					return assert.NotNil(t, body) &&
						assert.Equal(t, fakeQuery, *body) &&
						assert.Equal(t, orgID.String(), *in.GetOrgID()) &&
						assert.Nil(t, in.GetOrg())
				})).Return(ioutil.NopCloser(strings.NewReader(fakeResults)), nil)
			},
		},
		{
			name: "by org name",
			params: query.Params{
				OrgParams: clients.OrgParams{
					OrgName: "my-org",
				},
				Query: fakeQuery.Query,
			},
			configOrgName: "default-org",
			registerExpectations: func(t *testing.T, queryApi *mock.MockQueryApi) {
				queryApi.EXPECT().PostQuery(gomock.Any()).Return(api.ApiPostQueryRequest{ApiService: queryApi})
				queryApi.EXPECT().PostQueryExecute(tmock.MatchedBy(func(in api.ApiPostQueryRequest) bool {
					body := in.GetQuery()
					return assert.NotNil(t, body) &&
						assert.Equal(t, fakeQuery, *body) &&
						assert.Equal(t, "my-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetOrgID())
				})).Return(ioutil.NopCloser(strings.NewReader(fakeResults)), nil)
			},
		},
		{
			name: "by org name from config",
			params: query.Params{
				OrgParams: clients.OrgParams{},
				Query:     fakeQuery.Query,
			},
			configOrgName: "default-org",
			registerExpectations: func(t *testing.T, queryApi *mock.MockQueryApi) {
				queryApi.EXPECT().PostQuery(gomock.Any()).Return(api.ApiPostQueryRequest{ApiService: queryApi})
				queryApi.EXPECT().PostQueryExecute(tmock.MatchedBy(func(in api.ApiPostQueryRequest) bool {
					body := in.GetQuery()
					return assert.NotNil(t, body) &&
						assert.Equal(t, fakeQuery, *body) &&
						assert.Equal(t, "default-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetOrgID())
				})).Return(ioutil.NopCloser(strings.NewReader(fakeResults)), nil)
			},
		},
		{
			name: "no org specified",
			params: query.Params{
				OrgParams: clients.OrgParams{},
				Query:     fakeQuery.Query,
			},
			expectInErr: clients.ErrMustSpecifyOrg.Error(),
		},
		{
			name: "with profilers",
			params: query.Params{
				OrgParams: clients.OrgParams{},
				Query:     fakeQuery.Query,
				Profilers: []string{"foo", "bar"},
			},
			configOrgName: "default-org",
			registerExpectations: func(t *testing.T, queryApi *mock.MockQueryApi) {
				queryApi.EXPECT().PostQuery(gomock.Any()).Return(api.ApiPostQueryRequest{ApiService: queryApi})

				expectedBody := fakeQuery
				expectedBody.Extern = query.BuildExternAST([]string{"foo", "bar"})

				queryApi.EXPECT().PostQueryExecute(tmock.MatchedBy(func(in api.ApiPostQueryRequest) bool {
					body := in.GetQuery()
					return assert.NotNil(t, body) &&
						assert.Equal(t, expectedBody, *body) &&
						assert.Equal(t, "default-org", *in.GetOrg()) &&
						assert.Nil(t, in.GetOrgID())
				})).Return(ioutil.NopCloser(strings.NewReader(fakeResults)), nil)
			},
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

			queryApi := mock.NewMockQueryApi(ctrl)
			if tc.registerExpectations != nil {
				tc.registerExpectations(t, queryApi)
			}
			cli := query.Client{
				CLI:           clients.CLI{ActiveConfig: config.Config{Org: tc.configOrgName}, StdIO: stdio},
				QueryApi:      queryApi,
				ResultPrinter: query.RawResultPrinter,
			}

			err := cli.Query(context.Background(), &tc.params)
			if tc.expectInErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectInErr)
				require.Empty(t, writtenBytes.String())
				return
			}
			require.NoError(t, err)
			require.Equal(t, fakeResults, writtenBytes.String())
		})
	}
}
