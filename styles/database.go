package styles

import "github.com/charmbracelet/lipgloss"

var (
	HeaderDatabaseStyle = lipgloss.NewStyle().
				Background(PrimaryColor).
				Foreground(BackgroundColor).
				Bold(true).
				Align(lipgloss.Left)

	MidBlock          = lipgloss.NewStyle()
	DatabaseLeftPanel = lipgloss.NewStyle().
				Background(BackgroundColor).
				Foreground(SecondaryColor).
				Border(lipgloss.NormalBorder(), false, true, false, false).
				BorderForeground(SecondaryColor).Align(lipgloss.Center)
	SearchBox = lipgloss.NewStyle().
			Background(BackgroundColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(PrimaryColor)
	DatabaseList = lipgloss.NewStyle().
			Background(BackgroundColor).
			Foreground(PrimaryColor).
			AlignHorizontal(lipgloss.Left)
)
