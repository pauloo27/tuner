package home

import "github.com/charmbracelet/lipgloss"

var (
	textStyle = lipgloss.NewStyle().
		Bold(true).
		PaddingLeft(1).
		PaddingRight(1).
		Background(lipgloss.Color("4")).
		Foreground(lipgloss.Color("0"))
)
