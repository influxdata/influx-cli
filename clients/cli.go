package clients

import (
	"encoding/json"

	"github.com/influxdata/influx-cli/v2/config"
	"github.com/influxdata/influx-cli/v2/internal/tabwriter"
	"github.com/influxdata/influx-cli/v2/pkg/stdio"
)

// CLI is a container for common functionality used to execute commands.
type CLI struct {
	StdIO stdio.StdIO

	HideTableHeaders bool
	PrintAsJSON      bool

	ActiveConfig  config.Config
	ConfigService config.Service
}

func (c *CLI) PrintJSON(v interface{}) error {
	enc := json.NewEncoder(c.StdIO)
	enc.SetIndent("", "\t")
	return enc.Encode(v)
}

func (c *CLI) PrintTable(headers []string, rows ...map[string]interface{}) error {
	w := tabwriter.NewTabWriter(c.StdIO, c.HideTableHeaders)
	defer w.Flush()
	if err := w.WriteHeaders(headers...); err != nil {
		return err
	}
	for _, r := range rows {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	return nil
}
