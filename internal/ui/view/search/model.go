package search

import (
	"fmt"
	"log/slog"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/core"
)

type model struct {
	isTyping        bool
	searchCompleted bool
	err             error
	list            list.Model
	searchInput     textinput.Model
}

func NewModel() model {
	ti := textinput.New()
	ti.Placeholder = "Search..."
	ti.Focus()
	ti.CharLimit = 100

	l := list.New(nil, itemDelegate{}, defaultListWidth, defaultListHeight)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.DisableQuitKeybindings()
	l.SetFilteringEnabled(false)
	// TODO: cancel (aka esc) bind to take back to the home view

	return model{isTyping: true, list: l, searchInput: ti}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.isTyping = false
			cmd = startSearch
		default:
			if m.isTyping {
				m.searchInput, cmd = m.searchInput.Update(msg)
			} else if m.searchCompleted {
				m.list, cmd = m.list.Update(msg)
			}
		}
	case startSearchMsg:
		m.searchCompleted = false
		m.err = nil
		cmd = tea.Batch(doSearch(m.searchInput.Value()), m.list.SetItems(nil))
	case searchCompletedMsg:
		slog.Info("Search completed", "res", msg.Results, "err", msg.Err)
		m.err = msg.Err
		m.searchCompleted = true
		cmd = m.list.SetItems(searchResultsToItems(msg.Results))
	}

	return m, cmd
}

func (m model) View() string {
	if m.isTyping {
		return fmt.Sprintf(
			"%s\n\n%s",
			textStyle.Render(fmt.Sprintf("Tuner - %s", core.Version)),
			m.searchInput.View(),
		) + "\n"
	}

	if !m.searchCompleted {
		return fmt.Sprintf(
			"%s\n",
			textStyle.Render(fmt.Sprintf("Searching for %s...", m.searchInput.Value())),
		) + "\n"
	}

	if m.err != nil {
		return fmt.Sprintf(
			"%s\n",
			errorStyle.Render(fmt.Sprintf("An error occurred: %v", m.err)),
		) + "\n"
	}

	return fmt.Sprintf(
		"%s\n%s",
		textStyle.Render(fmt.Sprintf("Results for %s...", m.searchInput.Value())),
		m.list.View(),
	) + "\n"
}
