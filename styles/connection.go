package styles

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor    = lipgloss.Color("#00FFFF") // Neon Cyan
	SecondaryColor  = lipgloss.Color("#FF00FF") // Electric Magenta
	AccentColor     = lipgloss.Color("#FF4500") // HUD Orange
	BackgroundColor = lipgloss.Color("#0F0F1B") // Deep space gray
	TextColor       = lipgloss.Color("#C0C0C0") // Soft metallic gray

	Screen = lipgloss.NewStyle().Background(BackgroundColor)

	TopbarStyle = lipgloss.NewStyle()

	LeftPanelHomeStyle = lipgloss.NewStyle().
				Background(BackgroundColor)
	WelcomeHeaderStyle = lipgloss.NewStyle().
				Background(AccentColor).
				Foreground(BackgroundColor).
				Bold(true).
				Align(lipgloss.Center).AlignVertical(lipgloss.Center)

	SavedConnectionHeaderStle = lipgloss.NewStyle().
					Background(BackgroundColor).
					Height(3).
					AlignVertical(lipgloss.Bottom).
					Align(lipgloss.Center).Italic(true)
	SavedConnectionStyle     = lipgloss.NewStyle()
	SavedConnectionItemStyle = lipgloss.NewStyle().
					Border(lipgloss.BlockBorder(), false, false, false, true).Background(BackgroundColor).
					BorderForeground(SecondaryColor).
					Align(lipgloss.Left).
					AlignVertical(lipgloss.Center).
					Height(5).Padding(0, 0, 0, 2)
	SavedConnectionItemSelectedStyle = lipgloss.NewStyle().
						Border(lipgloss.BlockBorder(), false, false, false, true).
						Background(BackgroundColor).Foreground(SecondaryColor).
						BorderForeground(AccentColor).
						Align(lipgloss.Left).
						AlignVertical(lipgloss.Center).
						Height(5).Padding(0, 0, 0, 2)

	RightPanelHomeStyle = lipgloss.NewStyle().Align(lipgloss.Center)
	TitleBlockStyle     = lipgloss.NewStyle().Foreground(PrimaryColor).
				Bold(true).
				Italic(true).
				AlignVertical(lipgloss.Center)
	FormStyle = lipgloss.NewStyle().
			Background(BackgroundColor).
			Border(lipgloss.RoundedBorder()).
			BorderBackground(BackgroundColor).
			BorderForeground(AccentColor).Padding(1).Align(lipgloss.Center)
	InputStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(BackgroundColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(SecondaryColor).
			Padding(0, 1).
			Width(40)

	BootomBarStyle           = lipgloss.NewStyle().Background(BackgroundColor).Align().PaddingLeft(0)
	BottomBarItemStyle       = lipgloss.NewStyle().Foreground(SecondaryColor).Background(BackgroundColor)
	BottomBarItemHeaderStyle = lipgloss.NewStyle().Background(SecondaryColor).Foreground(BackgroundColor).Bold(true)
	BottomBarItemTextStyle   = lipgloss.NewStyle().Background(BackgroundColor).Foreground(SecondaryColor).Bold(true)

	// Header for top logos/titles
	HeaderStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Bold(true).
			Underline(true).
			Padding(1, 0).
			Align(lipgloss.Center)

	// Left Panel – Logo + Past Connections
	LeftPanel = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Background(BackgroundColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2).
			Width(50).
			Height(30)

	// Right Panel – Welcome + Form
	RightPanel = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(BackgroundColor).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 4).
			Align(lipgloss.Center).
			Width(80).
			Height(30)

	// Title box – "TermSQL 🚀"
	TitleBox = lipgloss.NewStyle().
			Foreground(AccentColor).
			Border(lipgloss.HiddenBorder()).
			Bold(true).
			AlignHorizontal(lipgloss.Left).
			AlignVertical(lipgloss.Center)

	// Saved Connections list
	ConnectionHistoryBox = lipgloss.NewStyle().
				Foreground(TextColor).
				Background(lipgloss.Color("#1A1A2E")).
				Border(lipgloss.ThickBorder(), false).
				BorderForeground(SecondaryColor).
				Padding(1, 2).
				Width(44).
				MarginTop(2)

	// Welcome message
	WelcomeBanner = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Italic(true).
			AlignVertical(lipgloss.Bottom)
)
