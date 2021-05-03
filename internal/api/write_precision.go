package api

import (
	"fmt"
)

func (v WritePrecision) String() string {
	return string(v)
}

func (v *WritePrecision) Set(s string) error {
	switch s {
	case "ms":
		*v = WRITEPRECISION_MS
	case "s":
		*v = WRITEPRECISION_S
	case "us":
		*v = WRITEPRECISION_US
	case "ns":
		*v = WRITEPRECISION_NS
	default:
		return fmt.Errorf("unsupported precision: %q", s)
	}
	return nil
}
