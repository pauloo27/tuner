package search

import (
	"fmt"
	"log/slog"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/ui/view/root"
)

type model struct {
	query           string
	searchCompleted bool
	err             error
	list            list.Model
}

func NewModel() model {
	l := list.New(nil, itemDelegate{}, defaultListWidth, defaultListHeight)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.DisableQuitKeybindings()
	l.SetFilteringEnabled(false)
	// TODO: cancel (aka esc) bind to take back to the home view
	return model{list: l}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		default:
			if m.searchCompleted {
				m.list, cmd = m.list.Update(msg)
			}
		}
	case root.StartSearchMsg:
		m.query = msg.Query
		m.searchCompleted = false
		m.err = nil
		cmd = tea.Batch(doSearch(m.query), m.list.SetItems(nil))
	case searchCompletedMsg:
		slog.Info("Search completed", "res", msg.Results, "err", msg.Err)
		m.err = msg.Err
		m.searchCompleted = true
		cmd = m.list.SetItems(searchResultsToItems(msg.Results))
	}

	return m, cmd
}

func (m model) View() string {
	if !m.searchCompleted {
		return fmt.Sprintf(
			"%s\n",
			textStyle.Render(fmt.Sprintf("Searching for %s...", m.query)),
		) + "\n"
	}

	return fmt.Sprintf(
		"%s\n%s",
		textStyle.Render(fmt.Sprintf("Results for %s...", m.query)),
		m.list.View(),
	) + "\n"
}
