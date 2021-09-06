package player

// temporary thing since it's too early to write this
type PlayerProvider interface {
	GetName() string
	Play(url string) error
}

// TODO: what should it be?
var playerProvider PlayerProvider

func Play(url string) error {
	return playerProvider.Play(url)
}
