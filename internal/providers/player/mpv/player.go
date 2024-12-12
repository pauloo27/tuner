package mpv

import (
	"log/slog"

	"github.com/pauloo27/tuner/internal/providers/player"
	"github.com/pauloo27/tuner/internal/providers/player/mpv/libmpv"
	"github.com/pauloo27/tuner/internal/providers/source"
)

type MpvPlayer struct {
	Instance *libmpv.Mpv
}

var _ player.PlayerProvider = &MpvPlayer{}

func (p *MpvPlayer) Play(result source.SearchResult) error {
	info, err := result.GetAudioInfo()
	if err != nil {
		return err
	}
	p.Instance.Command([]string{"loadfile", info.StreamURL})
	return nil
}

func (*MpvPlayer) GetName() string {
	return "MPV"
}

func NewMpvPlayer() (*MpvPlayer, error) {
	instance := libmpv.Create()

	mustSetOption := func(name string, data string) {
		err := instance.SetOptionString(name, data)
		if err != nil {
			slog.Error("Failed to set option", "name", name, "data", data, "err", err)
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
