package search

import (
	"io"
	"time"
)

type SearchResult struct {
	Artist, Title string
	IsLive        bool
	Length        time.Duration
	URL           string
	Provider      SearchProvider
}

type AudioFormat struct {
	Name string
}

type AudioInfo struct {
	Stream io.Reader
	Size   int64
	Format *AudioFormat
}

type SearchProvider interface {
	GetName() (name string)
	SearchFor(searchQuery string) (results []*SearchResult, err error)
	GetAudioInfo(result *SearchResult) (*AudioInfo, error)
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
