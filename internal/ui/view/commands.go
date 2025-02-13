package view

import tea "github.com/charmbracelet/bubbletea"

type ViewMessage interface {
	ForwardTo() ViewName
}

type GotoViewMsg struct {
	Name   ViewName
	Params []any
}

func GoToView(viewName ViewName, params ...any) func() tea.Msg {
	return func() tea.Msg {
		return GotoViewMsg{viewName, params}
	}
}

type VisibleMsg struct {
	Width, Height int
}

func Visible(width, height int) tea.Cmd {
	return func() tea.Msg {
		return VisibleMsg{width, height}
	}
}
