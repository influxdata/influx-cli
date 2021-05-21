package bucket_schema

import (
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/influxdata/influx-cli/v2/api"
)

type csvColumn struct {
	Name     string                 `csv:"name"`
	Type     api.ColumnSemanticType `csv:"type"`
	DataType *api.ColumnDataType    `csv:"data_type,omitempty"`
}

func init() {
	gocsv.FailIfUnmatchedStructTags = true
}

func decodeCSV(r io.Reader) ([]api.MeasurementSchemaColumn, error) {
	var cols []csvColumn
	err := gocsv.Unmarshal(r, &cols)
	if err != nil {
		return nil, fmt.Errorf("failed to decode CSV: %w", err)
	}

	rows := make([]api.MeasurementSchemaColumn, 0, len(cols))
	for i := range cols {
		c := &cols[i]
		rows = append(rows, api.MeasurementSchemaColumn{
			Name:     c.Name,
			Type:     c.Type,
			DataType: c.DataType,
		})
	}
	return rows, nil
}
