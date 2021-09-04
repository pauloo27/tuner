package utils

import (
	"github.com/Pauloo27/tuner/ui"
	"github.com/rivo/tview"
)

func CreateContainer() *tview.Flex {
	container := tview.NewFlex()
	container.SetBackgroundColor(ui.GetTheme().PrimitiveBackgroundColor)
	return container
}
