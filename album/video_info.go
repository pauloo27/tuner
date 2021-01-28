package album

import (
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

	id, _ := jsonparser.GetString(buffer, "id")
	duration, _ := jsonparser.GetInt(buffer, "duration")
	title, _ := jsonparser.GetString(buffer, "title")
	artist, _ := jsonparser.GetString(buffer, "artist")
	track, _ := jsonparser.GetString(buffer, "track")

	return &VideoInfo{artist, track, title, id, duration}, nil
}
