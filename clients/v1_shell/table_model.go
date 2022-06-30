package v1shell

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

type EndingStatus uint

const (
	quitStatus EndingStatus = iota
	goToPrevTableStatus
	goToPrevTableJumpFirstPageStatus
	goToNextTableStatus
)

type Model struct {
	name               string
	tags               map[string]string
	curResult          int
	resultMax          int
	curSeries          int
	seriesMax          int
	rows               []table.Row
	allCols            []table.Column
	colNames           []string
	simpleTable        table.Model
	totalWidth         int
	endingStatus       EndingStatus
	tableScreenPadding int
}

func NewModel(
	res api.InfluxqlJsonResponseSeries,
	jumpToFirstPage bool,
	name string,
	tags map[string]string,
	curRes int,
	resMax int,
	curSer int,
	serMax int,
	scientific bool) Model {

	cols := make([]table.Column, len(*res.Columns)+1)
	colWidths := make([]int, len(*res.Columns)+1)
	alignment := make([]lipgloss.Position, len(*res.Columns)+1)
	rows := make([]table.Row, len(*res.Values))
	colNames := *res.Columns
	for rowI, row := range *res.Values {
		rd := table.RowData{}
		rd["index"] = fmt.Sprintf("%d", rowI+1)
		alignment[0] = lipgloss.Right
		colWidths[0] = len("index") + colPadding
		for colI, rowVal := range row {
			var item string
			var colLen int
			switch val := rowVal.(type) {
			case int:
				item = fmt.Sprintf("%d", val)
				colLen = len(item)
				alignment[colI+1] = lipgloss.Right
			case string:
				item = color.YellowString(val)
				colLen = len(val)
				alignment[colI+1] = lipgloss.Left
			case float32, float64:
				if scientific {
					item = fmt.Sprintf("%.10e", val)
				} else {
					item = fmt.Sprintf("%.10f", val)
				}
				colLen = len(item)
				alignment[colI+1] = lipgloss.Right
			default:
				item = fmt.Sprintf("%v", val)
				colLen = len(item)
				alignment[colI+1] = lipgloss.Right
			}
			rd[colNames[colI]] = item
			if colWidths[colI+1] < colLen+colPadding {
				colWidths[colI+1] = colLen + colPadding
			}
		}
		rows[rowI] = table.NewRow(rd).
			WithStyle(lipgloss.NewStyle())
	}
	cols[0] = table.NewColumn("index", "index", colWidths[0])
	indexStyle := lipgloss.NewStyle()
	if isatty.IsTerminal(os.Stdout.Fd()) {
		indexStyle = indexStyle.
			Faint(true).
			Align(lipgloss.Right)
	}
	cols[0] = cols[0].WithStyle(indexStyle)
	for colI, colTitle := range colNames {
		if colWidths[colI+1] < len(colTitle)+colPadding {
			colWidths[colI+1] = len(colTitle) + colPadding
		}
		cols[colI+1] = table.NewColumn(colTitle, color.HiCyanString(colTitle), colWidths[colI+1]).
			WithStyle(lipgloss.NewStyle().Align(alignment[colI+1]))
	}
	colNames = append([]string{"index"}, colNames...)
	screenPadding := 10
	if len(tags) > 0 {
		screenPadding++
	}
	m := Model{
		name:               name,
		tags:               tags,
		curResult:          curRes,
		resultMax:          resMax,
		curSeries:          curSer,
		seriesMax:          serMax,
		rows:               rows,
		allCols:            cols,
		colNames:           colNames,
		tableScreenPadding: screenPadding,
	}
	keybind := table.DefaultKeyMap()
	keybind.RowUp.SetEnabled(false)
	keybind.RowDown.SetEnabled(false)
	keybind.PageUp.SetEnabled(false)
	keybind.PageDown.SetEnabled(false)
	keybind.ScrollLeft.SetKeys("left")
	keybind.ScrollRight.SetKeys("right")
	keybind.Filter.Unbind()
	keybind.FilterBlur.Unbind()
	keybind.FilterClear.Unbind()

	m.simpleTable = table.New(m.allCols).
		HeaderStyle(lipgloss.NewStyle().Align(lipgloss.Center)).
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

	if jumpToFirstPage {
		m.simpleTable = m.simpleTable.PageLast()
	}
	return m
}

func (m Model) Init() tea.Cmd {
	color.Magenta("Interactive Table View (press q to exit mode, shift+up/down to navigate tables):")
	builder := strings.Builder{}
	if m.name != "" {
		fmt.Printf("Name: %s\n", color.GreenString(m.name))
	} else {
		fmt.Println("") // keep a consistent height, so print an empty line
	}
	if len(m.tags) > 0 {
		fmt.Print("Tags: ")
		for key, val := range m.tags {
			if key == "" || val == "" {
				continue
			}
			builder.WriteString(fmt.Sprintf("%s=%s, ", color.YellowString(key), color.CyanString(val)))
		}
		tagline := builder.String()
		fmt.Print(tagline[:len(tagline)-2])
		fmt.Println("")
	}
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
			m.simpleTable = m.simpleTable.Focused(false)
			fmt.Printf("\n")
			m.endingStatus = quitStatus
			cmds = append(cmds, tea.Quit)
		case "shift+up":
			if !(m.curResult == 1 && m.curSeries == 1) {
				m.endingStatus = goToPrevTableJumpFirstPageStatus
				cmds = append(cmds, tea.Quit)
			}
		case "shift+down":
			if !(m.curResult == m.resultMax && m.curSeries == m.seriesMax) {
				m.endingStatus = goToNextTableStatus
				cmds = append(cmds, tea.Quit)
			}
		case "up":
			if m.simpleTable.CurrentPage() == 1 {
				if !(m.curResult == 1 && m.curSeries == 1) {
					m.endingStatus = goToPrevTableStatus
					cmds = append(cmds, tea.Quit)
				}
			} else {
				m.simpleTable = m.simpleTable.PageUp()
			}
		case "down":
			if m.simpleTable.CurrentPage() == m.simpleTable.MaxPages() {
				if !(m.curResult == m.resultMax && m.curSeries == m.seriesMax) {
					m.endingStatus = goToNextTableStatus
					cmds = append(cmds, tea.Quit)
				}
			} else {
				m.simpleTable = m.simpleTable.PageDown()
			}
		case "[":
			m.simpleTable = m.simpleTable.PageFirst()
		case "]":
			m.simpleTable = m.simpleTable.PageLast()
		}
	case tea.WindowSizeMsg:
		m.totalWidth = msg.Width
		m.simpleTable = m.simpleTable.WithPageSize(msg.Height - m.tableScreenPadding).
			WithMaxTotalWidth(msg.Width)
	}
	m.simpleTable = m.simpleTable.WithStaticFooter(
		fmt.Sprintf("%d Columns, %d Rows, Page %d/%d\nTable %d/%d, Statement %d/%d",
			len(m.allCols), len(m.rows), m.simpleTable.CurrentPage(), m.simpleTable.MaxPages(),
			m.curSeries, m.seriesMax,
			m.curResult, m.resultMax))
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString(m.simpleTable.View())
	body.WriteString("\n")
	return body.String()
}
