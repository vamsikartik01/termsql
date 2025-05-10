package main

import (
	"fmt"
	"termsql/db"
	"termsql/types"
	"termsql/ui"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	inputStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	outputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Italic(true)
	headerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("111")).Bold(true).Underline(true)
)

type model struct {
	input  textinput.Model
	output string
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.output = m.input.Value()
			m.input.Reset()

		case "q":
			return m, tea.Quit
		default:
			m.input, cmd = m.input.Update(msg)
		}
	default:
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m *model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s",
		headerStyle.Render("Type something below:"),
		inputStyle.Render(m.input.View()),
		headerStyle.Render("Output:"),
		outputStyle.Render(m.output),
	)
}
func main() {

	sconnections := []types.Connection{}
	sconnections = append(
		sconnections,
		types.Connection{Name: "Renderer", Host: "127.0.0.1", Port: "3306", Username: "root", Password: "zino"},
	)
	sconnections = append(
		sconnections,
		types.Connection{Name: "Builder", Host: "127.0.0.1", Port: "3306", Username: "root", Password: "zino"},
	)
	sconnections = append(
		sconnections,
		types.Connection{Name: "Development", Host: "127.0.0.1", Port: "3306", Username: "root", Password: "zino"},
	)
	sconnections = append(
		sconnections,
		types.Connection{Name: "Sandbox", Host: "127.0.0.1", Port: "3306", Username: "root", Password: "zino"},
	)

	var conn types.Connection
	// conn = ui.RunConnectionForm(sconnections)
	// fmt.Println(conn)

	conn.Host = "127.0.0.1"
	conn.Port = "3306"
	conn.Username = "root"
	conn.Password = "zino"
	// bootupBubble := ui.NewLoadingModel()
	// b := tea.NewProgram(bootupBubble, tea.WithAltScreen())
	// go func() {
	// 	_, err := b.Run()
	// 	if err != nil {
	// 		fmt.Println("Error bootup", err)
	// 	}
	// }()

	dbConn, err := db.Init(conn)
	if err != nil {
		fmt.Println("Error initiating mysql conn", err)
	}
	defer dbConn.Close()
	// b.Kill()

	db := ui.RunDatabaseSelector(dbConn, conn)
	fmt.Println(db)

}
