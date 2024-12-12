package player

import (
	"github.com/pauloo27/tuner/internal/providers/source"
)

// temporary thing since it's too early to write this
type PlayerProvider interface {
	GetName() string
	Play(*source.SearchResult) error
}

var Player PlayerProvider
