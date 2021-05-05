package bucket_schema

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeCSV(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		exp    []api.MeasurementSchemaColumn
		expErr string
	}{
		{
			name: "valid",
			data: heredoc.Doc(`
				name,type,data_type
				time,timestamp,
				host,tag,
				usage_user,field,float
			`),
			exp: cols(colTs(), colT("host"), colFF("usage_user")),
		},
		{
			name: "valid with alternate order",
			data: heredoc.Doc(`
				type,data_type,name
				timestamp,,time
				tag,,host
				field,float,usage_user
			`),
			exp: cols(colTs(), colT("host"), colFF("usage_user")),
		},
		{
			name: "invalid column type",
			data: heredoc.Doc(`
				name,type,data_type
				time,foo,
			`),
			expErr: `failed to decode CSV: record on line 0; parse error on line 2, column 2: "foo" is not a valid column type. Valid values are [timestamp, tag, field]`,
		},
		{
			name: "invalid column data type",
			data: heredoc.Doc(`
				name,type,data_type
				time,field,floaty
			`),
			expErr: `failed to decode CSV: record on line 0; parse error on line 2, column 3: "floaty" is not a valid column data type. Valid values are [integer, float, boolean, string, unsigned]`,
		},
		{
			name: "invalid headers",
			data: heredoc.Doc(`
				name,foo
				time,field
			`),
			expErr: `failed to decode CSV: found unmatched struct field with tags [type data_type]`,
		},
		{
			name: "invalid CSV",
			data: heredoc.Doc(`
				type,type,data_type
				time,timestamp
			`),
			expErr: `failed to decode CSV: record on line 2: wrong number of fields`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.data)
			got, err := decodeCSV(r)
			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.exp, got)
			}
		})
	}
}
