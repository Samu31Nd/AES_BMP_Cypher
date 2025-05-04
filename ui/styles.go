package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	fourMarginLeft      = lipgloss.NewStyle().MarginLeft(4)
	borderUnder         = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false).BorderBottom(true).MarginLeft(4)
	prevTitle           = lipgloss.NewStyle().Background(lipgloss.Color("6")).Width(2).Border(lipgloss.NormalBorder(), false)
	titleStyle          = lipgloss.NewStyle().Bold(true).PaddingLeft(1)
	titleListStyle      = lipgloss.NewStyle().MarginLeft(2)
	itemStyle           = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle   = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle     = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle           = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	errorStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	filePickerHelpStyle = helpStyle.Foreground(lipgloss.Color("240"))
)
