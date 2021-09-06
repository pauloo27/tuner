package player

import "os/exec"

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
	cmd := exec.Command("mpv", url)
	return cmd.Run()
}
