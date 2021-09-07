package player

// temporary thing since it's too early to write this
type PlayerProvider interface {
	GetName() string
	Play(url string) error
}

var DefaultProvider PlayerProvider

func Play(url string) error {
	return DefaultProvider.Play(url)
}
