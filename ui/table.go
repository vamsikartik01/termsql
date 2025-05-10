package ui

import (
	"termsql/types"

	"github.com/charmbracelet/bubbles/textinput"
)

type TableSelectorModel struct {
	tables         []string
	filteredTables []string
	textInput      textinput.Model
	cursor         int
	page           int
	pageSize       int
	selectedTable  string
	windowWidth    int
	windowHeight   int
	conn           types.Connection
	selectedDB     string
}
