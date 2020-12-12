package display

import (
	"fmt"

	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

func ListPlaylists(playlists []*storage.Playlist) {
	for i, playlist := range playlists {
		bold := ""
		if i%2 == 0 {
			bold = utils.ColorBold
		}

		defaultColor := bold + utils.ColorWhite
		altColor := bold + utils.ColorGreen

		fmt.Printf("  %s#%d: %s%s\n",
			defaultColor, i+1,
			altColor+playlist.Name,
			utils.ColorReset,
		)
	}
}
