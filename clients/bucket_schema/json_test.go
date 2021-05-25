package bucket_schema

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeJson(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		exp    []api.MeasurementSchemaColumn
		expErr string
	}{
		{
			name: "valid",
			data: heredoc.Doc(`
				[
					{ "name": "time", "type": "timestamp" },
					{ "name": "host", "type": "tag" },
					{ "name": "usage_user", "type": "field", "dataType": "float" }
				]
			`),
			exp: cols(colTs(), colT("host"), colFF("usage_user")),
		},
		{
			name: "invalid column type",
			data: heredoc.Doc(`
				[
					{ "name": "time", "type": "foo" }
				]
			`),
			expErr: `error decoding JSON: foo is not a valid ColumnSemanticType`,
		},
		{
			name: "invalid column data type",
			data: heredoc.Doc(`
				[
					{ "name": "time", "type": "field", "dataType": "floaty" }
				]
			`),
			expErr: `error decoding JSON: floaty is not a valid ColumnDataType`,
		},
		{
			name: "invalid JSON",
			data: heredoc.Doc(`
				[
					{ "name": "time", "type": "field", "dataType": "floaty" }
			`),
			expErr: `error decoding JSON: unexpected EOF`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.data)
			got, err := decodeJson(r)
			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.exp, got)
			}
		})
	}
}
