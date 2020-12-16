package state

import (
	"github.com/Pauloo27/tuner/storage"
)

var (
	Data    *storage.TunerData
	Warning string
)

func Start() {
	Data = storage.Load()
}
