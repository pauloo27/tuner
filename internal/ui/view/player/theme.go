package player

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryTextStyle = lipgloss.NewStyle().
				Bold(true).
				PaddingLeft(1).
				PaddingRight(1).
				Background(lipgloss.Color("2")).
				Foreground(lipgloss.Color("0"))

	secondaryTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("2"))

	errorStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1).
			Background(lipgloss.Color("1")).
			Foreground(lipgloss.Color("0"))

	progressFill = lipgloss.Color("4")
)
