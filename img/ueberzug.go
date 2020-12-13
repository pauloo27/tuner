package img

import (
	"fmt"
	"io"
	"os/exec"

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
			utils.HandleError(err, "Cannot start ueberzug")
		} else {
			fmt.Println("done")
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
