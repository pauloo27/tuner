package root

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	searchView tea.Model
	debugView  tea.Model

	isInDebug bool
}

func NewModel(
	searchModel tea.Model,
	debugModel tea.Model,
) model {
	return model{
		searchView: searchModel,
		debugView:  debugModel,
	}
}

func (m model) Init() tea.Cmd {
	// TODO: send the visible to the initialActiveView?
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlD:
			m.isInDebug = !m.isInDebug
			// ensure the activeView knows its now visible
			cmd = visible
		}
	}

	var viewCmd tea.Cmd

	if m.isInDebug {
		m.debugView, viewCmd = m.debugView.Update(msg)
	} else {
		m.searchView, viewCmd = m.searchView.Update(msg)
	}

	return m, tea.Batch(cmd, viewCmd)
}

func (m model) View() string {
	if m.isInDebug {
		return m.debugView.View()
	}

	return m.searchView.View()
}
