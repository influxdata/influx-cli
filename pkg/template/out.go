package template

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

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

func ParseOutParams(path string, fallback io.Writer) (OutParams, func(), error){
	if path == "" {
		return OutParams{Out: fallback, Encoding: YamlEncoding}, nil, nil
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return OutParams{}, nil, fmt.Errorf("failed to open output path %q: %w", path, err)
	}
	params := OutParams{Out: f}
	switch filepath.Ext(path) {
	case ".json":
		params.Encoding = JsonEncoding
	default:
		params.Encoding = YamlEncoding
	}

	return params, func() { _ = f.Close() }, nil
}

func (o OutParams) WriteTemplate(template []api.TemplateEntry) error {
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
