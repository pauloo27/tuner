package mpv

import (
	"github.com/Pauloo27/tuner/internal/providers/player"
	"github.com/Pauloo27/tuner/internal/providers/player/mpv/libmpv"
	"github.com/Pauloo27/tuner/internal/providers/source"
)

// TODO:
type MpvPlayer struct {
	Instance *libmpv.Mpv
}

var _ player.PlayerProvider = MpvPlayer{}

func (p MpvPlayer) Play(result *source.SearchResult) error {
	p.Instance.Command([]string{"loadfile", result.URL})
	return nil
}

func (MpvPlayer) GetName() string {
	return "MPV"
}

func newMpvPlayer() (*MpvPlayer, error) {
	instance := libmpv.Create()

	mustSetOption := func(name string, data string) {
		err := instance.SetOptionString(name, data)
		if err != nil {
			panic(err)
		}
	}

	mustSetOption("video", "no")
	mustSetOption("cache", "no")

	err := instance.Initialize()

	return &MpvPlayer{
		instance,
	}, err
}

func init() {
	mpv, err := newMpvPlayer()
	if err != nil {
		panic(err)
	}
	player.DefaultProvider = mpv
}
