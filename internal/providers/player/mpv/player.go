package mpv

import (
	"log/slog"
	"os"

	"github.com/pauloo27/libmpv"
	"github.com/pauloo27/tuner/internal/core/event"
	"github.com/pauloo27/tuner/internal/providers/player"
	"github.com/pauloo27/tuner/internal/providers/source"
)

type MpvPlayer struct {
	instance *libmpv.Mpv
	logger   *slog.Logger
	*event.EventEmitter[player.PlayerEvent]
}

var _ player.PlayerProvider = &MpvPlayer{}

func (p *MpvPlayer) Play(result source.SearchResult) error {
	info, err := result.GetAudioInfo()
	if err != nil {
		return err
	}
	p.logger.Info("Loading file", "url", info.StreamURL, "format", info.Format)
	p.instance.Command([]string{"loadfile", info.StreamURL})
	return nil
}

func (*MpvPlayer) Name() string {
	return "MPV"
}

func NewMpvPlayer() (*MpvPlayer, error) {
	instance := libmpv.Create()

	mustSetOption := func(name string, data string) {
		err := instance.SetOptionString(name, data)
		if err != nil {
			slog.Error("Failed to set option", "name", name, "data", data, "err", err)
			os.Exit(1)
		}
	}

	mustSetOption("video", "no")
	mustSetOption("cache", "no")

	err := instance.Initialize()
	if err != nil {
		return nil, err
	}

	p := &MpvPlayer{
		instance,
		slog.With("player", "mpv"),
		event.NewEventEmitter[player.PlayerEvent](),
	}

	go func() {
		// ignore errors for now
		_ = p.listenToEvents()
	}()

	if err := p.loadMpris(); err != nil {
		p.logger.Warn("Failed to load mpris", "err", err)
	}

	return p, nil
}

func (p *MpvPlayer) Pause() error {
	return p.instance.SetProperty("pause", libmpv.FORMAT_FLAG, true)
}

func (p *MpvPlayer) UnPause() error {
	return p.instance.SetProperty("pause", libmpv.FORMAT_FLAG, false)
}

func (p *MpvPlayer) TogglePause() error {
	isPaused, err := p.IsPaused()
	if err != nil {
		return err
	}

	if isPaused {
		return p.UnPause()
	}
	return p.Pause()
}

func (p *MpvPlayer) IsPaused() (bool, error) {
	isPaused, err := p.instance.GetProperty("pause", libmpv.FORMAT_FLAG)
	if err != nil {
		return false, err
	}

	return isPaused.(bool), err
}
