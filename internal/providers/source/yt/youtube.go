package yt

import (
	"errors"
	"time"

	"github.com/pauloo27/searchtube"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/youtube/v2"
)

type YouTubeSource struct {
	client youtube.Client
}

var _ source.Source = &YouTubeSource{}

func NewYoutubeSource() *YouTubeSource {
	return &YouTubeSource{}
}

func (*YouTubeSource) Name() (name string) {
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

	formats := video.Formats

	formats = formats.WithAudioChannels()
	formats.Sort()

	if len(formats) == 0 {
		return source.AudioInfo{}, errors.New("no audio info found")
	}
	format := formats[0]

	uri, err := s.client.GetStreamURL(video, &format)
	if err != nil {
		return source.AudioInfo{}, err
	}

	return source.AudioInfo{
		StreamURL: uri,
		Format:    source.AudioFormat{MimeType: format.MimeType},
	}, nil
}
