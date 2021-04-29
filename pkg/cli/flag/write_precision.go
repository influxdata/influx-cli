package flag

import (
	"fmt"

	"github.com/influxdata/influx-cli/v2/internal/api"
)

type WritePrecision api.WritePrecision

func WritePrecisionVar(v *api.WritePrecision) *WritePrecision {
	return (*WritePrecision)(v)
}

func (i WritePrecision) String() string {
	return string(i)
}

func (i *WritePrecision) Set(s string) error {
	switch s {
	case "ms":
		*i = WritePrecision(api.WRITEPRECISION_MS)
	case "s":
		*i = WritePrecision(api.WRITEPRECISION_S)
	case "us":
		*i = WritePrecision(api.WRITEPRECISION_US)
	case "ns":
		*i = WritePrecision(api.WRITEPRECISION_NS)
	default:
		return fmt.Errorf("unsupported precision: %q", s)
	}
	return nil
}
