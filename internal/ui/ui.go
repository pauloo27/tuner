package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/ui/view/home"
	"github.com/pauloo27/tuner/internal/ui/view/root"
	"github.com/pauloo27/tuner/internal/ui/view/search"
)

func StartTUI() error {
	homeModel := home.NewModel()
	searchModel := search.NewModel()

	rootModel := root.NewModel(
		homeModel,
		searchModel,
	)
	p := tea.NewProgram(rootModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
