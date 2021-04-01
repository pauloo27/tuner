package integrations

import (
	"os"

	"github.com/Pauloo27/tuner/player"
	"github.com/ananagame/rich-go/client"
)

func setActivity(state, details string) {
	err := client.SetActivity(client.Activity{
		State:      state,
		Details:    details,
		LargeImage: "tune",
		LargeText:  "Tune",
	})
	if err != nil {
		return
	}
}

func ConnectToDiscord() {
	err := client.Login(os.Getenv("DISCORD_APP_ID"))
	if err != nil {
		return
	}

	player.RegisterHook(func(param ...interface{}) {
		setActivity("Home screen", "Idle")
	}, player.HOOK_IDLE)

	player.RegisterHook(func(param ...interface{}) {
		setActivity(player.State.GetPlaying().Title, "Playing")
	}, player.HOOK_FILE_LOADED)

	setActivity("Home screen", "Idle")
}
