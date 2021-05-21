package bucket_schema

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func colTs() api.MeasurementSchemaColumn {
	return api.MeasurementSchemaColumn{Name: "time", Type: api.COLUMNSEMANTICTYPE_TIMESTAMP}
}

func colT(n string) api.MeasurementSchemaColumn {
	return api.MeasurementSchemaColumn{Name: n, Type: api.COLUMNSEMANTICTYPE_TAG}
}

func colFF(n string) api.MeasurementSchemaColumn {
	return api.MeasurementSchemaColumn{Name: n, Type: api.COLUMNSEMANTICTYPE_FIELD, DataType: api.COLUMNDATATYPE_FLOAT.Ptr()}
}

func cols(c ...api.MeasurementSchemaColumn) []api.MeasurementSchemaColumn { return c }

func TestDecodeNDJson(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		exp    []api.MeasurementSchemaColumn
		expErr string
	}{
		{
			name: "valid",
			data: heredoc.Doc(`
				{ "name": "time", "type": "timestamp" }
				{ "name": "host", "type": "tag" }
				{ "name": "usage_user", "type": "field", "dataType": "float" }
			`),
			exp: cols(colTs(), colT("host"), colFF("usage_user")),
		},
		{
			name: "invalid column type",
			data: heredoc.Doc(`
				{ "name": "time", "type": "foo" }
			`),
			expErr: `error decoding JSON at line 1: foo is not a valid ColumnSemanticType`,
		},
		{
			name: "invalid column data type",
			data: heredoc.Doc(`
				{ "name": "time", "type": "field", "dataType": "floaty" }
			`),
			expErr: `error decoding JSON at line 1: floaty is not a valid ColumnDataType`,
		},
		{
			name: "invalid JSON",
			data: heredoc.Doc(`
				{ "name": "usage_user", "type": "field", "dataType": "float" }
				{ "name": "time", "type": "field", "dataType": "float"
			`),
			expErr: `error decoding JSON at line 2: unexpected end of JSON input`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.data)
			got, err := decodeNDJson(r)
			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.exp, got)
			}
		})
	}
}
