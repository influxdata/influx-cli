package template_test

import (
	"bytes"
	"testing"

	"github.com/influxdata/influx-cli/v2/pkg/template"
	"github.com/stretchr/testify/require"
)

func TestTablePrinter_Empty(t *testing.T) {
	t.Parallel()

	out := bytes.Buffer{}
	printer := template.NewTablePrinter(&out, false, true).
		Title("Example").
		SetHeaders("Wow", "Such", "A", "Fancy", "Printer")

	printer.Render()
	require.Empty(t, out.String())
}

func TestTablePrinter(t *testing.T) {
	t.Parallel()

	out := bytes.Buffer{}
	printer := template.NewTablePrinter(&out, false, true).
		Title("Example").
		SetHeaders("Wow", "Such", "A", "Fancy", "Printer")

	printer.Append([]string{"foo", "bar", "baz", "qux", "wat"})
	printer.Append([]string{"veryveryverylongggg", "", "a", "b", "c"})

	printer.Render()
	expected := `EXAMPLE
+---------------------+------+-----+-------+---------+
|         WOW         | SUCH |  A  | FANCY | PRINTER |
+---------------------+------+-----+-------+---------+
|         foo         | bar  | baz |  qux  |   wat   |
+---------------------+------+-----+-------+---------+
| veryveryverylongggg |      |  a  |   b   |    c    |
+---------------------+------+-----+-------+---------+
|                                    TOTAL |    2    |
+---------------------+------+-----+-------+---------+
`
	require.Equal(t, expected, out.String())
}

func TestTablePrinter_Description(t *testing.T) {
	t.Parallel()

	out := bytes.Buffer{}
	printer := template.NewTablePrinter(&out, false, true).
		Title("Example").
		SetHeaders("Wow", "Such", "A", "Fancy", "Description")
	printer.Append([]string{"once", "upon", "a", "time", "short description"})

	printer.Render()
	// Expect that the description is left-aligned with a min width.
	expected := `EXAMPLE
+------+------+---+-------+--------------------------------+
| WOW  | SUCH | A | FANCY |          DESCRIPTION           |
+------+------+---+-------+--------------------------------+
| once | upon | a | time  | short description              |
+------+------+---+-------+--------------------------------+
|                   TOTAL |               1                |
+------+------+---+-------+--------------------------------+
`
	require.Equal(t, expected, out.String())
}
