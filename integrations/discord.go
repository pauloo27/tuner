package integrations

import (
	"github.com/Pauloo27/tuner/player"
	"github.com/ananagame/rich-go/client"
)

func setActivity(state, details string) {
	err := client.SetActivity(client.Activity{
		State:      state,
		Details:    details,
		LargeImage: "music",
		LargeText:  "Search and Play songs from YouTube inside your terminal",
	})
	if err != nil {
		return
	}
}

func ConnectToDiscord() {
	err := client.Login("827039629114867724")
	if err != nil {
		return
	}

	setActivity("Home screen", "Just started")

	player.RegisterHook(func(param ...interface{}) {
		setActivity("Home screen", "Idle")
	}, player.HOOK_IDLE)

	player.RegisterHook(func(param ...interface{}) {
		setActivity(player.State.GetPlaying().Title, "Playing")
	}, player.HOOK_FILE_LOADED)
}
