package root

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	homeView   tea.Model
	searchView tea.Model
	// yeah i know i should avoid pointers and etc, but happens
	activeView *tea.Model
}

func NewModel(
	homeModel tea.Model,
	searchView tea.Model,
) model {
	return model{
		homeView:   homeModel,
		searchView: searchView,
		activeView: &homeModel,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case StartSearchMsg:
		m.activeView = &m.searchView
	}

	updatedActiveView, cmd := (*m.activeView).Update(msg)
	*m.activeView = updatedActiveView
	return m, cmd
}

func (m model) View() string {
	return (*m.activeView).View()
}
