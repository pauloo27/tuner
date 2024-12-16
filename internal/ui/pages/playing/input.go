package playing

import (
	"fmt"
	"log/slog"

	"github.com/pauloo27/tuner/internal/providers"
)

type InputHandler struct {
	Name string
	Fn   func() error
}

func (p *playingPage) registerInputHandlers() {
	p.inputHandler = map[rune]InputHandler{
		' ': {Name: "Play/Pause", Fn: providers.Player.TogglePause},
		'+': {Name: "Increment volume", Fn: incrementVolume},
		'=': {Name: "Increment volume", Fn: incrementVolume},
		'-': {Name: "Decrement volume", Fn: decrementVolume},
		'c': {Name: "Stop", Fn: stopPlaylist},
	}
}

func (p *playingPage) handleInput(key rune) {
	handler, found := p.inputHandler[key]
	if !found {
		slog.Debug("Unhandled input", "key", key)
		return
	}

	err := handler.Fn()
	if err != nil {
		slog.Error("Failed to handle input", "input", key, "handler", handler.Name, "err", err)
	}
}

func incrementVolume() error {
	curVolume, err := providers.Player.GetVolume()
	if err != nil {
		return fmt.Errorf("failed to get current volume: %w", err)
	}

	newVolume := curVolume + volumeStep
	newVolume = min(volumeMax, newVolume)

	return providers.Player.SetVolume(newVolume)
}

func decrementVolume() error {
	curVolume, err := providers.Player.GetVolume()
	if err != nil {
		return fmt.Errorf("failed to get current volume: %w", err)
	}
	slog.Info("curVolume", "a", curVolume)

	newVolume := curVolume - volumeStep
	newVolume = max(volumeMin, newVolume)

	return providers.Player.SetVolume(newVolume)
}

func stopPlaylist() error {
	return providers.Player.Stop()
}
