package v1repl

import (
	"fmt"
	"math"
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
	visibleCols []table.Column
	allCols     []table.Column
	colWidths   []int
	colNames    []string
	colOffset   int
	simpleTable table.Model
	totalMargin int
	totalWidth  int
	rowsPerPage int
	currentPage int
	pageCount   int
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
		rows:        rows,
		visibleCols: cols,
		allCols:     cols,
		colNames:    colNames,
		colWidths:   colWidths,
		rowsPerPage: 15,
		pageCount:   int(math.Ceil(float64(len(rows)) / float64(15))),
		currentPage: 0,
	}
	m.regeneratePage()
	return m
}

func (m *Model) regeneratePage() {
	pageEnd := (m.currentPage + 1) * m.rowsPerPage
	if len(m.rows) < pageEnd {
		pageEnd = len(m.rows)
	}
	targetWidth := m.totalWidth - m.totalMargin
	visibleWidth := m.getVisibleTableWidth()
	if targetWidth > visibleWidth {
		targetWidth = visibleWidth
	}
	m.visibleCols = append([]table.Column{m.allCols[0]},
		m.allCols[m.colOffset+1:]...)
	m.simpleTable = table.New(m.visibleCols).
		WithRows(m.rows[m.currentPage*m.rowsPerPage : pageEnd]).
		WithStaticFooter(
			fmt.Sprintf("%d Columns, %d Rows, Page %d/%d",
				len(m.visibleCols), len(m.rows), m.currentPage+1, m.pageCount)).
		WithTargetWidth(targetWidth)
}

func (m Model) getVisibleTableWidth() int {
	width := 2 + m.colWidths[0]
	for _, colWidth := range m.colWidths[1+m.colOffset:] {
		width += colWidth + 1
	}
	return width
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

	if m.pageCount == 1 {
		fmt.Printf("\n")
		cmds = append(cmds, tea.Quit)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+d", "esc", "q":
			fmt.Printf("\n")
			cmds = append(cmds, tea.Quit)

		case "left":
			if m.colOffset > 0 {
				m.colOffset--
				m.regeneratePage()
			}

		case "right":
			if m.colOffset < len(m.allCols)-2 {
				m.colOffset++
				m.regeneratePage()
			}
		case "down":
			if m.currentPage < m.pageCount-1 {
				m.currentPage++
				m.regeneratePage()
			}
		case "up":
			if m.currentPage > 0 {
				m.currentPage--
				m.regeneratePage()
			}
		}
	case tea.WindowSizeMsg:
		m.totalWidth = msg.Width
		m.rowsPerPage = msg.Height - 7
		m.pageCount = int(math.Ceil(float64(len(m.rows)) / float64(m.rowsPerPage)))
		m.regeneratePage()
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString(m.simpleTable.View())

	return body.String()
}
