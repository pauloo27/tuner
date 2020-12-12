package display

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

func ListPlaylists() {
	for i, playlist := range state.Data.Playlists {
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

func SaveToPlaylist(result *search.YouTubeResult, mpv *player.MPV) {
	// stop keyboard
	keyboard.Close()
	utils.ClearScreen()

	fmt.Printf("Save to:\n")
	for i, playlist := range state.Data.Playlists {
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

			if err == nil && index <= int64(len(state.Data.Playlists)) && index > 0 {
				index--
				state.Data.Playlists[index].Songs = append(state.Data.Playlists[index].Songs, result)
				storage.Save(state.Data)
			}
		} else {
			newPlaylist := &storage.Playlist{Name: rawPlaylist, Songs: []*search.YouTubeResult{result}}
			state.Data.Playlists = append(state.Data.Playlists, newPlaylist)
			storage.Save(state.Data)
		}
	}

	// restore keyboard
	mpv.Saving = false
	go keybind.Listen()
	mpv.Update()
}
