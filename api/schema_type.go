package api

import (
	"fmt"
)

func (v SchemaType) String() string {
	return string(v)
}

// Set implements the cli.Generic interface for parsing
// flags.
func (v *SchemaType) Set(s string) error {
	switch s {
	case string(SCHEMATYPE_IMPLICIT):
		*v = SCHEMATYPE_IMPLICIT
	case string(SCHEMATYPE_EXPLICIT):
		*v = SCHEMATYPE_EXPLICIT
	default:
		return fmt.Errorf("unsupported schema type: %q", s)
	}
	return nil
}
