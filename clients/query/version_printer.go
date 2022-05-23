package query

import (
	"fmt"
	"io"

	"github.com/influxdata/influx-cli/v2/pkg/fluxcsv"
)

type versionPrinter struct{}

// VersionPrinter reads the result of the version query and outputs a single line.
var VersionPrinter ResultPrinter = &versionPrinter{}

func (v versionPrinter) PrintQueryResults(resultStream io.ReadCloser, out io.Writer) error {
	res := fluxcsv.NewQueryTableResult(resultStream)
	defer res.Close()

	// Only one row.
	if !res.Next() {
		return res.Err()
	}

	// The key doesn't matter except that it's not table.
	record := res.Record()
	for k, v := range record.Values() {
		if k == "table" {
			continue
		}
		_, _ = fmt.Fprintf(out, "flux %s\n", v)
		break
	}
	return nil
}
