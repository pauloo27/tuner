package search

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
)

type searchCompletedMsg struct {
	Results []source.SearchResult
	Err     error
}

func doSearch(query string) tea.Cmd {
	return func() tea.Msg {
		var results []source.SearchResult
		for _, provider := range providers.Sources {
			r, err := provider.SearchFor(query)
			if err != nil {
				return searchCompletedMsg{nil, err}
			}
			results = append(results, r...)
		}
		return searchCompletedMsg{results, nil}
	}
}
