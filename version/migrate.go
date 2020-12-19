package version

import (
	"fmt"
	"strings"

	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/storage"
)

func Migrate(currentVersion string) {
	if currentVersion != player.State.Data.Version {
		switch strings.Split(currentVersion, "-")[0] {
		case "0.0.2":
			// create a URL from the id
			for _, playlist := range player.State.Data.Playlists {
				for i := 0; i < len(playlist.Songs); i++ {
					playlist.Songs[i].URL = fmt.Sprintf("https://youtube.com/watch?v=%s",
						playlist.Songs[i].ID,
					)
				}
			}
			storage.Save(player.State.Data)
		}
		player.State.Data.Version = currentVersion
		storage.Save(player.State.Data)
	}
}
