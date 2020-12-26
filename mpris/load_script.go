package mpris

import (
	"fmt"
	"os"

	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

var (
	USER_HOME_PATHS = []string{"/.config/mpv/scripts/mpris.so"}
	SYSTEM_PATHS    = []string{"/etc/mpv/scripts/mpris.so"}
)

func LoadScript() {
	if !player.State.Data.LoadMPRIS {
		return
	}

	loadScript := func(path string) bool {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return false
		}

		err := player.MpvInstance.Command([]string{"load-script", path})
		if err != nil {
			player.State.Data.LoadMPRIS = false
			storage.Save(player.State.Data)
			fmt.Printf("%sLoad MPRIS disabled...%s\n", utils.ColorYellow, utils.ColorReset)
			utils.HandleError(err, "Cannot load mpris script at "+path)
		}
		return true
	}

	home := utils.GetUserHome()
	for _, path := range USER_HOME_PATHS {
		if loadScript(home + path) {
			return
		}
	}

	for _, path := range SYSTEM_PATHS {
		if loadScript(path) {
			return
		}
	}
}
