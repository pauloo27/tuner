package display

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

func ListPlaylists() {
	for i, playlist := range player.State.Data.Playlists {
		bold := ""
		if i%2 == 0 {
			bold = utils.ColorBold
		}

		defaultColor := bold + utils.ColorWhite
		altColor := bold + utils.ColorGreen

		fmt.Printf("  %s#%d: %s%s\n",
			defaultColor, i+1,
			altColor+playlist.Name,
			utils.ColorReset,
		)
	}
}

func startPlaylistSaveHooks() {
	player.RegisterHook(saveToPlaylist, player.HOOK_SAVING_TRACK_TO_PLAYLIST)
}

func saveToPlaylist(params ...interface{}) {
	// stop keyboard
	keyboard.Close()
	utils.ClearScreen()

	result := player.State.GetPlaying()

	fmt.Printf("Save to:\n")
	for i, playlist := range player.State.Data.Playlists {
		bold := ""
		if i%2 == 0 {
			bold = utils.ColorBold
		}

		defaultColor := bold + utils.ColorWhite
		altColor := bold + utils.ColorGreen

		fmt.Printf("  %s#%d: %s%s\n",
			defaultColor, i+1,
			altColor+playlist.Name,
			utils.ColorReset,
		)
	}

	rawPlaylist, err := utils.AskFor("Save to (#<id>, the name of a new one or nothing to cancel)")
	utils.HandleError(err, "Cannot read user input")

	if rawPlaylist != "" {
		if strings.HasPrefix(rawPlaylist, "#") {
			index, err := strconv.ParseInt(strings.TrimPrefix(rawPlaylist, "#"), 10, 64)

			if err == nil && index <= int64(len(player.State.Data.Playlists)) && index > 0 {
				index--
				player.State.Data.Playlists[index].Songs = append(player.State.Data.Playlists[index].Songs, result)
				storage.Save(player.State.Data)
			}
		} else {
			newPlaylist := &storage.Playlist{Name: rawPlaylist, Songs: []*search.YouTubeResult{result}}
			player.State.Data.Playlists = append(player.State.Data.Playlists, newPlaylist)
			storage.Save(player.State.Data)
		}
	}

	// restore keyboard
	player.State.SavingToPlaylist = false
	go keybind.Listen()
	player.ForceUpdate()
}
