package template

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type DiffPrinter struct {
	w      io.Writer
	writer *tablewriter.Table

	useColor bool
	title    string

	appendCalls int
	headerLen   int
}

func NewDiffPrinter(w io.Writer, hasColor, hasBorder bool) *DiffPrinter {
	wr := tablewriter.NewWriter(w)
	wr.SetBorder(hasBorder)
	wr.SetRowLine(hasBorder)

	return &DiffPrinter{
		w:        w,
		writer:   wr,
		useColor: hasColor,
	}
}

func (d *DiffPrinter) Render() {
	if d.appendCalls == 0 {
		return
	}

	// set the title and the add/remove legend
	title := strings.ToUpper(d.title)
	add := "+add"
	remove := "-remove"
	if d.useColor {
		title = colorTitle.Sprint(title)
		add = colorTotalAdd.Sprint(add)
		remove = colorTotalRemove.Sprint(remove)
	}
	fmt.Fprintf(d.w, "%s    %s | %s | unchanged\n", title, add, remove)

	d.setFooter()
	d.writer.Render()
}

func (d *DiffPrinter) Title(title string) *DiffPrinter {
	d.title = title
	return d
}

func (d *DiffPrinter) SetHeaders(headers ...string) *DiffPrinter {
	headers = d.prepend(headers, "+/-")
	d.headerLen = len(headers)

	d.writer.SetHeader(headers)

	headerColors := make([]tablewriter.Colors, d.headerLen)
	color := noColor
	if d.useColor {
		color = colorHeader
	}
	for i := range headerColors {
		headerColors[i] = color
	}
	d.writer.SetHeaderColor(headerColors...)

	return d
}

func (d *DiffPrinter) setFooter() *DiffPrinter {
	footers := make([]string, d.headerLen)
	if d.headerLen > 1 {
		footers[len(footers)-2] = "TOTAL"
		footers[len(footers)-1] = strconv.Itoa(d.appendCalls)
	} else {
		footers[0] = "TOTAL: " + strconv.Itoa(d.appendCalls)
	}

	d.writer.SetFooter(footers)
	colors := make([]tablewriter.Colors, d.headerLen)
	color := noColor
	if d.useColor {
		color = colorFooter
	}
	if d.headerLen > 1 {
		colors[len(colors)-2] = color
		colors[len(colors)-1] = color
	} else {
		colors[0] = color
	}
	d.writer.SetFooterColor(colors...)

	return d
}

func (d *DiffPrinter) Append(slc []string) {
	d.writer.Append(d.prepend(slc, ""))
}

func (d *DiffPrinter) AppendDiff(remove, add []string) {
	defer func() { d.appendCalls++ }()

	if d.appendCalls > 0 {
		d.appendBufferLine()
	}

	lenAdd, lenRemove := len(add), len(remove)
	preppedAdd, preppedRemove := d.prepend(add, "+"), d.prepend(remove, "-")
	if lenRemove > 0 && lenAdd == 0 {
		d.writer.Rich(preppedRemove, d.redRow(len(preppedRemove)))
		return
	}
	if lenAdd > 0 && lenRemove == 0 {
		d.writer.Rich(preppedAdd, d.greenRow(len(preppedAdd)))
		return
	}

	var (
		addColors    = make([]tablewriter.Colors, len(preppedAdd))
		removeColors = make([]tablewriter.Colors, len(preppedRemove))
		hasDiff      bool
	)
	addColor, removeColor := noColor, noColor
	if d.useColor {
		addColor, removeColor = colorAdd, colorRemove
	}
	for i := 0; i < lenRemove; i++ {
		if add[i] != remove[i] {
			hasDiff = true
			// offset to skip prepended +/- column
			addColors[i+1], removeColors[i+1] = addColor, removeColor
		}
	}

	if !hasDiff {
		d.writer.Append(d.prepend(add, ""))
		return
	}

	addColors[0], removeColors[0] = addColor, removeColor
	d.writer.Rich(d.prepend(remove, "-"), removeColors)
	d.writer.Rich(d.prepend(add, "+"), addColors)
}

func (d *DiffPrinter) appendBufferLine() {
	d.writer.Append([]string{})
}

func (d *DiffPrinter) redRow(i int) []tablewriter.Colors {
	return d.colorRow(colorRemove, i)
}

func (d *DiffPrinter) greenRow(i int) []tablewriter.Colors {
	return d.colorRow(colorAdd, i)
}

func (d *DiffPrinter) prepend(slc []string, val string) []string {
	return append([]string{val}, slc...)
}

func (d *DiffPrinter) colorRow(color tablewriter.Colors, i int) []tablewriter.Colors {
	colors := make([]tablewriter.Colors, i)
	for i := range colors {
		if d.useColor {
			colors[i] = color
		} else {
			colors[i] = noColor
		}
	}
	return colors
}
