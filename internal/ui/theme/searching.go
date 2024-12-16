package theme

import (
	"github.com/gdamore/tcell/v2"
)

type searchingPageTheme struct {
	Title     tcell.Color
	ItemColor tcell.Color
}

var (
	SearchingPageTheme = searchingPageTheme{
		Title:     tcell.ColorGreen,
		ItemColor: tcell.ColorBlue,
	}
)
