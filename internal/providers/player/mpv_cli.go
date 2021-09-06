package player

type MPVProvider struct {
}

var _ PlayerProvider = MPVProvider{}

func init() {
	playerProvider = MPVProvider{}
}

func (MPVProvider) GetName() string {
	return "MPV CLI"
}

func (MPVProvider) Play(url string) error {
	return nil
}
