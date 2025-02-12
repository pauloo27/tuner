package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/ui/view"
	"github.com/pauloo27/tuner/internal/ui/view/debug"
	"github.com/pauloo27/tuner/internal/ui/view/root"
	"github.com/pauloo27/tuner/internal/ui/view/search"
)

func StartTUI() error {
	views := map[view.ViewName]tea.Model{
		view.SearchViewName: search.NewModel(),
		view.DebugViewName:  debug.NewModel(),
	}

	rootModel := root.NewModel(
		views,
		view.SearchViewName,
	)
	p := tea.NewProgram(rootModel, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
