package mpv

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

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
	err = p.instance.Command([]string{"loadfile", info.StreamURL})
	if err != nil {
		return err
	}

	return p.instance.SetPropertyString("force-media-title", result.Title)
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

	return isPaused.(bool), nil
}

func (p *MpvPlayer) GetVolume() (float64, error) {
	volume, err := p.instance.GetProperty("volume", libmpv.FORMAT_DOUBLE)
	if err != nil {
		return 0, err
	}

	return volume.(float64), nil
}

func (p *MpvPlayer) SetVolume(volume float64) error {
	return p.instance.SetProperty("volume", libmpv.FORMAT_DOUBLE, volume)
}

func (p *MpvPlayer) Stop() error {
	return p.instance.Command([]string{"stop"})
}

func (p *MpvPlayer) GetDuration() (time.Duration, error) {
	duration, err := p.instance.GetProperty("duration", libmpv.FORMAT_DOUBLE)
	if err != nil {
		return 0, err
	}

	return time.Duration(duration.(float64) * float64(time.Second)), nil
}

func (p *MpvPlayer) GetPosition() (time.Duration, error) {
	position, err := p.instance.GetProperty("time-pos", libmpv.FORMAT_DOUBLE)
	if err != nil {
		return 0, err
	}

	return time.Duration(position.(float64) * float64(time.Second)), nil
}

func (p *MpvPlayer) Seek(offset time.Duration) error {
	return p.instance.Command([]string{"seek", fmt.Sprintf("%.2f", offset.Seconds())})
}

func (p *MpvPlayer) IsErrPropertyUnavailable(err error) bool {
	return strings.Trim(err.Error(), "\n") == "-10 property unavailable"
}
