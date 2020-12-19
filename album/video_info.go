package album

import (
	"fmt"
	"os/exec"

	"github.com/Pauloo27/tuner/search"
	"github.com/buger/jsonparser"
)

type VideoInfo struct {
	Artist, Track, Title, ID string
	Duration                 int64
}

var youtubeDLPath = ""

func GetYouTubeDLPath() string {
	return "/usr/bin/youtube-dl"
}

func FetchVideoInfo(result *search.SearchResult) (*VideoInfo, error) {
	if youtubeDLPath == "" {
		youtubeDLPath = GetYouTubeDLPath()
	}

	cmd := exec.Command(youtubeDLPath, result.URL, "-j")

	buffer, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	artist, err := jsonparser.GetString(buffer, "artist")
	if err != nil {
		return nil, fmt.Errorf("Cannot get artist: %v", err)
	}
	track, err := jsonparser.GetString(buffer, "track")
	if err != nil {
		return nil, fmt.Errorf("Cannot get track: %v", err)
	}
	title, err := jsonparser.GetString(buffer, "title")
	if err != nil {
		return nil, fmt.Errorf("Cannot get title: %v", err)
	}
	id, err := jsonparser.GetString(buffer, "id")
	if err != nil {
		return nil, fmt.Errorf("Cannot get id: %v", err)
	}
	duration, err := jsonparser.GetInt(buffer, "duration")
	if err != nil {
		return nil, fmt.Errorf("Cannot get duration: %v", err)
	}

	return &VideoInfo{artist, track, title, id, duration}, nil
}
