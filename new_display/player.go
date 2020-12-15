package new_display

import (
	"fmt"
	"sync"

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

func DisplayPlayer() {
	var lock sync.Mutex

	render := func() {
		lock.Lock()
		defer lock.Unlock()
		utils.ClearScreen()

		result := new_player.State.GetPlaying()

		if !result.Live {
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

		/*
			if mpv.IsPlaylist() {
				fmt.Printf("Playing: %s (%d/%d)\n",
					mpv.Playlist.Name,
					mpv.PlaylistIndex+1,
					len(mpv.Playlist.Songs),
				)
			}
		*/

		icon := playingIcon

		if new_player.State.Paused {
			icon = pausedIcon
		}

		extra := ""
		// TODO: loop
		/*
			if status, err := mpv.Player.GetLoopStatus(); err == nil {
				if status == mpris.LoopTrack {
					extra += utils.ColorWhite + "  "
				} else if status == mpris.LoopPlaylist {
					extra += utils.ColorBlue + "  "
				}
			}
		*/

		fmt.Printf(" %s  %s %sfrom %s%s%s\n",
			icon,
			utils.ColorGreen+result.Title,
			utils.ColorWhite,
			utils.ColorGreen+result.Uploader,
			extra,
			utils.ColorReset,
		)
		fmt.Printf("Volume: %s%.0f%%%s\n", utils.ColorGreen, new_player.State.Volume, utils.ColorReset)
	}

	new_player.RegisterHooks(
		func(param ...interface{}) {
			render()
		},
		new_player.HOOK_PLAYBACK_PAUSED, new_player.HOOK_PLAYBACK_RESUMED,
		new_player.HOOK_VOLUME_CHANGED, new_player.HOOK_POSITION_CHANGED,
	)

	render()
}
