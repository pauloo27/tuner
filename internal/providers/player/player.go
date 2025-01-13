package player

import (
	"time"

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
	GetVolume() (float64, error)
	SetVolume(float64) error
	Stop() error
	GetDuration() (time.Duration, error)
	GetPosition() (time.Duration, error)
}
