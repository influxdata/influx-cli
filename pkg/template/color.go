package template

import (
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

var (
	colorTitle       = color.New(color.FgYellow, color.Bold)
	colorTotalAdd    = color.New(color.FgHiGreen, color.Bold)
	colorTotalRemove = color.New(color.FgRed, color.Bold)

	noColor     = tablewriter.Colors{}
	colorAdd    = tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold}
	colorFooter = tablewriter.Color(tablewriter.FgHiBlueColor, tablewriter.Bold)
	colorHeader = tablewriter.Colors{tablewriter.FgHiCyanColor, tablewriter.Bold}
	colorRemove = tablewriter.Colors{tablewriter.FgRedColor, tablewriter.Bold}
)
