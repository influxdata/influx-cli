package template_test

import (
	"bytes"
	"testing"

	"github.com/influxdata/influx-cli/v2/pkg/template"
	"github.com/stretchr/testify/require"
)

func TestDiffPrinter_Empty(t *testing.T) {
	t.Parallel()

	out := bytes.Buffer{}
	printer := template.NewDiffPrinter(&out, false, true).
		Title("Example").
		SetHeaders("Wow", "Such", "A", "Fancy", "Printer")

	printer.Render()
	require.Empty(t, out.String())
}

func TestDiffPrinter(t *testing.T) {
	t.Parallel()

	out := bytes.Buffer{}
	printer := template.NewDiffPrinter(&out, false, true).
		Title("Example").
		SetHeaders("Wow", "Such", "A", "Fancy", "Printer")

	// Add
	printer.AppendDiff(nil, []string{"A", "B", "C", "D", "E"})

	// No change
	printer.Append([]string{"foo", "bar", "baz", "qux", "wat"})

	// Replace
	printer.AppendDiff(
		[]string{"1", "200000000000000", "3", "4", "5"},
		[]string{"9", "8", "7", "6", "5"},
	)

	// Remove
	printer.AppendDiff([]string{"x y", "z x", "x y z", "", "y z"}, nil)

	printer.Render()
	expected := `EXAMPLE    +add | -remove | unchanged
+-----+-----+-----------------+-------+-------+---------+
| +/- | WOW |      SUCH       |   A   | FANCY | PRINTER |
+-----+-----+-----------------+-------+-------+---------+
| +   | A   | B               | C     | D     | E       |
+-----+-----+-----------------+-------+-------+---------+
|     | foo | bar             | baz   | qux   | wat     |
+-----+-----+-----------------+-------+-------+---------+
+-----+-----+-----------------+-------+-------+---------+
| -   |   1 | 200000000000000 |     3 |     4 |       5 |
+-----+-----+-----------------+-------+-------+---------+
| +   |   9 |               8 |     7 |     6 |       5 |
+-----+-----+-----------------+-------+-------+---------+
+-----+-----+-----------------+-------+-------+---------+
| -   | x y | z x             | x y z |       | y z     |
+-----+-----+-----------------+-------+-------+---------+
|                                       TOTAL |    3    |
+-----+-----+-----------------+-------+-------+---------+
`
	require.Equal(t, expected, out.String())
}
