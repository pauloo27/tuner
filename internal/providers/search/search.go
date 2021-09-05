package search

type SearchResult struct {
	Artist, Title, Length string
	IsLive                bool
}

type SearchProvider interface {
	GetName() (name string)
	SearchFor(searchQuery string) []*SearchResult
}

var searchProviders []SearchProvider

func SearchInAll(query string) (results []*SearchResult) {
	for _, provider := range searchProviders {
		r := provider.SearchFor(query)
		results = append(results, r...)
	}
	return
}
