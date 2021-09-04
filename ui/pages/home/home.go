package pages

import (
	"github.com/Pauloo27/tuner/ui"
	"github.com/Pauloo27/tuner/ui/utils"
)

func init() {
	container := utils.CreateContainer()
	ui.RegisterPage("home", container)
}
