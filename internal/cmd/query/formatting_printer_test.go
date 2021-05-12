package query_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/influxdata/influx-cli/v2/internal/cmd/query"
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
TableId: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                      _time:time                  _value:float  
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ------------------------------  ----------------------------  
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T18:29:52.764702000Z                         12345  
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T19:30:59.675550000Z                         12345  
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T19:31:01.876079000Z                         12345  
1921-05-08T15:46:22.507379000Z  2021-05-08T14:46:22.507379000Z                     qux                     foo                   "baz"  2021-05-04T19:31:02.499461000Z                         12345  
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
TableId: keys: [_start, _stop, _field, _measurement, bar]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string                      _time:time                  _value:float  
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ------------------------------  ----------------------------  
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T18:29:52.764702000Z                         12345  
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T19:30:59.675550000Z                         12345  
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T19:31:01.876079000Z                         12345  
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"  2021-05-04T19:31:02.499461000Z                         12345  
TableId: keys: [_start, _stop, _field, _measurement, bar, is_foo]
                   _start:time                      _stop:time           _field:string     _measurement:string              bar:string           is_foo:string                      _time:time                  _value:float  
------------------------------  ------------------------------  ----------------------  ----------------------  ----------------------  ----------------------  ------------------------------  ----------------------------  
1921-05-08T15:42:58.218436000Z  2021-05-08T15:42:58.218436000Z                     qux                     foo                   "baz"                       t  2021-05-08T15:42:19.567667000Z                         12345  
Result: _profiler
TableId: keys: [_measurement]
   _measurement:string                 Type:string            Label:string                   Count:int             MinDuration:int             MaxDuration:int             DurationSum:int            MeanDuration:float  
----------------------  --------------------------  ----------------------  --------------------------  --------------------------  --------------------------  --------------------------  ----------------------------  
     profiler/operator  *influxdb.readFilterSource              ReadRange2                           1                      367331                      367331                      367331                        367331  
`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			in := ioutil.NopCloser(strings.NewReader(tc.in))
			out := bytes.Buffer{}
			require.NoError(t, query.NewFormattingPrinter().PrintQueryResults(in, &out))
			require.Equal(t, tc.expected, out.String())
		})
	}
}
