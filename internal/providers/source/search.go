package source

import (
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
	MimeType string
}

type AudioInfo struct {
	StreamURL string
	Format    *AudioFormat
}

type SearchProvider interface {
	GetName() (name string)
	SearchFor(searchQuery string) (results []*SearchResult, err error)
	GetAudioInfo(result *SearchResult) (*AudioInfo, error)
}

var searchProviders []SearchProvider

func (r *SearchResult) GetAudioInfo() (*AudioInfo, error) {
	return r.Provider.GetAudioInfo(r)
}

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
