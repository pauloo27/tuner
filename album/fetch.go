package album

import (
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
)

func FetchAlbum(result *search.YouTubeResult, mpv *player.MPV) {
	videoInfo, err := FetchVideoInfo(result)
	if err != nil {
		return
	}

	trackInfo, err := FetchTrackInfo(videoInfo.Artist, videoInfo.Track)
	if err != nil {
		return
	}

	// TODO:
	//utils.Download(trackInfo.Album.ImageURL, "album")
}
