package img

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

var stdin io.WriteCloser

func StartDaemon() {
	cmd := exec.Command("ueberzug", "layer", "--parser", "json")
	var err error
	stdin, err = cmd.StdinPipe()
	utils.HandleError(err, "Cannot connect to stdin")
	go func() {
		err := cmd.Run()
		if err != nil {
			player.State.Data.FetchAlbum = false
			storage.Save(player.State.Data)
			fmt.Printf("%sShow album disabled...%s\n", utils.ColorYellow, utils.ColorReset)
			utils.HandleError(err, "Cannot start ueberzug")
		}
	}()
}

func SendCommand(command string) {
	if stdin == nil {
		return
	}
	_, err := io.WriteString(stdin, command+"\n")
	utils.HandleError(err, "Cannot send command")
}
