package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type bootupModel struct {
	spinner spinner.Model
	message string
}

func (m bootupModel) Init() tea.Cmd {
	return spinner.Tick
}

func (m bootupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m bootupModel) View() string {
	return fmt.Sprintf("\n\n  %s %s\n\n", m.spinner.View(), m.message)
}

func NewLoadingModel() bootupModel {
	s := spinner.New()
	s.Spinner = spinner.Points // Can customize it
	return bootupModel{
		spinner: s,
		message: "Connecting to server...",
	}
}

func NewBootupBubble() {
	p := tea.NewProgram(NewLoadingModel(), tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
