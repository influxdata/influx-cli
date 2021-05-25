package bucket_schema

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/api"
)

func decodeJson(r io.Reader) ([]api.MeasurementSchemaColumn, error) {
	var rows []api.MeasurementSchemaColumn
	if err := json.NewDecoder(r).Decode(&rows); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	return rows, nil
}
