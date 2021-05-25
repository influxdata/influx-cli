package bucket_schema

import (
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/influxdata/influx-cli/v2/api"
)

// ColumnsFormat is a type which defines the supported formats
//
type ColumnsFormat int

const (
	ColumnsFormatAuto ColumnsFormat = iota
	ColumnsFormatCSV
	ColumnsFormatNDJson
	ColumnsFormatJson
)

func (f *ColumnsFormat) Set(v string) error {
	switch v {
	case "auto":
		*f = ColumnsFormatAuto
	case "csv":
		*f = ColumnsFormatCSV
	case "ndjson":
		*f = ColumnsFormatNDJson
	case "json":
		*f = ColumnsFormatJson
	default:
		return fmt.Errorf("invalid columns-format: %s, expected [csv, ndjson, json, auto]", v)
	}
	return nil
}

func (f ColumnsFormat) String() string {
	switch f {
	case ColumnsFormatAuto:
		return "auto"
	case ColumnsFormatCSV:
		return "csv"
	case ColumnsFormatNDJson:
		return "ndjson"
	case ColumnsFormatJson:
		return "json"
	default:
		return "schema.Format(" + strconv.FormatInt(int64(f), 10) + ")"
	}
}

// DecoderFn uses f and path to return a function capable of decoding
// measurement schema columns from a given io.Reader. If no combination
// of decoder exists for f and path, DecoderFn returns an error.
func (f ColumnsFormat) DecoderFn(path string) (ColumnsDecoderFn, error) {
	ff := f
	if ff == ColumnsFormatAuto {
		ext := filepath.Ext(path)
		switch {
		case strings.EqualFold(ext, ".csv"):
			ff = ColumnsFormatCSV
		case strings.EqualFold(ext, ".json"):
			ff = ColumnsFormatJson
		case strings.EqualFold(ext, ".ndjson") || strings.EqualFold(ext, ".jsonl"):
			ff = ColumnsFormatNDJson
		}
	}

	switch ff {
	case ColumnsFormatCSV:
		return decodeCSV, nil
	case ColumnsFormatNDJson:
		return decodeNDJson, nil
	case ColumnsFormatJson:
		return decodeJson, nil
	}

	return nil, fmt.Errorf("unable to guess format for file %q", path)
}

// ColumnsDecoderFn is a function which decodes a slice of api.MeasurementSchemaColumn
// elements from r.
type ColumnsDecoderFn func(r io.Reader) ([]api.MeasurementSchemaColumn, error)
