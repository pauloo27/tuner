package display

import (
	"fmt"
	"sync"

	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/new_player"
	"github.com/Pauloo27/tuner/utils"
	"golang.org/x/term"
)

const (
	pausedIcon  = ""
	playingIcon = ""
)

var (
	horizontalBars = []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}
)

func startPlayerHooks() {
	var lock sync.Mutex

	render := func() {
		lock.Lock()
		defer lock.Unlock()

		result := new_player.State.GetPlaying()
		if result == nil {
			return
		}

		utils.ClearScreen()

		// TODO: fix the progress bar
		if !result.Live && false {
			length := new_player.State.Duration
			position, err := new_player.GetPosition()
			if err == nil && position > 0 {
				columns, _, err := term.GetSize(0)
				utils.HandleError(err, "Cannot get term size")
				barSize := float64(columns) * float64(len(horizontalBars))
				progress := int((barSize * position) / length)
				fullBlocks := progress / len(horizontalBars)
				missing := progress % len(horizontalBars)
				fmt.Print(utils.ColorBlue)
				for i := 0; i < fullBlocks; i++ {
					fmt.Print(horizontalBars[len(horizontalBars)-1])
				}
				if missing != 0 {
					fmt.Print(horizontalBars[missing-1])
				}
				fmt.Println(utils.ColorReset)
			}
		}

		if new_player.State.IsPlaylist() {
			fmt.Printf("Playing: %s (%d/%d)\n",
				new_player.State.Playlist.Name,
				new_player.State.PlaylistIndex+1,
				len(new_player.State.Playlist.Songs),
			)
		}

		icon := playingIcon

		if new_player.State.Paused {
			icon = pausedIcon
		}

		extra := ""
		switch new_player.State.Loop {
		case new_player.LOOP_TRACK:
			extra += utils.ColorWhite + "  "
		case new_player.LOOP_PLAYLIST:
			extra += utils.ColorBlue + "  "
		}

		fmt.Printf(" %s  %s %sfrom %s%s%s\n",
			icon,
			utils.ColorGreen+result.Title,
			utils.ColorWhite,
			utils.ColorGreen+result.Uploader,
			extra,
			utils.ColorReset,
		)

		fmt.Printf("Volume: %s%.0f%%%s\n", utils.ColorGreen, new_player.State.Volume, utils.ColorReset)

		if new_player.State.ShowURL {
			fmt.Printf("%s%s%s\n", utils.ColorBlue, result.URL(), utils.ColorReset)
		}

		if new_player.State.ShowHelp {
			fmt.Println("\n" + utils.ColorBlue + "Keybinds:")
			for _, bind := range keybind.ListBinds() {
				fmt.Printf("  %s: %s\n", bind.KeyName, bind.Description)
			}
		}

		if new_player.State.ShowLyric {
			fmt.Println(utils.ColorBlue)
			lyric := new_player.State.Lyric
			lines := len(lyric.Lines)
			if lines == 0 {
				fmt.Println("Fetching lyric...")
			}
			for i := lyric.Index; i < lyric.Index+14; i++ {
				if i == lines {
					break
				}
				fmt.Println(lyric.Lines[i])
			}
		}

		fmt.Print(utils.ColorReset)
	}

	new_player.RegisterHooks(
		func(param ...interface{}) {
			render()
		},
		new_player.HOOK_PLAYBACK_PAUSED, new_player.HOOK_PLAYBACK_RESUMED,
		new_player.HOOK_VOLUME_CHANGED, new_player.HOOK_POSITION_CHANGED,
		new_player.HOOK_GENERIC_UPDATE, new_player.HOOK_LOOP_PLAYLIST_CHANGED,
		new_player.HOOK_LOOP_TRACK_CHANGED, new_player.HOOK_FILE_LOAD_STARTED,
		new_player.HOOK_FILE_ENDED,
	)
}
