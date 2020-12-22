package display

import (
	"fmt"
	"sync"
	"time"

	"github.com/Pauloo27/tuner/icons"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/utils"
	"golang.org/x/term"
)

func startPlayerHooks() {
	var playerLock sync.Mutex
	var progressBarLock sync.Mutex

	renderProgressBar := func() {
		if player.State.SavingToPlaylist {
			return
		}

		progressBarLock.Lock()
		defer progressBarLock.Unlock()

		result := player.State.GetPlaying()
		if result == nil {
			return
		}

		if result.Live {
			return
		}

		utils.MoveCursorTo(1, 1)
		utils.ClearLine()
		utils.HideCursor()

		length := player.State.Duration
		position, err := player.GetPosition()
		if position < 0 {
			position = 0
		}
		if err == nil {
			columns, _, err := term.GetSize(0)
			utils.HandleError(err, "Cannot get term size")
			barSize := float64(columns) * float64(len(icons.HORIZONTAL_BARS))
			progress := int((barSize * position) / length)
			fullBlocks := progress / len(icons.HORIZONTAL_BARS)
			missing := progress % len(icons.HORIZONTAL_BARS)
			fmt.Print(utils.ColorBlue)
			for i := 0; i < fullBlocks; i++ {
				fmt.Print(icons.HORIZONTAL_BARS[len(icons.HORIZONTAL_BARS)-1])
			}
			if missing != 0 {
				fmt.Print(icons.HORIZONTAL_BARS[missing-1])
			}
			fmt.Println(utils.ColorReset)
		}
	}

	renderPlayer := func() {
		if player.State.SavingToPlaylist {
			return
		}

		playerLock.Lock()
		defer playerLock.Unlock()

		result := player.State.GetPlaying()
		if result == nil {
			return
		}

		if result.Live {
			utils.MoveCursorTo(1, 0)
		} else {
			utils.MoveCursorTo(2, 0)
		}
		utils.ClearAfterCursor()

		if player.State.IsPlaylist() {
			shuffled := ""
			if player.State.Playlist.IsShuffled() {
				shuffled = icons.PLAYLIST_SHUFFLED
			}
			fmt.Printf("Playing: %s (%d/%d) %s\n",
				player.State.Playlist.Name,
				player.State.PlaylistIndex+1,
				len(player.State.Playlist.Songs),
				shuffled,
			)
		}

		icon := icons.PLAYING

		if player.State.Paused {
			icon = icons.PAUSED
		}

		extra := " "
		switch player.State.Loop {
		case player.LOOP_TRACK:
			extra += utils.ColorWhite + icons.LOOPED
		case player.LOOP_PLAYLIST:
			extra += utils.ColorBlue + icons.LOOPED
		}

		fmt.Printf(" %s %s %sfrom %s%s%s\n",
			icon,
			utils.ColorGreen+result.Title,
			utils.ColorWhite,
			utils.ColorGreen+result.Uploader,
			extra,
			utils.ColorReset,
		)

		fmt.Printf("Volume: %s%.0f%%%s\n", utils.ColorGreen, player.State.Volume, utils.ColorReset)

		if player.State.ShowURL {
			fmt.Printf("%s%s%s\n", utils.ColorBlue, result.URL, utils.ColorReset)
		}

		if player.State.ShowHelp {
			fmt.Println("\n" + utils.ColorBlue + "Keybinds:")
			for _, bind := range keybind.ListBinds() {
				fmt.Printf("  %s: %s\n", bind.KeyName, bind.Description)
			}
		}

		if player.State.ShowLyric {
			fmt.Println(utils.ColorBlue)
			lyric := player.State.Lyric
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

	player.RegisterHooks(
		func(param ...interface{}) {
			renderPlayer()
		},
		player.HOOK_PLAYBACK_PAUSED, player.HOOK_PLAYBACK_RESUMED,
		player.HOOK_VOLUME_CHANGED, player.HOOK_POSITION_CHANGED,
		player.HOOK_GENERIC_UPDATE, player.HOOK_LOOP_PLAYLIST_CHANGED,
		player.HOOK_LOOP_TRACK_CHANGED, player.HOOK_FILE_LOAD_STARTED,
		player.HOOK_FILE_ENDED,
	)

	// progress bar updater
	player.RegisterHook(func(param ...interface{}) {
		renderProgressBar()
	}, player.HOOK_SEEK)
	go func() {
		for {
			if !player.State.Idle && !player.State.Paused {
				renderProgressBar()
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
