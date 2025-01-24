package search

import (
	"fmt"
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui/view/root"
)

type model struct {
	query           string
	searchCompleted bool
	results         []source.SearchResult
	err             error
}

func NewModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case root.StartSearchMsg:
		m.query = msg.Query
		m.searchCompleted = false
		m.err = nil
		m.results = nil
		cmd = doSearch(m.query)
	case searchCompletedMsg:
		slog.Info("Search completed", "res", msg.Results, "err", msg.Err)
		m.results = msg.Results
		m.err = msg.Err
		m.searchCompleted = true
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

	return "hello"
}
