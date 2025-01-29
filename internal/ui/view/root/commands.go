package root

import tea "github.com/charmbracelet/bubbletea"

type VisibleMsg struct {
}

func visible() tea.Msg {
	return VisibleMsg{}
}
