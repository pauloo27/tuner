package theme

import (
	"github.com/gdamore/tcell/v2"
)

type playingPageTheme struct {
	SongInfoColor tcell.Color
	VolumeColor   tcell.Color
}

var (
	PlayingPageTheme = playingPageTheme{
		SongInfoColor: tcell.ColorYellow,
		VolumeColor:   tcell.ColorYellow,
	}
)
