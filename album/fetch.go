package album

import (
	"github.com/Pauloo27/tuner/img"
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
		path := utils.LoadDataFolder() + "/album"
		utils.DownloadFile(trackInfo.Album.ImageURL, path)

		size := 30
		x := int(utils.GetTerminalSize().Col) - size

		img.SendCommand(
			utils.Fmt(`{"action": "add", "x": %d, "y": 0, "width": %d, "height": %d, "path": "%s", "identifier": "asda"}`,
				x, size, size, path,
			),
		)
	}()
}
