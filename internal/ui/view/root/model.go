package root

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/ui/view"
)

type model struct {
	views          map[view.ViewName]tea.Model
	activeViewName view.ViewName
	width, height  int

	isInDebug bool
}

func NewModel(
	views map[view.ViewName]tea.Model,
	initialViewName view.ViewName,
) model {
	return model{
		views:          views,
		activeViewName: initialViewName,
	}
}

func (m model) Init() tea.Cmd {
	// TODO: send the visible to the initialActiveView?
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	visibleViewName := m.activeViewName
	if m.isInDebug {
		visibleViewName = view.DebugViewName
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlD:
			m.isInDebug = !m.isInDebug
			slog.Info("Debug toggled", "isInDebug", m.isInDebug)
			// ensure the activeView knows its now visible
			cmds = append(cmds, visible(m.width, m.height))
		}
	case view.ViewMessage:
		forwardTo := msg.ForwardTo()
		if visibleViewName != forwardTo {
			var viewCmd tea.Cmd
			m.views[forwardTo], viewCmd = m.views[forwardTo].Update(msg)
			cmds = append(cmds, viewCmd)
		}
	}

	var viewCmd tea.Cmd

	m.views[visibleViewName], viewCmd = m.views[visibleViewName].Update(msg)

	cmds = append(cmds, viewCmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.isInDebug {
		return m.views[view.DebugViewName].View()
	}

	return m.views[m.activeViewName].View()
}
