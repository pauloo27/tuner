package keybind

import (
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/utils"
	"github.com/eiannone/keyboard"
)

func Listen() {
	err := keyboard.Open()
	utils.HandleError(err, "Cannot open keyboard")
	for {
		c, key, err := keyboard.GetKey()
		if err != nil {
			if !state.Playing {
				break
			}
		} else {
			HandlePress(c, key, state.MPVInstance)
		}
	}
}
