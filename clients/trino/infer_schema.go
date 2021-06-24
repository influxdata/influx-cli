package trino

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/influxdata/influx-cli/v2/api"
	v1 "github.com/influxdata/influx-cli/v2/api/v1"
)

type InferSchemaParams struct {
	DB string
	RP string
}

type columnSchema struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	DataType string `json:"dataType,omitempty"`
}

type tableSchema struct {
	Name    string         `json:"name"`
	Columns []columnSchema `json:"columns"`
}

type databaseSchema map[string]*tableSchema

func (d databaseSchema) Measurements() []string {
	res := make([]string, 0, len(d))
	for k := range d {
		res = append(res, k)
	}
	return res
}

func (c Client) getCSV(ctx context.Context, database, query string) ([][]string, error) {
	res, err := c.QueryApi.GetQueryV1(ctx).
		Accept(v1.INFLUXQLCONTENTTYPE_TEXT_CSV).
		Db(database).
		Q(query).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to infer schema: %w", err)
	}

	r := csv.NewReader(strings.NewReader(res))
	return r.ReadAll()
}

func (c Client) addMeasurements(ctx context.Context, database string, schema databaseSchema) error {
	rows, err := c.getCSV(ctx, database, "SHOW MEASUREMENTS")
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	} else if len(rows) == 0 {
		return nil
	}

	for _, row := range rows[1:] {
		if len(row) < 3 {
			return fmt.Errorf("getMeasurements: unexpected number of columns")
		}

		schema[row[2]] = &tableSchema{
			Name: row[2],
			Columns: []columnSchema{
				// time column is implicit
				{Name: "time", Type: string(api.COLUMNSEMANTICTYPE_TIMESTAMP)},
			},
		}
	}

	return nil
}

func (c Client) addTagColumns(ctx context.Context, database string, schema databaseSchema) error {
	measurements := schema.Measurements()
	var query string
	if len(measurements) == 1 {
		query = fmt.Sprintf("SHOW TAG KEYS FROM %s", measurements[0])
	} else {
		query = fmt.Sprintf("SHOW TAG KEYS FROM /^(%s)$/", strings.Join(measurements, "|"))
	}

	rows, err := c.getCSV(ctx, database, query)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	} else if len(rows) == 0 {
		return nil
	}

	for _, row := range rows {
		if len(row) < 3 {
			return fmt.Errorf("getTagKeys: unexpected number of columns")
		}

		var (
			measurement = row[0]
			tagKey      = row[2]
		)
		if table := schema[measurement]; table != nil {
			table.Columns = append(table.Columns, columnSchema{
				Name: tagKey,
				Type: string(api.COLUMNSEMANTICTYPE_TAG),
			})
		}
	}

	return nil
}

func (c Client) addFieldColumns(ctx context.Context, database string, schema databaseSchema) error {
	measurements := schema.Measurements()
	var query string
	if len(measurements) == 1 {
		query = fmt.Sprintf("SHOW FIELD KEYS FROM %s", measurements[0])
	} else {
		query = fmt.Sprintf("SHOW FIELD KEYS FROM /^(%s)$/", strings.Join(measurements, "|"))
	}

	rows, err := c.getCSV(ctx, database, query)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	} else if len(rows) == 0 {
		return nil
	}

	for _, row := range rows {
		if len(row) < 4 {
			return fmt.Errorf("getTagKeys: unexpected number of columns")
		}

		// NOTE: fieldType is one of https://github.com/influxdata/influxql/blob/b9e521f450818862cab2bc31a30293b192439728/ast.go#L177-L186
		//  which is synonymous with the DataType field in a measurement schema column.

		var (
			measurement = row[0]
			fieldKey    = row[2]
			fieldType   = row[3]
		)
		if table := schema[measurement]; table != nil {
			table.Columns = append(table.Columns, columnSchema{
				Name:     fieldKey,
				Type:     string(api.COLUMNSEMANTICTYPE_FIELD),
				DataType: fieldType,
			})
		}
	}

	return nil
}

func (c Client) InferSchema(ctx context.Context, params *InferSchemaParams) error {
	schema := make(databaseSchema)
	err := c.addMeasurements(ctx, params.DB, schema)
	if err != nil {
		return err
	}

	err = c.addTagColumns(ctx, params.DB, schema)
	if err != nil {
		return err
	}

	err = c.addFieldColumns(ctx, params.DB, schema)
	if err != nil {
		return err
	}

	schemas := make([]*tableSchema, 0, len(schema))
	for _, table := range schema {
		schemas = append(schemas, table)
	}

	enc := json.NewEncoder(c.StdIO)
	return enc.Encode(schemas)
}

type InferExplicitBucketParams struct {
	Name string
	DB   string
	RP   string
}

func (c Client) InferExplicitBucket(ctx context.Context, params *InferExplicitBucketParams) error {
	schema := make(databaseSchema)
	err := c.addMeasurements(ctx, params.DB, schema)
	if err != nil {
		return err
	}

	err = c.addTagColumns(ctx, params.DB, schema)
	if err != nil {
		return err
	}

	err = c.addFieldColumns(ctx, params.DB, schema)
	if err != nil {
		return err
	}

	schemas := make([]*tableSchema, 0, len(schema))
	for _, table := range schema {
		schemas = append(schemas, table)
	}

	enc := json.NewEncoder(c.StdIO)
	return enc.Encode(schemas)
}
