package search

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/ui/view"
)

type startSearchMsg struct {
}

func startSearch() tea.Msg {
	return startSearchMsg{}
}

type searchCompletedMsg struct {
	Results []source.SearchResult
	Err     error
}

var _ view.ViewMessage = searchCompletedMsg{}

func (searchCompletedMsg) ForwardTo() view.ViewName {
	return view.SearchViewName
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
