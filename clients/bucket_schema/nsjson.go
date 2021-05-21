package bucket_schema

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/api"
)

func decodeNDJson(r io.Reader) ([]api.MeasurementSchemaColumn, error) {
	scan := bufio.NewScanner(r)
	var rows []api.MeasurementSchemaColumn
	n := 1
	for scan.Scan() {
		if line := bytes.TrimSpace(scan.Bytes()); len(line) > 0 {
			var row api.MeasurementSchemaColumn
			if err := json.Unmarshal(line, &row); err != nil {
				return nil, fmt.Errorf("error decoding JSON at line %d: %w", n, err)
			}
			rows = append(rows, row)
		}
		n++
	}
	return rows, nil
}
