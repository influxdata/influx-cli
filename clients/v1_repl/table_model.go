package v1repl

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/fatih/color"
	"github.com/influxdata/influx-cli/v2/api"
)

const (
	minWidth = 15
)

type Model struct {
	simpleTable table.Model
	totalMargin int
	totalWidth  int
}

func NewModel(res api.InfluxqlJsonResponseSeries) Model {
	cols := make([]table.Column, len(*res.Columns))
	rows := make([]table.Row, len(*res.Values))
	colNames := *res.Columns
	for rowI, row := range *res.Values {
		rd := table.RowData{}
		for i, rowVal := range row {
			if rowValStr, ok := rowVal.(string); ok {
				rd[colNames[i]] = fmt.Sprintf("%q", rowValStr)
			} else {
				rd[colNames[i]] = fmt.Sprintf("%v", rowVal)
			}
		}
		rows[rowI] = table.NewRow(rd).
			WithStyle(lipgloss.NewStyle().Align(lipgloss.Center))
	}
	for colI, colTitle := range colNames {
		titleWidth := len(colTitle) + 2
		colWidth := 10
		if len(rows) > 0 {
			colWidth = len(fmt.Sprintf("%q", rows[0].Data[colTitle])) + 2
		}
		if colWidth < titleWidth {
			colWidth = titleWidth
		}
		cols[colI] = table.NewFlexColumn(colTitle, color.HiCyanString(colTitle), colWidth).
			WithStyle(lipgloss.NewStyle().Align(lipgloss.Center))
	}

	return Model{
		simpleTable: table.New(cols).WithRows(rows).WithStaticFooter(fmt.Sprintf("%d Columns, %d Rows", len(cols), len(rows))),
	}
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+d", "esc", "q":
			cmds = append(cmds, tea.Quit)

		case "left":
			if m.totalWidth-m.totalMargin > minWidth {
				m.totalMargin++
				m.recalculateTable()
			}

		case "right":
			if m.totalMargin > 0 {
				m.totalMargin--
				m.recalculateTable()
			}
		}
	case tea.WindowSizeMsg:
		m.totalWidth = msg.Width
		m.recalculateTable()
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) recalculateTable() {
	m.simpleTable = m.simpleTable.WithTargetWidth(m.totalWidth - m.totalMargin)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString("Query Response: (non-interactive)\nPress q or ctrl+c to quit\n\n")

	body.WriteString(m.simpleTable.View())

	return body.String() + "\n"
}
