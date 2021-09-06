package search

import (
	"time"

	"github.com/Pauloo27/searchtube"
)

type YouTubeSearch struct {
}

var _ SearchProvider = &YouTubeSearch{}

func (YouTubeSearch) GetName() (name string) {
	return "YouTube"
}

func init() {
	searchProviders = append(searchProviders, &YouTubeSearch{})
}

func (YouTubeSearch) SearchFor(searchQuery string) ([]*SearchResult, error) {
	youtubeResults, err := searchtube.Search(searchQuery, 10)
	if err != nil {
		return nil, err
	}
	var results []*SearchResult
	for _, result := range youtubeResults {
		var duration time.Duration
		if !result.Live {
			duration, err = result.GetDuration()
			if err != nil {
				return nil, err
			}
		}
		results = append(results, &SearchResult{
			Artist: result.Uploader,
			Title:  result.Title,
			IsLive: result.Live,
			Length: duration,
		})
	}
	return results, nil
}
