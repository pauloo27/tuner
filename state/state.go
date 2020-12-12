package state

import (
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/storage"
)

var (
	Data        *storage.TunerData
	MPVInstance *player.MPV
	Playing     bool
	Warning     string
)

func Start() {
	Data = storage.Load()
	Playing = false
}
