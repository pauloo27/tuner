package playing

import (
	"log/slog"

	"github.com/pauloo27/tuner/internal/providers"
)

func (p *playingPage) handleInput(key rune) {
	switch key {
	case ' ':
		err := providers.Player.TogglePause()
		if err != nil {
			slog.Error("Failed to toggle pause", "err", err)
		}
	case '+':
		err := p.incrementVolume()
		if err != nil {
			slog.Error("Failed to increment volume", "err", err)
		}
	case '-':
		err := p.decrementVolume()
		if err != nil {
			slog.Error("Failed to decrement volume", "err", err)
		}
	default:
		slog.Info("Unhandled input", "key", key)
	}
}
