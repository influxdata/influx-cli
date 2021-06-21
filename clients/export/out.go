package export

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/api"
	"gopkg.in/yaml.v3"
)

type OutEncoding int

const (
	YamlEncoding OutEncoding = iota
	JsonEncoding
)

type OutParams struct {
	Out      io.Writer
	Encoding OutEncoding
}

func (o OutParams) writeTemplate(template []api.TemplateEntry) error {
	switch o.Encoding {
	case JsonEncoding:
		enc := json.NewEncoder(o.Out)
		enc.SetIndent("", "\t")
		return enc.Encode(template)
	case YamlEncoding:
		enc := yaml.NewEncoder(o.Out)
		for _, entry := range template {
			if err := enc.Encode(entry); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("encoding %q is not recognized", o.Encoding)
	}
	return nil
}
