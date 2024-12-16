package theme

import (
	"github.com/gdamore/tcell/v2"
)

type homePageTheme struct {
	TitleText       tcell.Color
	SearchLabelText tcell.Color
	SearchField     FgBg
}

var (
	HomePageTheme = homePageTheme{
		TitleText:       tcell.ColorGreen,
		SearchLabelText: tcell.ColorGreen,
		SearchField:     FgBg{tcell.ColorWhite, tcell.ColorDefault},
	}
)
