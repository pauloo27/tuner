package mpv

import (
	"github.com/Pauloo27/logger"
	"github.com/Pauloo27/tuner/internal/providers/player"
)

func init() {
	err := Initialize()
	if err != nil {
		logger.Fatal(err)
	}
	player.DefaultProvider = &MpvProvider{}
}

type MpvProvider struct {
}

func (MpvProvider) GetName() string {
	return "MPV"
}

func (MpvProvider) Play(url string) error {
	if err := ClearPlaylist(); err != nil {
		return err
	}

	if err := LoadFile(url); err != nil {
		return err
	}

	return Play()
}
