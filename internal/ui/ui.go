package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/ui/view/home"
)

type rootModel struct {
	activePage tea.Model
}

func StartTUI() error {
	root := rootModel{
		activePage: home.InitialModel(),
	}

	p := tea.NewProgram(root, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (r rootModel) Init() tea.Cmd {
	return nil
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.activePage.Update(msg)
}

func (m rootModel) View() string {
	return m.activePage.View()
}
