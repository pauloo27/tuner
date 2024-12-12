package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	defaultTheme = &tview.Theme{
		PrimitiveBackgroundColor:    tcell.ColorDefault, // "transparent"
		ContrastBackgroundColor:     tcell.ColorBlue,
		MoreContrastBackgroundColor: tcell.ColorGreen,
		BorderColor:                 tcell.ColorWhite,
		TitleColor:                  tcell.ColorGreen,
		GraphicsColor:               tcell.ColorWhite,
		PrimaryTextColor:            tcell.ColorGreen,
		SecondaryTextColor:          tcell.ColorBlue,
		TertiaryTextColor:           tcell.ColorWhite,
		InverseTextColor:            tcell.ColorBlue,
		ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
	}
)
