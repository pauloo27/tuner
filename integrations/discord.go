package integrations

import (
	"github.com/Pauloo27/tuner/album"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/hugolgst/rich-go/client"
)

func setActivity(state, details, image string) {
	if image == "" {
		image = "music"
	}
	// TODO: handle error?
	_ = client.SetActivity(client.Activity{
		State:      state,
		Details:    details,
		LargeImage: image,
		LargeText:  "Play songs from YouTube inside your terminal",
	})
}

func getImage(result *search.SearchResult) string {
	return ""
}

func ConnectToDiscord() {
	err := client.Login("827039629114867724")
	if err != nil {
		return
	}

	setActivity("Home screen", "Just started", "")

	player.RegisterHook(func(param ...interface{}) {
		setActivity("Home screen", "Idle", "")
	}, player.HookIdle)

	player.RegisterHook(func(param ...interface{}) {
		setActivity(player.State.GetPlaying().Title, "Playing", album.GetAlbumURL(player.State.GetPlaying()))
	}, player.HookFileLoaded)
}
