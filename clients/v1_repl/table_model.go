package v1repl

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/fatih/color"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/mattn/go-isatty"
)

const (
	colPadding int = 2
)

type Model struct {
	rows        []table.Row
	allCols     []table.Column
	colNames    []string
	simpleTable table.Model
	totalWidth  int
}

func NewModel(res api.InfluxqlJsonResponseSeries) Model {
	cols := make([]table.Column, len(*res.Columns)+1)
	colWidths := make([]int, len(*res.Columns)+1)
	rows := make([]table.Row, len(*res.Values))
	colNames := *res.Columns
	for rowI, row := range *res.Values {
		rd := table.RowData{}
		rd["index"] = fmt.Sprintf("%d", rowI+1)
		colWidths[0] = len("index") + colPadding
		for colI, rowVal := range row {
			var item string
			switch val := rowVal.(type) {
			case int:
				item = fmt.Sprintf("%d", val)
			case string:
				item = fmt.Sprintf("%q", val)
			default:
				item = fmt.Sprintf("%v", val)
			}
			rd[colNames[colI]] = item
			if colWidths[colI+1] < len(item)+colPadding {
				colWidths[colI+1] = len(item) + colPadding
			}
		}
		rows[rowI] = table.NewRow(rd).
			WithStyle(lipgloss.NewStyle().Align(lipgloss.Center))
	}
	cols[0] = table.NewColumn("index", "index", colWidths[0])
	indexStyle := lipgloss.NewStyle()
	if isatty.IsTerminal(os.Stdout.Fd()) {
		indexStyle = indexStyle.
			Faint(true).
			Align(lipgloss.Center)
	}
	cols[0] = cols[0].WithStyle(indexStyle)
	for colI, colTitle := range colNames {
		if colWidths[colI+1] < len(colTitle)+colPadding {
			colWidths[colI+1] = len(colTitle) + colPadding
		}
		cols[colI+1] = table.NewColumn(colTitle, color.HiCyanString(colTitle), colWidths[colI+1]).
			WithStyle(lipgloss.NewStyle().Align(lipgloss.Center))
	}
	colNames = append([]string{"index"}, colNames...)
	m := Model{
		rows:     rows,
		allCols:  cols,
		colNames: colNames,
	}
	keybind := table.DefaultKeyMap()
	keybind.RowUp.SetEnabled(false)
	keybind.RowDown.SetEnabled(false)
	keybind.PageDown.SetKeys("down")
	keybind.PageUp.SetKeys("up")
	keybind.ScrollLeft.SetKeys("left")
	keybind.ScrollRight.SetKeys("right")
	keybind.Filter.Unbind()
	keybind.FilterBlur.Unbind()
	keybind.FilterClear.Unbind()

	m.simpleTable = table.New(m.allCols).
		WithRows(m.rows).
		WithPageSize(15).
		WithMaxTotalWidth(500).
		WithHorizontalFreezeColumnCount(1).
		WithStaticFooter(
			fmt.Sprintf("%d Columns, %d Rows, Page %d/%d",
				len(m.allCols), len(m.rows), m.simpleTable.CurrentPage(), m.simpleTable.MaxPages())).
		HighlightStyle(lipgloss.NewStyle()).
		Focused(true).
		SelectableRows(false).
		WithKeyMap(keybind)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.simpleTable, cmd = m.simpleTable.Update(msg)
	cmds = append(cmds, cmd)

	if m.simpleTable.MaxPages() == 1 {
		m.simpleTable = m.simpleTable.Focused(false)
		cmds = append(cmds, tea.Quit)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+d", "esc", "q":
			m.simpleTable = m.simpleTable.Focused(false)
			fmt.Printf("\n")
			cmds = append(cmds, tea.Quit)
		}
	case tea.WindowSizeMsg:
		m.totalWidth = msg.Width
		m.simpleTable = m.simpleTable.WithPageSize(msg.Height - 7).
			WithMaxTotalWidth(msg.Width)
	}
	m.simpleTable = m.simpleTable.WithStaticFooter(
		fmt.Sprintf("%d Columns, %d Rows, Page %d/%d",
			len(m.allCols), len(m.rows), m.simpleTable.CurrentPage(), m.simpleTable.MaxPages()))
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString(m.simpleTable.View())
	if m.simpleTable.MaxPages() == 1 {
		body.WriteString("\n")
	}
	return body.String()
}
