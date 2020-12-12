package album

import (
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/utils"
)

func FetchAlbum(result *search.YouTubeResult, mpv *player.MPV) {
	if !state.Data.FetchAlbum {
		return
	}
	go func() {

		videoInfo, err := FetchVideoInfo(result)
		if err != nil {
			return
		}

		trackInfo, err := FetchTrackInfo(videoInfo.Artist, videoInfo.Track)
		if err != nil {
			return
		}

		utils.DownloadFile(trackInfo.Album.ImageURL, utils.LoadDataFolder()+"/album")
	}()
}
