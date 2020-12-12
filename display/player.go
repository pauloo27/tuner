package display

import (
	"fmt"
	"sync"

	"github.com/Pauloo27/go-mpris"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/utils"
)

const (
	pausedIcon  = ""
	playingIcon = ""
)

var (
	horizontalBars = []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}
)

var updateLock sync.Mutex

func ShowPlaying(result *search.YouTubeResult, mpv *player.MPV) {
	if !state.Playing || mpv.Saving {
		return
	}

	updateLock.Lock()
	defer updateLock.Unlock()

	utils.ClearScreen()

	icon := playingIcon

	playback, _ := mpv.Player.GetPlaybackStatus()
	if playback != mpris.PlaybackPlaying {
		icon = pausedIcon
	}

	extra := ""
	if status, err := mpv.Player.GetLoopStatus(); err == nil {
		if status == mpris.LoopTrack {
			extra += utils.ColorWhite + "  "
		} else if status == mpris.LoopPlaylist {
			extra += utils.ColorBlue + "  "
		}
	}

	if mpv.IsPlaylist() {
		fmt.Printf("Playing: %s (%d/%d)\n",
			mpv.Playlist.Name,
			mpv.PlaylistIndex+1,
			len(mpv.Playlist.Songs),
		)
	}

	fmt.Printf(" %s  %s %sfrom %s%s%s\n",
		icon,
		utils.ColorGreen+result.Title,
		utils.ColorWhite,
		utils.ColorGreen+result.Uploader,
		extra,
		utils.ColorReset,
	)

	if !mpv.GetPlaying().Live {
		length, err := mpv.Player.GetLength()
		if err == nil {
			position, err := mpv.Player.GetPosition()
			if err == nil && position > 0 {
				columns := float64(utils.GetTerminalSize().Col)
				barSize := columns * float64(len(horizontalBars))
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
	}

	if status, _ := mpv.Player.GetPlaybackStatus(); status != "" {
		volume, _ := mpv.Player.GetVolume()
		fmt.Printf("Volume: %s%.0f%%%s\n", utils.ColorGreen, volume*100, utils.ColorReset)
	}

	if mpv.ShowURL {
		fmt.Printf("%s%s%s\n", utils.ColorBlue, result.URL(), utils.ColorReset)
	}

	if mpv.ShowHelp {
		fmt.Println("\n" + utils.ColorBlue + "Keybinds:")
		for _, bind := range keybind.ListBinds() {
			fmt.Printf("  %s: %s\n", bind.KeyName, bind.Description)
		}
	}

	if mpv.ShowLyric {
		fmt.Println(utils.ColorBlue)
		lines := len(mpv.LyricLines)
		if lines == 0 {
			fmt.Println("Fetching lyric...")
		}
		for i := mpv.LyricIndex; i < mpv.LyricIndex+15; i++ {
			if i == lines {
				break
			}
			fmt.Println(mpv.LyricLines[i])
		}
	}

	fmt.Print(utils.ColorReset)

}
