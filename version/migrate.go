package version

import (
	"strings"

	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/storage"
)

func Migrate(currentVersion string) {
	if currentVersion != player.State.Data.Version {
		switch strings.Split(currentVersion, "-")[0] {
		}
		player.State.Data.Version = currentVersion
		storage.Save(player.State.Data)
	}
}
