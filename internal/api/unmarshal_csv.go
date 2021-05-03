package api

import (
	"fmt"
	"strings"
)

// UnmarshalCSV implements the gocsv.TypeUnmarshaller interface
// for decoding CSV.
func (v *ColumnSemanticType) UnmarshalCSV(s string) error {
	types := []string{string(COLUMNSEMANTICTYPE_TIMESTAMP), string(COLUMNSEMANTICTYPE_TAG), string(COLUMNSEMANTICTYPE_FIELD)}
	for _, t := range types {
		if s == t {
			*v = ColumnSemanticType(t)
			return nil
		}
	}
	return fmt.Errorf("%q is not a valid column type. Valid values are [%s]", s, strings.Join(types, ", "))
}

// UnmarshalCSV implements the gocsv.TypeUnmarshaller interface
// for decoding CSV.
func (v *ColumnDataType) UnmarshalCSV(s string) error {
	types := []string{
		string(COLUMNDATATYPE_INTEGER),
		string(COLUMNDATATYPE_FLOAT),
		string(COLUMNDATATYPE_BOOLEAN),
		string(COLUMNDATATYPE_STRING),
		string(COLUMNDATATYPE_UNSIGNED),
	}

	for _, t := range types {
		if s == t {
			*v = ColumnDataType(t)
			return nil
		}
	}
	return fmt.Errorf("%q is not a valid column data type. Valid values are [%s]", s, strings.Join(types, ", "))
}
