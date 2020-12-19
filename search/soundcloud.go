package search

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/Pauloo27/tuner/utils"
	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

type SoundCloudSource struct{}

var sc *soundcloudapi.API

var re = regexp.MustCompile(`large\.(\w{3})$`)

func (s *SoundCloudSource) Search(query string, limit int) (results []*SearchResult) {
	var err error
	if sc == nil {
		sc, err = soundcloudapi.New("", http.DefaultClient)
		utils.HandleError(err, "Cannot start soundcloud API")
	}

	searchQuery, err := sc.Search(soundcloudapi.SearchOptions{
		Query: query,
	})
	utils.HandleError(err, "Cannot search soundcloud")

	tracks, err := searchQuery.GetTracks()
	utils.HandleError(err, "Cannot get tracks from query result")

	for _, track := range tracks {
		if len(results) >= limit {
			return
		}

		duration := utils.FormatTime(int(track.FullDurationMS / 1000))
		album := re.ReplaceAllString(track.ArtworkURL, "original.$1")

		results = append(results, &SearchResult{
			SourceName: "soundcloud",
			Title:      track.Title,
			Uploader:   track.User.Username,
			ID:         strconv.Itoa(int(track.ID)),
			Duration:   duration,
			URL:        track.URI,
			Extra:      []string{album},
		})
	}
	return
}
