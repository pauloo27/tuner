package source

import (
	"time"
)

type SearchResult struct {
	Artist, Title string
	IsLive        bool
	Length        time.Duration
	URL           string
	Source        Source
}

type AudioFormat struct {
	MimeType string
}

type AudioInfo struct {
	StreamURL string
	Format    AudioFormat
}

type Source interface {
	GetName() (name string)
	SearchFor(searchQuery string) (results []SearchResult, err error)
	GetAudioInfo(result SearchResult) (AudioInfo, error)
}

func (r SearchResult) GetAudioInfo() (AudioInfo, error) {
	return r.Source.GetAudioInfo(r)
}
