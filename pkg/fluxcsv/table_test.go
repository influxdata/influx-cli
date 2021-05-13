package fluxcsv_test

import (
	"testing"
	"time"

	"github.com/influxdata/influx-cli/v2/pkg/fluxcsv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustParseTime(t *testing.T, s string) time.Time {
	t.Helper()
	time, err := time.Parse(time.RFC3339, s)
	require.NoError(t, err)
	return time
}

func TestTable(t *testing.T) {
	t.Parallel()

	table := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "10", "table", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.TimeDatatypeRFC, "", "_start", true),
		fluxcsv.NewFluxColumnFull(fluxcsv.DoubleDatatype, "1.1", "_value", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "", "_field", true),
	)
	require.Len(t, table.Columns(), 3)

	require.NotNil(t, table.ResultColumn())
	require.Equal(t, "_result", table.ResultColumn().DefaultValue())
	require.Equal(t, fluxcsv.StringDatatype, table.ResultColumn().DataType())
	require.Equal(t, "result", table.ResultColumn().Name())
	require.Equal(t, false, table.ResultColumn().IsGroup())

	require.NotNil(t, table.TableIdColumn())
	require.Equal(t, "10", table.TableIdColumn().DefaultValue())
	require.Equal(t, fluxcsv.LongDatatype, table.TableIdColumn().DataType())
	require.Equal(t, "table", table.TableIdColumn().Name())
	require.Equal(t, false, table.TableIdColumn().IsGroup())

	require.NotNil(t, table.Column(0))
	require.Equal(t, "", table.Column(0).DefaultValue())
	require.Equal(t, fluxcsv.TimeDatatypeRFC, table.Column(0).DataType())
	require.Equal(t, "_start", table.Column(0).Name())
	require.Equal(t, true, table.Column(0).IsGroup())

	require.NotNil(t, table.Column(1))
	require.Equal(t, "1.1", table.Column(1).DefaultValue())
	require.Equal(t, fluxcsv.DoubleDatatype, table.Column(1).DataType())
	require.Equal(t, "_value", table.Column(1).Name())
	require.Equal(t, false, table.Column(1).IsGroup())

	require.NotNil(t, table.Column(2))
	require.Equal(t, "", table.Column(2).DefaultValue())
	require.Equal(t, fluxcsv.StringDatatype, table.Column(2).DataType())
	require.Equal(t, "_field", table.Column(2).Name())
	require.Equal(t, true, table.Column(2).IsGroup())
}

func TestRecord(t *testing.T) {
	t.Parallel()

	table := fluxcsv.NewFluxTableMetadataFull(
		fluxcsv.NewFluxColumnFull(fluxcsv.StringDatatype, "_result", "result", false),
		fluxcsv.NewFluxColumnFull(fluxcsv.LongDatatype, "10", "table", false),
	)

	record, err := fluxcsv.NewFluxRecord(table, map[string]interface{}{
		"table":        int64(0),
		"_start":       mustParseTime(t, "2020-02-17T22:19:49.747562847Z"),
		"_stop":        mustParseTime(t, "2020-02-18T22:19:49.747562847Z"),
		"_time":        mustParseTime(t, "2020-02-18T10:34:08.135814545Z"),
		"_value":       1.4,
		"_field":       "f",
		"_measurement": "test",
		"a":            "1",
		"b":            "adsfasdf",
	})
	require.NoError(t, err)
	require.Len(t, record.Values(), 9)
	require.Equal(t, mustParseTime(t, "2020-02-17T22:19:49.747562847Z"), record.Start())
	require.Equal(t, mustParseTime(t, "2020-02-18T22:19:49.747562847Z"), record.Stop())
	require.Equal(t, mustParseTime(t, "2020-02-18T10:34:08.135814545Z"), record.Time())
	require.Equal(t, "f", record.Field())
	require.Equal(t, 1.4, record.Value())
	require.Equal(t, "test", record.Measurement())
	require.Equal(t, int64(0), record.TableId())

	agRec, err := fluxcsv.NewFluxRecord(table, map[string]interface{}{
		"result": "foo",
		"room":   "bathroom",
		"sensor": "SHT",
		"temp":   24.3,
		"hum":    42,
	})
	require.NoError(t, err)
	require.Len(t, agRec.Values(), 4)
	require.Equal(t, time.Time{}, agRec.Start())
	require.Equal(t, time.Time{}, agRec.Stop())
	require.Equal(t, time.Time{}, agRec.Time())
	require.Equal(t, "", agRec.Field())
	assert.Nil(t, agRec.Value())
	require.Equal(t, "", agRec.Measurement())
	require.Equal(t, int64(10), agRec.TableId())
	require.Equal(t, 24.3, agRec.ValueByKey("temp"))
	require.Equal(t, 42, agRec.ValueByKey("hum"))
	assert.Nil(t, agRec.ValueByKey("notexist"))
}
