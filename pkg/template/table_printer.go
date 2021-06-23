package template

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type TablePrinter struct {
	w      io.Writer
	writer *tablewriter.Table

	useColor bool
	title    string

	headerLen   int
	appendCalls int
}

func NewTablePrinter(w io.Writer, hasColor, hasBorder bool) *TablePrinter {
	wr := tablewriter.NewWriter(w)
	wr.SetBorder(hasBorder)
	wr.SetRowLine(hasBorder)

	return &TablePrinter{
		w:        w,
		writer:   wr,
		useColor: hasColor,
	}
}

func (t *TablePrinter) Render() {
	if t.appendCalls == 0 {
		return
	}

	title := strings.ToUpper(t.title)
	if t.useColor {
		title = colorTitle.Sprint(title)
	}
	fmt.Fprintln(t.w, title)

	t.setFooter()
	t.writer.Render()
}

func (t *TablePrinter) Title(title string) *TablePrinter {
	t.title = title
	return t
}

func (t *TablePrinter) SetHeaders(headers ...string) *TablePrinter {
	t.headerLen = len(headers)
	t.writer.SetHeader(headers)

	headerColors := make([]tablewriter.Colors, t.headerLen)
	alignments := make([]int, t.headerLen)

	color := noColor
	if t.useColor {
		color = colorHeader
	}
	for i, header := range headers {
		headerColors[i] = color
		if strings.EqualFold("description", header) {
			t.writer.SetColMinWidth(i, 30)
			alignments[i] = tablewriter.ALIGN_LEFT
		} else {
			alignments[i] = tablewriter.ALIGN_CENTER
		}
	}
	t.writer.SetHeaderColor(headerColors...)
	t.writer.SetColumnAlignment(alignments)

	return t
}

func (t *TablePrinter) setFooter() *TablePrinter {
	footers := make([]string, t.headerLen)
	if t.headerLen > 1 {
		footers[len(footers)-2] = "TOTAL"
		footers[len(footers)-1] = strconv.Itoa(t.appendCalls)
	} else {
		footers[0] = "TOTAL: " + strconv.Itoa(t.appendCalls)
	}
	t.writer.SetFooter(footers)

	colors := make([]tablewriter.Colors, t.headerLen)
	color := noColor
	if t.useColor {
		color = colorFooter
	}
	if t.headerLen > 1 {
		colors[len(colors)-2] = color
		colors[len(colors)-1] = color
	} else {
		colors[0] = color
	}
	t.writer.SetFooterColor(colors...)

	return t
}

func (t *TablePrinter) Append(slc []string) {
	t.appendCalls++
	t.writer.Append(slc)
}
