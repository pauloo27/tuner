package new_display

import (
	"fmt"

	"github.com/Pauloo27/tuner/new_player"
	"github.com/Pauloo27/tuner/utils"
)

func DisplayPlayer() {
	render := func() {
		// TODO: lock
		utils.ClearScreen()
		fmt.Println("TODO")
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
