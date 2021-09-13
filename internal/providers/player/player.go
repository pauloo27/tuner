package player

import "github.com/Pauloo27/tuner/internal/providers/search"

// temporary thing since it's too early to write this
type PlayerProvider interface {
	GetName() string
	Play(*search.SearchResult) error
}

var DefaultProvider PlayerProvider

func Play(result *search.SearchResult) error {
	return DefaultProvider.Play(result)
}
