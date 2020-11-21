package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Pauloo27/go-mpris"
	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/commands"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/options"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
	"github.com/eiannone/keyboard"
)

var close = make(chan os.Signal)
var playing = false
var mpvInstance *player.MPV
var data *storage.TunerData

const (
	pausedIcon  = ""
	playingIcon = ""
)

func doSearch(searchTerm string, limit int) (results []search.YouTubeResult, err error) {
	c := make(chan bool)

	go utils.PrintWithLoadIcon(fmt.Sprintf("Searching for %s", searchTerm), c, 100*time.Millisecond, true)
	results = search.SearchYouTube(searchTerm, limit)

	c <- true
	<-c
	return
}

func setupCloseHandler() {
	signal.Notify(close, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			<-close
			if !playing {
				utils.MoveCursorTo(1, 1)
				utils.ClearScreen()
				fmt.Println("Bye!")
				os.Exit(0)
			}
		}
	}()
}

func listResults(results []search.YouTubeResult) {
	utils.ClearScreen()
	for i, result := range results {
		bold := ""
		if i%2 == 0 {
			bold = utils.ColorBold
		}

		defaultColor := bold + utils.ColorWhite
		altColor := bold + utils.ColorGreen

		duration := result.Duration

		if duration == "" {
			duration = utils.ColorRed + "LIVE"
		}

		fmt.Printf("  %s%d: %s %sfrom %s - %s%s\n",
			defaultColor, i+1,
			altColor+result.Title,
			defaultColor,
			altColor+result.Uploader,
			defaultColor+duration,
			utils.ColorReset,
		)
	}
}

func listenToKeyboard(mpv *player.MPV) {
	err := keyboard.Open()
	utils.HandleError(err, "Cannot open keyboard")
	for {
		c, key, err := keyboard.GetKey()
		if err != nil {
			if !playing {
				break
			}
		} else {
			keybind.HandlePress(c, key, mpv)
		}
	}
}

func saveToPlaylist(result *search.YouTubeResult, mpv *player.MPV) {
	keyboard.Close()
	utils.ClearScreen()

	fmt.Printf("Save to:\n")
	for i, playlist := range data.Playlists {
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

			if err == nil && index <= int64(len(data.Playlists)) && index > 0 {
				index--
				data.Playlists[index].Songs = append(data.Playlists[index].Songs, result)
				storage.Save(data)
			}
		} else {
			newPlaylist := &storage.Playlist{Name: rawPlaylist, Songs: []*search.YouTubeResult{result}}
			data.Playlists = append(data.Playlists, newPlaylist)
			storage.Save(data)
		}
	}

	// restore keyboard
	mpv.Saving = false
	go listenToKeyboard(mpv)
	mpv.Update()
}

func showPlayingScreen(result *search.YouTubeResult, mpv *player.MPV) {
	if !playing || mpv.Saving {
		return
	}

	utils.ClearScreen()

	icon := playingIcon

	playback, _ := mpv.Player.GetPlaybackStatus()
	if playback != mpris.PlaybackPlaying {
		icon = pausedIcon
	}

	extra := utils.ColorWhite
	if status, err := mpv.Player.GetLoopStatus(); err == nil && status == mpris.LoopTrack {
		extra += "  "
	}

	if result == nil {
		result = mpv.Playlist.Songs[mpv.PlaylistIndex]
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

func play(result *search.YouTubeResult, playlist *storage.Playlist) {
	parameters := []string{}
	if result == nil {
		for _, song := range playlist.Songs {
			parameters = append(parameters, song.URL())
		}
	} else {
		parameters = []string{result.URL()}
	}

	if !options.Options.ShowVideo {
		parameters = append(parameters, "--no-video", "--ytdl-format=worst")
	}
	if !options.Options.Cache {
		parameters = append(parameters, "--cache=no")
	}

	playing = true
	cmd := exec.Command("mpv", parameters...)

	go func() {
		if mpvInstance != nil && !mpvInstance.Exitted {
			mpvInstance.Exit()
		}
		mpvInstance = player.ConnectToMPV(cmd, result, playlist, showPlayingScreen, saveToPlaylist)
		go listenToKeyboard(mpvInstance)
	}()

	err := cmd.Run()

	if err != nil && err.Error() != "exit status 4" && err.Error() != "signal: killed" {
		utils.HandleError(err, "Cannot run MPV")
	}

	keyboard.Close()
	playing = false
	mpvInstance.Exit()
}

func findAndPlay(warning *string) {
	utils.ClearScreen()

	fmt.Printf("%sPlaylists:\n", utils.ColorBlue)
	for i, playlist := range data.Playlists {
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
	fmt.Printf("%sUse #<id> to start a playlist%s\n\n", utils.ColorBlue, utils.ColorReset)

	if *warning != "" {
		fmt.Printf("%s%s%s\n", utils.ColorYellow, *warning, utils.ColorReset)
		*warning = ""
	}

	rawSearchTerm, err := utils.AskFor("Search for")
	utils.HandleError(err, "Cannot read user input")

	if strings.HasPrefix(rawSearchTerm, "#") {
		index, err := strconv.ParseInt(strings.TrimPrefix(rawSearchTerm, "#"), 10, 64)

		if err == nil && index <= int64(len(data.Playlists)) && index > 0 {
			index--
			play(nil, data.Playlists[index])
		} else {
			*warning = "Invalid playlist"
		}
		return
	} else if strings.HasPrefix(rawSearchTerm, "/") {
		found, out := command.InvokeCommand(strings.TrimPrefix(rawSearchTerm, "/"))
		if !found {
			*warning = "Invalid command"
		} else {
			*warning = out
		}
		return
	}

	limit := 10
	if strings.HasPrefix(rawSearchTerm, "!") {
		limit = 1
	}
	searchTerm := strings.TrimPrefix(rawSearchTerm, "!")

	results, err := doSearch(searchTerm, limit)
	utils.HandleError(err, "Cannot search")

	if len(results) == 0 {
		*warning = "No results found"
		return
	}

	var entry search.YouTubeResult

	if len(results) == 1 {
		entry = results[0]
	} else {
		listResults(results)
		enteredIndex, err := utils.AskFor("Insert index of the video")
		utils.HandleError(err, "Cannot read user input")

		index, err := strconv.ParseInt(enteredIndex, 10, 64)

		if err != nil || index > int64(len(results)) || index <= 0 {
			*warning = "Invalid index"
			return
		}
		index--

		entry = results[index]
	}

	play(&entry, nil)
}

func main() {
	data = storage.Load()
	commands.SetupDefaultCommands()
	keybind.RegisterDefaultKeybinds(data)
	setupCloseHandler()
	warning := ""
	for {
		findAndPlay(&warning)
	}
}
