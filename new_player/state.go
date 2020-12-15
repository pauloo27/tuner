package new_player

import "github.com/Pauloo27/tuner/search"

type PlayerState struct {
	Paused   bool
	playing  *search.YouTubeResult
	Volume   float64
	Duration float64
}

func (s *PlayerState) GetPlaying() *search.YouTubeResult {
	return s.playing
}
