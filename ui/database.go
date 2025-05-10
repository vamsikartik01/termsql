package ui

import (
	"fmt"
	"strings"
	"termsql/constants"
	"termsql/db"
	"termsql/styles"
	"termsql/types"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DatabaseConnectionModel struct {
	conn types.Connection
	db   db.SQL

	allDatabases      []string
	filteredDatabases []string

	allTables      []string
	filteredTables []string

	textInput  textinput.Model
	searchText string
	cursor     int
	page       int
	pageSize   int
	selectDB   string

	windowWidth  int
	windowHeight int

	Help         []types.Help
	databaseMode bool

	// dbListViewport viewport.Model
}

func (m *DatabaseConnectionModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *DatabaseConnectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "esc":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.databaseMode {
				if m.cursor+1 < len(m.filteredDatabases) {
					m.cursor++
				}
			} else {
				if m.cursor+1 < len(m.filteredTables) {
					m.cursor++
				}
			}

		case "enter":
			if m.databaseMode {
				index := m.cursor
				if index < len(m.filteredDatabases) {
					m.selectDB = m.filteredDatabases[index]
					m.databaseMode = false
					m.fetchTables()
					m.filterTables()
				}
			} else {

			}

		}
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	curr := m.textInput.Value()
	if curr != m.searchText {
		if m.databaseMode {
			m.filterDatabases()
		} else {
			m.filterTables()
		}

	}

	return m, cmd
}

func (m *DatabaseConnectionModel) filterDatabases() {
	query := strings.ToLower(m.textInput.Value())
	var filtered []string
	for _, dbName := range m.allDatabases {
		if strings.Contains(strings.ToLower(dbName), query) {
			filtered = append(filtered, dbName)
		}
	}
	m.filteredDatabases = filtered
	m.cursor = 0
}

func (m *DatabaseConnectionModel) fetchTables() {

	err := m.db.SwitchDatabase(m.selectDB)
	if err != nil {

	}
	tables, err := m.db.GetTables()
	if err != nil {

	}
	m.allTables = tables
}

func (m *DatabaseConnectionModel) filterTables() {
	query := strings.ToLower(m.textInput.Value())
	var filtered []string
	if query == "" {
		m.filteredTables = m.allTables
	} else {
		for _, tableName := range m.allTables {
			if strings.Contains(strings.ToLower(tableName), query) {
				filtered = append(filtered, tableName)
			}
		}
		m.filteredTables = filtered
	}

	m.cursor = 0
}

func (m DatabaseConnectionModel) View() string {

	bottomBarHeight := 1
	headerHeight := 1
	midBlockHeight := m.windowHeight - bottomBarHeight - headerHeight
	leftPanelWidth := constants.DefaultDatabaseLeftPanelWidth
	rightPanelWidth := m.windowWidth - leftPanelWidth

	// Header
	headerBar := styles.HeaderDatabaseStyle.Width(m.windowWidth).Height(headerHeight).Render(
		m.selectDB + " | " + m.conn.Host + " | " + m.conn.Port + " | " + m.conn.Username,
	)

	// Mid Block
	midBlock := ""

	//mid left panel
	if m.databaseMode {
		m.textInput.Width = leftPanelWidth - 5

		searchBox := styles.SearchBox.Render(m.textInput.View())

		databases := []string{}
		for i, database := range m.filteredDatabases {
			cursor := "   "
			if i == m.cursor {
				cursor = "👉 "
			}
			databaseName := database
			databases = append(databases, cursor+databaseName)
		}

		list := styles.DatabaseList.Width(leftPanelWidth).Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				databases...,
			),
		)

		leftPanel := styles.DatabaseLeftPanel.Height(midBlockHeight).Width(leftPanelWidth).Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				searchBox,
				list,
			),
		)
		RightPane := styles.HeaderStyle.Height(midBlockHeight).
			Width(rightPanelWidth).
			AlignVertical(lipgloss.Center).
			Render("Select Database")

		midBlock = styles.MidBlock.Height(midBlockHeight).Width(m.windowWidth).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				leftPanel,
				RightPane,
			),
		)
	} else {
		m.textInput.Width = leftPanelWidth - 5

		searchBox := styles.SearchBox.Render(m.textInput.View())

		tables := []string{}
		for i, table := range m.filteredTables {
			cursor := "   "
			if i == m.cursor {
				cursor = "👉 "
			}
			tables = append(tables, cursor+table)
		}

		list := styles.DatabaseList.Width(leftPanelWidth).Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				tables...,
			),
		)

		leftPanel := styles.DatabaseLeftPanel.Height(midBlockHeight).Width(leftPanelWidth).Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				searchBox,
				list,
			),
		)
		RightPane := styles.HeaderStyle.Height(midBlockHeight).
			Width(rightPanelWidth).
			AlignVertical(lipgloss.Center).
			Render("Select Database")

		midBlock = styles.MidBlock.Height(midBlockHeight).Width(m.windowWidth).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				leftPanel,
				RightPane,
			),
		)
	}

	// Bottom help bar
	helpTexts := []string{}
	for _, helpText := range m.Help {
		helpTexts = append(helpTexts, styles.BottomBarItemStyle.Render(
			styles.BottomBarItemTextStyle.Render(
				" "+helpText.Key+" ",
			)+styles.BottomBarItemHeaderStyle.Render(
				" "+helpText.Title+" ",
			),
		))
	}
	bottomBar := styles.BootomBarStyle.Height(bottomBarHeight).Width(m.windowWidth).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			helpTexts...,
		),
	)

	// Page UI
	ui := styles.Screen.Height(m.windowHeight).Width(m.windowWidth).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			headerBar,
			midBlock,
			bottomBar,
		),
	)

	return ui
}

func NewDatabaseSelectorModel(dbConn db.SQL, conn types.Connection) DatabaseConnectionModel {
	ti := textinput.New()
	ti.Placeholder = "Search database..."
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30

	return DatabaseConnectionModel{
		conn:      conn,
		db:        dbConn,
		textInput: ti,
		cursor:    0,
		page:      0,
		pageSize:  10,

		databaseMode: true,
		selectDB:     "no database",
		Help:         constants.DefaultDatabaseHelp,
	}
}

func RunDatabaseSelector(dbConn db.SQL, conn types.Connection) string {
	model := NewDatabaseSelectorModel(dbConn, conn)

	// Fetch databases before starting UI
	databases, err := dbConn.ListDatabases("")
	if err != nil {
		fmt.Println("Error fetching databases:", err)
		return ""
	}
	model.allDatabases = databases
	model.filteredDatabases = databases

	p := tea.NewProgram(&model, tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		return ""
	}

	finalModel := m.(*DatabaseConnectionModel)
	return finalModel.selectDB
}
