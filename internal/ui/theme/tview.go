package theme

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	dullTheme = tview.Theme{
		PrimitiveBackgroundColor:    tcell.ColorDefault, // "transparent"
		ContrastBackgroundColor:     tcell.ColorGray,
		MoreContrastBackgroundColor: tcell.ColorBlack,
		BorderColor:                 tcell.ColorGrey,
		TitleColor:                  tcell.ColorWhite,
		GraphicsColor:               tcell.ColorWhite,
		PrimaryTextColor:            tcell.ColorWhite,
		SecondaryTextColor:          tcell.ColorWhite,
		TertiaryTextColor:           tcell.ColorWhite,
		InverseTextColor:            tcell.ColorWhite,
		ContrastSecondaryTextColor:  tcell.ColorWhite,
	}
)

func ResetTviewTheme() {
	tview.Styles = dullTheme
}
