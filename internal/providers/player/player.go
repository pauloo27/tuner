package player

import (
	"github.com/pauloo27/tuner/internal/providers/source"
)

type PlayerProvider interface {
	Name() string
	Play(source.SearchResult) error
	On(event PlayerEvent, handler PlayerEventCallback)
	Pause() error
	UnPause() error
	TogglePause() error
	IsPaused() (bool, error)
}
