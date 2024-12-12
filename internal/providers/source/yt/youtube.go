package yt

import (
	"time"

	"github.com/kkdai/youtube/v2"
	"github.com/pauloo27/searchtube"
	"github.com/pauloo27/tuner/internal/providers/source"
)

type YouTubeSource struct {
	client youtube.Client
}

var _ source.Source = &YouTubeSource{}

func NewYoutubeSource() *YouTubeSource {
	return &YouTubeSource{}
}

func (*YouTubeSource) GetName() (name string) {
	return "YouTube"
}

func (s *YouTubeSource) SearchFor(searchQuery string) ([]source.SearchResult, error) {
	youtubeResults, err := searchtube.Search(searchQuery, 10)
	if err != nil {
		return nil, err
	}
	var results []source.SearchResult
	for _, result := range youtubeResults {
		var duration time.Duration
		if !result.Live {
			duration, err = result.GetDuration()
			if err != nil {
				return nil, err
			}
		}
		results = append(results, source.SearchResult{
			Artist: result.Uploader,
			Title:  result.Title,
			URL:    result.URL,
			IsLive: result.Live,
			Length: duration,
			Source: s,
		})
	}
	return results, nil
}

func (s *YouTubeSource) GetAudioInfo(result source.SearchResult) (source.AudioInfo, error) {
	video, err := s.client.GetVideo(result.URL)
	if err != nil {
		return source.AudioInfo{}, err
	}

	format := video.Formats.FindByItag(250)
	uri, err := s.client.GetStreamURL(video, format)
	if err != nil {
		return source.AudioInfo{}, err
	}

	return source.AudioInfo{
		StreamURL: uri,
		Format:    source.AudioFormat{MimeType: format.MimeType},
	}, nil
}
