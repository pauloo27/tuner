package yt

import (
	"errors"
	"fmt"
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
	// since the go lib is not working, we leave it up to mpv to fetch the stream
	// url
	return source.AudioInfo{
		StreamURL: result.URL,
	}, nil

	// dead code::::::::
	video, err := s.client.GetVideo(result.URL)
	if err != nil {
		return source.AudioInfo{}, err
	}

	if result.IsLive {
		return s.getLiveAudioInfo(video)
	}

	formats := video.Formats.WithAudioChannels()
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

func (s *YouTubeSource) getLiveAudioInfo(video *youtube.Video) (source.AudioInfo, error) {
	// TODO: skill issue on my side, but i couldn't get the live to play...
	// so we fallback to youtube-dl integrated in MPV
	return source.AudioInfo{
		StreamURL: fmt.Sprintf("https://youtu.be/%s", video.ID),
		Format:    source.AudioFormat{},
	}, nil
}
