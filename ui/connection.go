package ui

import (
	"fmt"
	"os"
	"termsql/styles"
	"termsql/types"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	host             textinput.Model
	port             textinput.Model
	username         textinput.Model
	password         textinput.Model
	focusIndex       int
	panelIndex       int
	savedConnIndex   int
	savedConnections []types.Connection

	windowWidth  int
	windowHeight int
}

func InitialModel() model {
	hostTi := textinput.New()
	hostTi.Placeholder = "127.0.0.1"
	hostTi.Prompt = "Host : "

	portTI := textinput.New()
	portTI.Placeholder = "3306"
	portTI.Prompt = "Port : "

	UsernameTi := textinput.New()
	UsernameTi.Placeholder = "username"
	UsernameTi.Prompt = "Username : "

	PwasswordTi := textinput.New()
	PwasswordTi.Placeholder = "password"
	PwasswordTi.Prompt = "Password : "

	return model{
		host:       hostTi,
		port:       portTI,
		username:   UsernameTi,
		password:   PwasswordTi,
		focusIndex: 0,
		panelIndex: 0,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	total := len(m.savedConnections)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			fmt.Println("Exiting...")
			return m, tea.Quit
		case "left":
			m.panelIndex = 0
			m.password.Blur()
			m.port.Blur()
			m.username.Blur()
			m.host.Blur()
		case "right":
			m.panelIndex = 1
		case "up":
			if m.panelIndex == 0 {
				m.savedConnIndex--
				if m.savedConnIndex < 0 {
					m.savedConnIndex = 0
				}
			} else {
				m.focusIndex--
				if m.focusIndex < 0 {
					m.focusIndex = 0
				}
			}

		case "down":
			if m.panelIndex == 0 {
				m.savedConnIndex++
				if m.savedConnIndex > total-1 {
					m.savedConnIndex = 0
				}
			} else {
				m.focusIndex++
				if m.focusIndex > 3 {
					m.focusIndex = 3
				}
			}

		case "tab":
			if m.panelIndex == 1 {
				m.focusIndex++
			}

		case "enter":
			if m.panelIndex == 1 {
				if m.focusIndex == 3 {
					return m, tea.Quit
				}
				m.focusIndex++
			} else {
				return m, tea.Quit
			}

		}

		if m.panelIndex == 1 {
			switch m.focusIndex {
			case 0:
				m.host.Focus()
				m.port.Blur()
				m.username.Blur()
				m.password.Blur()
			case 1:
				m.port.Focus()
				m.host.Blur()
				m.username.Blur()
				m.password.Blur()
			case 2:
				m.username.Focus()
				m.port.Blur()
				m.host.Blur()
				m.password.Blur()
			case 3:
				m.password.Focus()
				m.port.Blur()
				m.username.Blur()
				m.host.Blur()
			}

		}

	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		return m, nil
	}

	var cmd tea.Cmd
	m.host, cmd = m.host.Update(msg)
	m.port, cmd = m.port.Update(msg)
	m.username, cmd = m.username.Update(msg)
	m.password, cmd = m.password.Update(msg)

	return m, cmd
}

func (m model) View() string {
	bottomBarHeight := 1
	topBarHeight := m.windowHeight - bottomBarHeight
	leftPanelWidth := m.windowWidth / 4
	rightPanleWidth := m.windowWidth - leftPanelWidth

	// Bottom help bar
	bottomBar := styles.BootomBarStyle.Height(bottomBarHeight).Width(m.windowWidth).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			styles.BottomBarItemStyle.Render(
				styles.BottomBarItemTextStyle.Render("Ctrl-C ")+styles.BottomBarItemHeaderStyle.Render(" Exit "),
			),
			styles.BottomBarItemStyle.Render(
				styles.BottomBarItemTextStyle.Render(" Ctrl-S ")+styles.BottomBarItemHeaderStyle.Render(" Settings "),
			),
		),
	)

	// Left Panel
	WelcomeHeader := styles.WelcomeHeaderStyle.Width(leftPanelWidth).Height(3).Render("WELCOME MR.KARTIK")

	SavedConnHeader := styles.SavedConnectionHeaderStle.Width(leftPanelWidth).Render("Saved Connections")

	var savedConnItems []string
	for i, sconn := range m.savedConnections {
		if i == m.savedConnIndex && m.panelIndex == 0 {
			savedConnItems = append(
				savedConnItems,
				styles.SavedConnectionItemSelectedStyle.Width(leftPanelWidth-1).
					Render(sconn.Name+"\n"+sconn.Host+", "+sconn.Port+", "+sconn.Username),
			)
		} else {
			savedConnItems = append(
				savedConnItems,
				styles.SavedConnectionItemStyle.Width(leftPanelWidth-1).
					Render(sconn.Name+"\n"+sconn.Host+", "+sconn.Port+", "+sconn.Username),
			)
		}
	}
	SavedConnections := styles.SavedConnectionStyle.Width(leftPanelWidth).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			savedConnItems...,
		),
	)

	leftPanel := styles.LeftPanelHomeStyle.Height(topBarHeight).Width(leftPanelWidth).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			WelcomeHeader,
			SavedConnHeader,
			SavedConnections,
		),
	)

	TitleBlock := styles.TitleBlockStyle.Height(topBarHeight / 3).Width(rightPanleWidth).Render("TermSQL 🚀")

	formContent := lipgloss.JoinVertical(lipgloss.Left,
		styles.InputStyle.Render(m.host.View()),
		styles.InputStyle.Render(m.port.View()),
		styles.InputStyle.Render(m.username.View()),
		styles.InputStyle.Render(m.password.View()),
	)
	Form := styles.FormStyle.UnsetBorderForeground().Width(rightPanleWidth / 2).Render(formContent)

	rightPanel := styles.RightPanelHomeStyle.Height(topBarHeight).Width(rightPanleWidth - 2).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			TitleBlock,
			Form,
			lipgloss.NewStyle().Render("\n(Press Enter to Connect)"),
		),
	)

	topBar := styles.TopbarStyle.Height(topBarHeight).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			leftPanel,
			rightPanel,
		),
	)

	ui := styles.Screen.Height(m.windowHeight).Width(m.windowWidth).Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			topBar,
			bottomBar,
		),
	)

	return ui
}

func (m model) ViewLegacy() string {
	leftContent := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.TitleBox.Width(m.windowWidth/4).Height(m.windowHeight/10).Render("WELCOME MR.KARTIK"),
		styles.ConnectionHistoryBox.Width(m.windowWidth/5).
			Height(m.windowHeight*4/5).
			Render("Saved connections"), // your logic to show past names
	)

	// Calculate heights
	welcomeHeight := m.windowHeight / 3

	welcomeBanner := styles.WelcomeBanner.
		Height(welcomeHeight).
		Width((m.windowWidth * 3) / 4).
		Render("\n\nTermSQL 🚀\n")

	formContent := lipgloss.JoinVertical(lipgloss.Left,
		styles.InputStyle.Render(m.host.View()),
		styles.InputStyle.Render(m.port.View()),
		styles.InputStyle.Render(m.username.View()),
		styles.InputStyle.Render(m.password.View()),
	)

	// Fancy "★ Favourite" button
	favButton := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")). // Gold-ish
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.AccentColor).
		Padding(0, 2).
		Align(lipgloss.Center).
		Width(20).
		Render("★ Favourite")

	rightContent := lipgloss.JoinVertical(lipgloss.Left,
		welcomeBanner,
		formContent,
		"press Enter to Connect\n",
		"\n"+favButton,
	)

	ui := lipgloss.JoinHorizontal(lipgloss.Top,
		styles.LeftPanel.Height(m.windowHeight).Width(m.windowWidth/4).Render(leftContent),
		styles.RightPanel.Height(m.windowHeight).Width(m.windowWidth*3/4).Render(rightContent),
	)

	return ui
}

func RunConnectionForm(savedConns []types.Connection) types.Connection {
	i := InitialModel()
	i.savedConnections = savedConns
	p := tea.NewProgram(i, tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	finalModel, ok := m.(model)

	var conn types.Connection

	if ok {
		if finalModel.panelIndex == 1 {
			conn.Host = finalModel.host.Value()
			conn.Port = finalModel.port.Value()
			conn.Username = finalModel.username.Value()
			conn.Password = finalModel.password.Value()
		} else {
			newConn := finalModel.savedConnections[finalModel.savedConnIndex]
			conn.Host = newConn.Host
			conn.Port = newConn.Port
			conn.Username = newConn.Username
			conn.Password = newConn.Password
		}
	}

	return conn
}
