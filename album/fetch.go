package album

import (
	"github.com/Pauloo27/tuner/img"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/utils"
	"golang.org/x/term"
)

func FetchAlbum(result *search.YouTubeResult, mpv *player.MPV) {
	if !state.Data.FetchAlbum || mpv.Exitted {
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

		size := 25
		width, _, err := term.GetSize(0)
		utils.HandleError(err, "Cannot get term size")
		x := width - size

		img.SendCommand(
			utils.Fmt(`{"action": "add", "x": %d, "y": 1, "width": %d, "height": %d, "path": "%s", "identifier": "album"}`,
				x, size, size, path,
			),
		)
	}()
}
