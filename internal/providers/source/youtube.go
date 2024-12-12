package source

import (
	"time"

	"github.com/Pauloo27/searchtube"
	"github.com/kkdai/youtube/v2"
)

type YouTubeSearch struct {
}

var _ SearchProvider = YouTubeSearch{}

func init() {
	searchProviders = append(searchProviders, &YouTubeSearch{})
}

func (YouTubeSearch) GetName() (name string) {
	return "YouTube"
}

var client youtube.Client

func (p YouTubeSearch) SearchFor(searchQuery string) ([]*SearchResult, error) {
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
			Artist:   result.Uploader,
			Title:    result.Title,
			URL:      result.URL,
			IsLive:   result.Live,
			Length:   duration,
			Provider: p,
		})
	}
	return results, nil
}

func (YouTubeSearch) GetAudioInfo(result *SearchResult) (*AudioInfo, error) {
	video, err := client.GetVideo(result.URL)
	if err != nil {
		return nil, err
	}

	format := video.Formats.FindByItag(250)
	uri, err := client.GetStreamURL(video, format)
	if err != nil {
		return nil, err
	}

	return &AudioInfo{
		StreamURL: uri,
		Format:    &AudioFormat{MimeType: format.MimeType},
	}, nil
}
