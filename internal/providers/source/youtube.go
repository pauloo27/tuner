package source

import (
	"strings"
	"time"

	"github.com/Pauloo27/searchtube"
	"github.com/kkdai/youtube/v2"
)

type YouTubeSearch struct {
}

var _ SearchProvider = YouTubeSearch{}

func (YouTubeSearch) GetName() (name string) {
	return "YouTube"
}

var client youtube.Client

func (YouTubeSearch) GetAudioInfo(result *SearchResult) (*AudioInfo, error) {
	id := strings.Split(result.URL, "=")[1]

	video, err := client.GetVideo(id)
	if err != nil {
		panic(err)
	}

	stream, size, err := client.GetStream(video, video.Formats.FindByItag(249))
	if err != nil {
		panic(err)
	}

	return &AudioInfo{
		Stream: stream,
		Size:   size,
		Format: FormatWebm,
	}, nil
}

func init() {
	searchProviders = append(searchProviders, &YouTubeSearch{})
}

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
