package root

import tea "github.com/charmbracelet/bubbletea"

type VisibleMsg struct {
	Width, Height int
}

func visible(width, height int) tea.Cmd {
	return func() tea.Msg {
		return VisibleMsg{width, height}
	}
}
