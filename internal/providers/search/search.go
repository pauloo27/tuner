package search

import "time"

type SearchResult struct {
	Artist, Title string
	IsLive        bool
	Length        time.Duration
}

type SearchProvider interface {
	GetName() (name string)
	SearchFor(searchQuery string) (results []*SearchResult, err error)
}

var searchProviders []SearchProvider

func SearchInAll(query string) (results []*SearchResult, err error) {
	for _, provider := range searchProviders {
		r, err := provider.SearchFor(query)
		if err != nil {
			return nil, err
		}
		results = append(results, r...)
	}
	return
}
