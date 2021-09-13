package player

import (
	"github.com/Pauloo27/tuner/internal/providers/source"
)

// temporary thing since it's too early to write this
type PlayerProvider interface {
	GetName() string
	Play(*source.SearchResult) error
}

var DefaultProvider PlayerProvider

func Play(result *source.SearchResult) error {
	return DefaultProvider.Play(result)
}
