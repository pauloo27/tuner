package search

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultListHeight = 13
	defaultListWidth  = 20
)

var (
	textStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1).
			Background(lipgloss.Color("4")).
			Foreground(lipgloss.Color("0"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("4")).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(0).PaddingTop(0)
)
