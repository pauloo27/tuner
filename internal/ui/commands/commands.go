package commands

import tea "github.com/charmbracelet/bubbletea"

type SearchCommand struct {
	Query string
}

func Search(query string) tea.Cmd {
	return func() tea.Msg {
		return SearchCommand{query}
	}
}
