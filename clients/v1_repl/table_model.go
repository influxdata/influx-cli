package v1repl

import (
	"fmt"
	"math"
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
	rows        []table.Row
	cols        []table.Column
	colNames    []string
	simpleTable table.Model
	totalMargin int
	totalWidth  int
	rowsPerPage int
	currentPage int
	numPages    int
}

func NewModel(res api.InfluxqlJsonResponseSeries) Model {
	cols := make([]table.Column, len(*res.Columns)+1)
	rows := make([]table.Row, len(*res.Values))
	colNames := *res.Columns
	for rowI, row := range *res.Values {
		rd := table.RowData{}
		rd["index"] = fmt.Sprintf("%d", rowI)
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
	cols[0] = table.NewColumn("index", "index", 10).WithStyle(lipgloss.NewStyle().
		Faint(true).
		Align(lipgloss.Center))
	for colI, colTitle := range colNames {
		cols[colI+1] = table.NewFlexColumn(colTitle, color.HiCyanString(colTitle), 1).
			WithStyle(lipgloss.NewStyle().Align(lipgloss.Center))
	}
	colNames = append([]string{"index"}, colNames...)
	m := Model{
		rows:        rows,
		cols:        cols,
		colNames:    colNames,
		rowsPerPage: 15,
		numPages:    int(math.Ceil(float64(len(rows)) / float64(15))),
		currentPage: 0,
	}
	m.regeneratePage()
	m.recalculateTable()
	return m
}

func (m *Model) regeneratePage() {
	pageEnd := (m.currentPage + 1) * m.rowsPerPage
	if len(m.rows) < pageEnd {
		pageEnd = len(m.rows)
	}
	m.simpleTable = table.New(m.cols).
		WithRows(m.rows[m.currentPage*m.rowsPerPage : pageEnd]).
		WithStaticFooter(
			fmt.Sprintf("%d Columns, %d Rows, Page %d/%d",
				len(m.cols), len(m.rows), m.currentPage+1, m.numPages))
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
			fmt.Printf("\n")
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
		case "down":
			if m.currentPage < m.numPages-1 {
				m.currentPage++
				m.regeneratePage()
				m.recalculateTable()
			}
		case "up":
			if m.currentPage > 0 {
				m.currentPage--
				m.regeneratePage()
				m.recalculateTable()
			}
		}
	case tea.WindowSizeMsg:
		m.totalWidth = msg.Width
		m.rowsPerPage = msg.Height - 7
		m.numPages = int(math.Ceil(float64(len(m.rows)) / float64(m.rowsPerPage)))
		m.regeneratePage()
		m.recalculateTable()
	}
	return m, tea.Batch(cmds...)
}

func (m *Model) recalculateTable() {
	m.simpleTable = m.simpleTable.WithTargetWidth(m.totalWidth - m.totalMargin)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString(m.simpleTable.View())

	return body.String()
}
