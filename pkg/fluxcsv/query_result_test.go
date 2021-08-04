package fluxcsv_test

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/influxdata/influx-cli/v2/pkg/fluxcsv"
	"github.com/stretchr/testify/require"
)

func TestQueryCVSResultSingleTable(t *testing.T) {
	csvTable := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string
#group,false,false,true,true,false,false,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T10:34:08.135814545Z,1.4,f,test,1,adsfasdf
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.850214724Z,6.6,f,test,1,adsfasdf

`
	expectedTable := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_stop", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.DoubleDatatype, "", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_measurement", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "a", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "b", true),
	)
	expectedRecord1, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T10:34:08.135814545Z"),
			"_value":       1.4,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	expectedRecord2, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.850214724Z"),
			"_value":       6.6,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	reader := strings.NewReader(csvTable)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord1, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())
	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord2, queryResult.Record())

	require.False(t, queryResult.Next())
	require.Nil(t, queryResult.Err())
}

func TestQueryCVSResultMultiTables(t *testing.T) {
	csvTable := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string
#group,false,false,true,true,false,false,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T10:34:08.135814545Z,1.4,f,test,1,adsfasdf
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.850214724Z,6.6,f,test,1,adsfasdf

#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,long,string,string,string,string
#group,false,false,true,true,false,false,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,1,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T10:34:08.135814545Z,4,i,test,1,adsfasdf
,,1,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.850214724Z,-1,i,test,1,adsfasdf

#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,boolean,string,string,string,string
#group,false,false,true,true,false,false,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,2,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.62797864Z,false,f,test,0,adsfasdf
,,2,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.969100374Z,true,f,test,0,adsfasdf

#datatype,string,long,dateTime:RFC3339Nano,dateTime:RFC3339Nano,dateTime:RFC3339Nano,unsignedLong,string,string,string,string
#group,false,false,true,true,false,false,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,3,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.62797864Z,0,i,test,0,adsfasdf
,,3,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.969100374Z,2,i,test,0,adsfasdf

`
	expectedTable1 := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_stop", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.DoubleDatatype, "", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_measurement", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "a", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "b", true),
	)
	expectedRecord11, err := fluxcsv.NewFluxRecord(
		expectedTable1,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T10:34:08.135814545Z"),
			"_value":       1.4,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)
	expectedRecord12, err := fluxcsv.NewFluxRecord(
		expectedTable1,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.850214724Z"),
			"_value":       6.6,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	expectedTable2 := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_stop", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_measurement", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "a", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "b", true),
	)
	expectedRecord21, err := fluxcsv.NewFluxRecord(
		expectedTable2,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(1),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T10:34:08.135814545Z"),
			"_value":       int64(4),
			"_field":       "i",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)
	expectedRecord22, err := fluxcsv.NewFluxRecord(
		expectedTable2,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(1),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.850214724Z"),
			"_value":       int64(-1),
			"_field":       "i",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	expectedTable3 := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_stop", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.BoolDatatype, "", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_measurement", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "a", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "b", true),
	)
	expectedRecord31, err := fluxcsv.NewFluxRecord(
		expectedTable3,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(2),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.62797864Z"),
			"_value":       false,
			"_field":       "f",
			"_measurement": "test",
			"a":            "0",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)
	expectedRecord32, err := fluxcsv.NewFluxRecord(
		expectedTable3,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(2),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.969100374Z"),
			"_value":       true,
			"_field":       "f",
			"_measurement": "test",
			"a":            "0",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	expectedTable4 := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFCNano, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFCNano, "", "_stop", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFCNano, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.ULongDatatype, "", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_measurement", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "a", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "b", true),
	)
	expectedRecord41, err := fluxcsv.NewFluxRecord(
		expectedTable4,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(3),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.62797864Z"),
			"_value":       uint64(0),
			"_field":       "i",
			"_measurement": "test",
			"a":            "0",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)
	expectedRecord42, err := fluxcsv.NewFluxRecord(
		expectedTable4,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(3),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.969100374Z"),
			"_value":       uint64(2),
			"_field":       "i",
			"_measurement": "test",
			"a":            "0",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	reader := strings.NewReader(csvTable)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord11, queryResult.Record())
	require.True(t, queryResult.AnnotationsChanged())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())
	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord12, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord21, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord22, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err(), queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord31, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord32, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord41, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord42, queryResult.Record())

	require.False(t, queryResult.Next())
	require.Nil(t, queryResult.Err())
}

func TestQueryCVSResultSingleTableMultiColumnsNoValue(t *testing.T) {
	csvTable := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,long,string,duration,base64Binary,dateTime:RFC3339
#group,false,false,true,true,false,true,true,false,false,false
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z
,,1,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:39:36.330153686Z,1467463,BME280,1h20m30.13245s,eHh4eHhjY2NjY2NkZGRkZA==,2020-04-28T00:00:00Z
`
	expectedTable := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_stop", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "deviceId", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "sensor", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.DurationDatatype, "", "elapsed", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.Base64BinaryDataType, "", "note", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "start", false),
	)
	expectedRecord1, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":   "_result",
			"table":    int64(0),
			"_start":   mustParseTime(t, "2020-04-28T12:36:50.990018157Z"),
			"_stop":    mustParseTime(t, "2020-04-28T12:51:50.990018157Z"),
			"_time":    mustParseTime(t, "2020-04-28T12:38:11.480545389Z"),
			"deviceId": int64(1467463),
			"sensor":   "BME280",
			"elapsed":  time.Minute + time.Second,
			"note":     []byte("datainbase64"),
			"start":    time.Date(2020, 4, 27, 0, 0, 0, 0, time.UTC),
		},
	)
	require.NoError(t, err)

	expectedRecord2, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":   "_result",
			"table":    int64(1),
			"_start":   mustParseTime(t, "2020-04-28T12:36:50.990018157Z"),
			"_stop":    mustParseTime(t, "2020-04-28T12:51:50.990018157Z"),
			"_time":    mustParseTime(t, "2020-04-28T12:39:36.330153686Z"),
			"deviceId": int64(1467463),
			"sensor":   "BME280",
			"elapsed":  time.Hour + 20*time.Minute + 30*time.Second + 132450000*time.Nanosecond,
			"note":     []byte("xxxxxccccccddddd"),
			"start":    time.Date(2020, 4, 28, 0, 0, 0, 0, time.UTC),
		},
	)
	require.NoError(t, err)

	reader := strings.NewReader(csvTable)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord1, queryResult.Record())
	require.Nil(t, queryResult.Record().Value())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())
	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord2, queryResult.Record())

	require.False(t, queryResult.Next())
	require.Nil(t, queryResult.Err())
}

func TestErrorInRow(t *testing.T) {
	csvRowsError := []string{
		`#datatype,string,string`,
		`#group,true,true`,
		`#default,,`,
		`,error,reference`,
		`,failed to create physical plan: invalid time bounds from procedure from: bounds contain zero time,897`}
	csvTable := makeCSVstring(csvRowsError)
	reader := strings.NewReader(csvTable)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))

	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "failed to create physical plan: invalid time bounds from procedure from: bounds contain zero time,897", queryResult.Err().Error())

	csvRowsErrorNoReference := []string{
		`#datatype,string,string`,
		`#group,true,true`,
		`#default,,`,
		`,error,reference`,
		`,failed to create physical plan: invalid time bounds from procedure from: bounds contain zero time,`}
	csvTable = makeCSVstring(csvRowsErrorNoReference)
	reader = strings.NewReader(csvTable)
	queryResult = fluxcsv.NewQueryTableResult(io.NopCloser(reader))

	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "failed to create physical plan: invalid time bounds from procedure from: bounds contain zero time", queryResult.Err().Error())

	csvRowsErrorNoMessage := []string{
		`#datatype,string,string`,
		`#group,true,true`,
		`#default,,`,
		`,error,reference`,
		`,,`}
	csvTable = makeCSVstring(csvRowsErrorNoMessage)
	reader = strings.NewReader(csvTable)
	queryResult = fluxcsv.NewQueryTableResult(io.NopCloser(reader))

	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "unknown query error", queryResult.Err().Error())
}

func TestInvalidDataType(t *testing.T) {
	csvTable := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,int,string,duration,base64Binary,dateTime:RFC3339
#group,false,false,true,true,false,true,true,false,false,false
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:39:36.330153686Z,1467463,BME280,1h20m30.13245s,eHh4eHhjY2NjY2NkZGRkZA==,2020-04-28T00:00:00Z
`

	reader := strings.NewReader(csvTable)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "unknown data type int", queryResult.Err().Error())
}

func TestReorderedAnnotations(t *testing.T) {
	expectedTable := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_stop", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.DoubleDatatype, "", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_measurement", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "a", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "b", true),
	)
	expectedRecord1, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T10:34:08.135814545Z"),
			"_value":       1.4,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	expectedRecord2, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":       "_result",
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.850214724Z"),
			"_value":       6.6,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	csvTable1 := `#group,false,false,true,true,false,false,true,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T10:34:08.135814545Z,1.4,f,test,1,adsfasdf
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.850214724Z,6.6,f,test,1,adsfasdf

`
	reader := strings.NewReader(csvTable1)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord1, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())
	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord2, queryResult.Record())

	require.False(t, queryResult.Next())
	require.Nil(t, queryResult.Err())

	csvTable2 := `#default,_result,,,,,,,,,
#group,false,false,true,true,false,false,true,true,true,true
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T10:34:08.135814545Z,1.4,f,test,1,adsfasdf
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.850214724Z,6.6,f,test,1,adsfasdf

`
	reader = strings.NewReader(csvTable2)
	queryResult = fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord1, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())
	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord2, queryResult.Record())

	require.False(t, queryResult.Next())
	require.Nil(t, queryResult.Err())
}

func TestDatatypeOnlyAnnotation(t *testing.T) {
	expectedTable := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_stop", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_time", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.DoubleDatatype, "", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_measurement", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "a", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "b", false),
	)
	expectedRecord1, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":       nil,
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T10:34:08.135814545Z"),
			"_value":       1.4,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	expectedRecord2, err := fluxcsv.NewFluxRecord(
		expectedTable,
		map[string]interface{}{
			"result":       nil,
			"table":        int64(0),
			"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
			"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
			"_time":        mustParseTime(t, "2020-02-18T22:08:44.850214724Z"),
			"_value":       6.6,
			"_field":       "f",
			"_measurement": "test",
			"a":            "1",
			"b":            "adsfasdf",
		},
	)
	require.NoError(t, err)

	csvTable1 := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T10:34:08.135814545Z,1.4,f,test,1,adsfasdf
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.850214724Z,6.6,f,test,1,adsfasdf

`
	reader := strings.NewReader(csvTable1)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.True(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord1, queryResult.Record())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())
	require.False(t, queryResult.AnnotationsChanged())
	require.NotNil(t, queryResult.Record())
	require.Equal(t, expectedRecord2, queryResult.Record())

	require.False(t, queryResult.Next())
	require.Nil(t, queryResult.Err())
}

func TestMissingDatatypeAnnotation(t *testing.T) {
	csvTable1 := `
#group,false,false,true,true,false,true,true,false,false,false
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:39:36.330153686Z,1467463,BME280,1h20m30.13245s,eHh4eHhjY2NjY2NkZGRkZA==,2020-04-28T00:00:00Z
`

	reader := strings.NewReader(csvTable1)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "parsing error, datatype annotation not found", queryResult.Err().Error())

	csvTable2 := `
#default,_result,,,,,,,,,
#group,false,false,true,true,false,true,true,false,false,false
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:39:36.330153686Z,1467463,BME280,1h20m30.13245s,eHh4eHhjY2NjY2NkZGRkZA==,2020-04-28T00:00:00Z
`

	reader = strings.NewReader(csvTable2)
	queryResult = fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "parsing error, datatype annotation not found", queryResult.Err().Error())
}

func TestMissingAnnotations(t *testing.T) {
	csvTable3 := `
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:39:36.330153686Z,1467463,BME280,1h20m30.13245s,eHh4eHhjY2NjY2NkZGRkZA==,2020-04-28T00:00:00Z

`
	reader := strings.NewReader(csvTable3)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "parsing error, annotations not found", queryResult.Err().Error())
}

func TestDifferentNumberOfColumns(t *testing.T) {
	csvTable := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,long,string,duration,base64Binary,dateTime:RFC3339
#group,false,false,true,true,false,true,true,false,false,
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z,2345234
`

	reader := strings.NewReader(csvTable)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "parsing error, row has different number of columns than the table: 11 vs 10", queryResult.Err().Error())

	csvTable2 := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,long,string,duration,base64Binary,dateTime:RFC3339
#group,false,false,true,true,false,true,true,false,false,
#default,_result,,,,,,,
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z,2345234
`

	reader = strings.NewReader(csvTable2)
	queryResult = fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "parsing error, row has different number of columns than the table: 8 vs 10", queryResult.Err().Error())

	csvTable3 := `#default,_result,,,,,,,
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,long,string,duration,base64Binary,dateTime:RFC3339
#group,false,false,true,true,false,true,true,false,false,
,result,table,_start,_stop,_time,deviceId,sensor,elapsed,note,start
,,0,2020-04-28T12:36:50.990018157Z,2020-04-28T12:51:50.990018157Z,2020-04-28T12:38:11.480545389Z,1467463,BME280,1m1s,ZGF0YWluYmFzZTY0,2020-04-27T00:00:00Z,2345234
`

	reader = strings.NewReader(csvTable3)
	queryResult = fluxcsv.NewQueryTableResult(io.NopCloser(reader))
	require.False(t, queryResult.Next())
	require.NotNil(t, queryResult.Err())
	require.Equal(t, "parsing error, row has different number of columns than the table: 10 vs 8", queryResult.Err().Error())
}

func TestEmptyValue(t *testing.T) {
	csvTable := `#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string
#group,false,false,true,true,false,false,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,a,b
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T10:34:08.135814545Z,,f,test,1,adsfasdf
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:08:44.850214724Z,6.6,f,test,,adsfasdf
,,0,2020-02-17T22:19:49.747562847Z,2020-02-18T22:19:49.747562847Z,2020-02-18T22:11:32.225467895Z,1122.45,f,test,3,
`

	reader := strings.NewReader(csvTable)
	queryResult := fluxcsv.NewQueryTableResult(io.NopCloser(reader))

	require.True(t, queryResult.Next(), queryResult.Err())
	require.Nil(t, queryResult.Err())

	require.NotNil(t, queryResult.Record())
	require.Nil(t, queryResult.Record().Value())

	require.True(t, queryResult.Next(), queryResult.Err())
	require.NotNil(t, queryResult.Record())
	require.Nil(t, queryResult.Record().ValueByKey("a"))

	require.True(t, queryResult.Next(), queryResult.Err())
	require.NotNil(t, queryResult.Record())
	require.Nil(t, queryResult.Record().ValueByKey("b"))

	require.False(t, queryResult.Next())
	require.Nil(t, queryResult.Err())
}

func makeCSVstring(rows []string) string {
	csvTable := strings.Join(rows, "\r\n")
	return fmt.Sprintf("%s\r\n", csvTable)
}
