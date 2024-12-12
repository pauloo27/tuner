package player

import (
	"github.com/pauloo27/tuner/internal/providers/source"
)

type PlayerProvider interface {
	GetName() string
	Play(source.SearchResult) error
}
