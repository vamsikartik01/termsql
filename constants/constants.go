package constants

import "termsql/types"

var (
	DefaultDatabaseHelp = []types.Help{
		{Key: "Ctrl-C", Title: "Exit"},
		{Key: "Up/Down", Title: "Navigate"},
		{Key: "Ctrl-N", Title: "New Database"},
		{Key: "Ctrl-R", Title: "Refresh"},
	}
	DefaultDatabaseLeftPanelWidth = 64
)
