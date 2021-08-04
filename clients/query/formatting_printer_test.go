package query_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/clients/query"
	"github.com/stretchr/testify/require"
)

func TestFormatterPrinter_PrintQueryResults(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "empty",
			in:       "",
			expected: "",
		},
		{
			name: "single table",
			in: `#group,false,false,true,true,false,false,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string
#default,_result,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,bar
,,0,1921-05-08T15:46:22.507379Z,2021-05-08T14:46:22.507379Z,2021-05-04T18:29:52.764702Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:46:22.507379Z,2021-05-08T14:46:22.507379Z,2021-05-04T19:30:59.67555Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:46:22.507379Z,2021-05-08T14:46:22.507379Z,2021-05-04T19:31:01.876079Z,12345,qux,foo,"""baz"""
,,0,1921-05-08T15:46:22.507379Z,2021-05-08T14:46:22.507379Z,2021-05-04T19:31:02.499461Z,12345,qux,foo,"""baz"""
`,
			expected: `Result: _result
Table: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                      _time:time                  _value:float
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ------------------------------  ----------------------------
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T18:29:52.764702000Z                         12345
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T19:30:59.675550000Z                         12345
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T19:31:01.876079000Z                         12345
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T19:31:02.499461000Z                         12345
`,
		},
		{
			name: "nil values",
			in: `#group,false,false,false,false,true,false,false
#datatype,string,long,string,string,string,string,long
#default,_result,,,,,,
,result,table,name,id,organizationID,retentionPolicy,retentionPeriod
,,0,_monitoring,1aa1e247d56a143f,b6b9cb281ae9583d,,604800000000000
,,0,_tasks,e03361698294077c,b6b9cb281ae9583d,,259200000000000
,,0,dan,57de01a0f4825d94,b6b9cb281ae9583d,,259200000000000
`,
			expected: `Result: _result
Table: keys: [organizationID]
 organizationID:string             name:string               id:string  retentionPolicy:string         retentionPeriod:int
----------------------  ----------------------  ----------------------  ----------------------  --------------------------
      b6b9cb281ae9583d             _monitoring        1aa1e247d56a143f                                     604800000000000
      b6b9cb281ae9583d                  _tasks        e03361698294077c                                     259200000000000
      b6b9cb281ae9583d                     dan        57de01a0f4825d94                                     259200000000000
`,
		},
		{
			name: "multi table",
			in: `#group,false,false,true,true,false,false,true,true,true
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
			expected: `Result: _result
Table: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                      _time:time                  _value:float
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ------------------------------  ----------------------------
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T18:29:52.764702000Z                         12345
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T19:30:59.675550000Z                         12345
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T19:31:01.876079000Z                         12345
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T19:31:02.499461000Z                         12345
Table: keys: [_start, _stop, _field, _measurement, bar, is_foo]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string           is_foo:string                      _time:time                  _value:float
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ----------------------  ------------------------------  ----------------------------
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"                       t  2021-05-08T15:42:19.567667000Z                         12345
Result: _profiler
Table: keys: [_measurement]
   _measurement:string                 Type:string            Label:string                   Count:int             MinDuration:int             MaxDuration:int             DurationSum:int            MeanDuration:float
----------------------  --------------------------  ----------------------  --------------------------  --------------------------  --------------------------  --------------------------  ----------------------------
     profiler/operator  *influxdb.readFilterSource              ReadRange2                           1                      367331                      367331                      367331                        367331
`,
		},
		{
			name: "multi table single result",
			in: `#group,false,false,true,true,false,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string
#default,_result,,,,,,,
,result,table,_start,_stop,_value,_field,_measurement,bar
,,0,2021-05-12T20:11:00Z,2021-05-12T20:12:00Z,518490,qux,foo,"""baz"""
,,1,2021-05-12T20:12:00Z,2021-05-12T20:13:00Z,703665,qux,foo,"""baz"""
,,2,2021-05-12T20:13:00Z,2021-05-12T20:14:00Z,703665,qux,foo,"""baz"""
,,3,2021-05-12T20:14:00Z,2021-05-12T20:15:00Z,444420,qux,foo,"""baz"""`,
			expected: `Result: _result
Table: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                  _value:float
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ----------------------------
2021-05-12T20:11:00.000000000Z  2021-05-12T20:12:00.000000000Z                     qux                     foo                   "baz"                        518490
Table: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                  _value:float
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ----------------------------
2021-05-12T20:12:00.000000000Z  2021-05-12T20:13:00.000000000Z                     qux                     foo                   "baz"                        703665
Table: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                  _value:float
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ----------------------------
2021-05-12T20:13:00.000000000Z  2021-05-12T20:14:00.000000000Z                     qux                     foo                   "baz"                        703665
Table: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                  _value:float
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ----------------------------
2021-05-12T20:14:00.000000000Z  2021-05-12T20:15:00.000000000Z                     qux                     foo                   "baz"                        444420
`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			in := io.NopCloser(strings.NewReader(tc.in))
			out := bytes.Buffer{}
			require.NoError(t, query.NewFormattingPrinter().PrintQueryResults(in, &out))
			require.Equal(t, tc.expected, out.String())
		})
	}
}
