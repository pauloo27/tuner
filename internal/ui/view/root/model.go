package root

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	searchView tea.Model
	debugView  tea.Model

	// yeah i know i should avoid pointers and etc, but happens
	activeView   *tea.Model
	previousView *tea.Model
}

func NewModel(
	searchModel tea.Model,
	debugModel tea.Model,
) model {
	return model{
		searchView: searchModel,
		debugView:  debugModel,

		activeView: &searchModel,
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
			if *m.activeView == m.debugView {
				m.activeView = m.previousView
			} else {
				m.previousView = m.activeView
				m.activeView = &m.debugView
			}
			// ensure the activeView knows its now visible
			cmd = visible
		}
	}

	var activeViewCmd tea.Cmd
	*m.activeView, activeViewCmd = (*m.activeView).Update(msg)

	return m, tea.Batch(cmd, activeViewCmd)
}

func (m model) View() string {
	return (*m.activeView).View()
}
