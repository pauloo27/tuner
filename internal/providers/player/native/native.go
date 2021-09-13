package native

import (
	"os"
	"time"

	"github.com/Pauloo27/tuner/internal/providers/player"
	"github.com/Pauloo27/tuner/internal/providers/search"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func init() {
	player.DefaultProvider = &NativeProvider{}
}

type NativeProvider struct {
}

var _ player.PlayerProvider = NativeProvider{}

func (NativeProvider) GetName() string {
	return "Native"
}

func (NativeProvider) Play(result *search.SearchResult) error {
	// TODO: get the link to the oppus file
	f, err := os.Open("test.mp3")
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}

	speaker.Play(streamer)
	select {}
}
