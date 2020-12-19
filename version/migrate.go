package version

import (
	"strings"

	"github.com/Pauloo27/tuner/player"
)

func Migrate(currentVersion string) {
	if currentVersion != player.State.Data.Version {
		switch strings.Split(currentVersion, "-")[0] {
		}
	}
}
