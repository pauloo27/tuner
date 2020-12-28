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

	if player.State.IsPlaylist() {
		fmt.Printf("%sTo edit the current playlist, use:\n", utils.ColorYellow)
		fmt.Printf("  %s%sD%s%s to delete current playlist\n%s",
			utils.ColorYellow, utils.ColorBold, utils.ColorReset, utils.ColorYellow,
			utils.ColorReset,
		)
		fmt.Printf("  %s%sR%s%s to remove the current song from the playlist\n%s",
			utils.ColorYellow, utils.ColorBold, utils.ColorReset, utils.ColorYellow,
			utils.ColorReset,
		)
	}
	rawPlaylist, err := utils.AskFor(
		"Save to (#<id>, a name (create if not exists) or nothing to cancel)",
	)
	utils.HandleError(err, "Cannot read user input")

	appendToPlaylist := func(index int) {
		player.State.Data.Playlists[index].Songs = append(player.State.Data.Playlists[index].Songs, result)
		storage.Save(player.State.Data)
	}

	if rawPlaylist != "" {
		if strings.HasPrefix(rawPlaylist, "#") {
			index, err := strconv.ParseInt(strings.TrimPrefix(rawPlaylist, "#"), 10, 64)
			if err == nil && index <= int64(len(player.State.Data.Playlists)) && index > 0 {
				appendToPlaylist(int(index - 1))
			}
		} else {
			switch strings.ToLower(rawPlaylist) {
			case "r":
				if player.State.IsPlaylist() {
					songs := []*search.SearchResult{}
					for _, song := range player.State.Playlist.Songs {
						if song.ID != result.ID {
							songs = append(songs, song)
						}
					}
					player.State.Playlist.Songs = songs
					storage.Save(player.State.Data)
					player.PlaySearchResult(nil, player.State.Playlist)
				}
			case "d":
				if player.State.IsPlaylist() {
					filteredPlaylist := []*storage.Playlist{}
					for _, playlist := range player.State.Data.Playlists {
						if playlist.Name != player.State.Playlist.Name {
							filteredPlaylist = append(filteredPlaylist, playlist)
						}
					}
					player.State.Data.Playlists = filteredPlaylist
					storage.Save(player.State.Data)
					player.Stop()
				}
			default:
				duplicate := false
				for index, playlist := range player.State.Data.Playlists {
					if playlist.Name == rawPlaylist {
						duplicate = true
						appendToPlaylist(index)
						break
					}
				}
				if !duplicate {
					newPlaylist := &storage.Playlist{Name: rawPlaylist, Songs: []*search.SearchResult{result}}
					player.State.Data.Playlists = append(player.State.Data.Playlists, newPlaylist)
					storage.Save(player.State.Data)
				}
			}
		}
	}

	// restore keyboard
	player.State.SavingToPlaylist = false
	go keybind.Listen()
	player.ForceUpdate()
}
