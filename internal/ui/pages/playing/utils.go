package playing

import (
	"fmt"
	"log/slog"

	"github.com/pauloo27/tuner/internal/providers"
)

const (
	volumeMax  = 100
	volumeMin  = 0
	volumeStep = 1
)

func (p *playingPage) buildSongLabel(icon string) string {
	return fmt.Sprintf("%s | %s from %s", icon, p.result.Title, p.result.Artist)
}

func (p *playingPage) updateVolumeLabel() error {
	volume, err := providers.Player.GetVolume()
	if err != nil {
		return err
	}
	p.volumeLabel.SetText(fmt.Sprintf("Volume: %.0f%%", volume))
	return nil
}

func (p *playingPage) incrementVolume() error {
	curVolume, err := providers.Player.GetVolume()
	if err != nil {
		return fmt.Errorf("failed to get current volume: %w", err)
	}

	newVolume := curVolume + volumeStep
	newVolume = min(volumeMax, newVolume)

	return providers.Player.SetVolume(newVolume)
}

func (p *playingPage) decrementVolume() error {
	curVolume, err := providers.Player.GetVolume()
	if err != nil {
		return fmt.Errorf("failed to get current volume: %w", err)
	}
	slog.Info("curVolume", "a", curVolume)

	newVolume := curVolume - volumeStep
	newVolume = max(volumeMin, newVolume)

	return providers.Player.SetVolume(newVolume)
}
