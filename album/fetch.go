package album

import (
	"path"
	"time"

	"github.com/Pauloo27/tuner/img"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/utils"
	"golang.org/x/term"
)

var albumPath string

func listenToResizes() {
	size := 25
	lastWidth := -1
	for {
		if albumPath == "" {
			continue
		}
		time.Sleep(1 * time.Second)

		width, _, err := term.GetSize(0)
		utils.HandleError(err, "Cannot get term size")
		if lastWidth == width {
			continue
		}
		x := width - size

		img.SendCommand(
			utils.Fmt(`{"action": "add", "x": %d, "y": 2, "width": %d, "height": %d, "path": "%s", "identifier": "album"}`,
				x, size, size, albumPath,
			),
		)
	}
}

func RegisterHooks() {
	if !player.State.Data.FetchAlbum {
		return
	}

	player.RegisterHook(func(param ...interface{}) {
		fetchAlbum()
	}, player.HookFileLoaded)

	player.RegisterHook(func(param ...interface{}) {
		albumPath = ""
		img.SendCommand(`{"action": "remove", "identifier": "album"}`)
	}, player.HookFileEnded)

	go listenToResizes()
}

func fetchAlbum() {
	if !player.State.Data.FetchAlbum {
		return
	}
	go func() {
		result := player.State.GetPlaying()

		var artURL string

		if result.SourceName == "soundcloud" {
			artURL = result.Extra[0]
		} else {
			videoInfo, err := FetchVideoInfo(result)
			if err != nil {
				return
			}
			artURL = utils.Fmt("https://i1.ytimg.com/vi/%s/hqdefault.jpg", videoInfo.ID)
			if videoInfo.Artist != "" && videoInfo.Track != "" {
				trackInfo, err := FetchTrackInfo(videoInfo.Artist, videoInfo.Track)
				if err == nil {
					artURL = trackInfo.Album.ImageURL
				}
			}
		}

		path := path.Join(utils.LoadDataFolder(), "album")
		utils.DownloadFile(artURL, path)

		albumPath = path
	}()
}
