package root

import (
	tea "github.com/charmbracelet/bubbletea"
)

type StartSearchMsg struct {
	Query string
}

func StartSearch(query string) tea.Cmd {
	return func() tea.Msg {
		return StartSearchMsg{query}
	}
}
