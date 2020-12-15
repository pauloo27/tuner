package new_player

import "github.com/Pauloo27/tuner/search"

type PlayerState struct {
	Paused   bool
	Playing  *search.YouTubeResult
	Volume   float64
	Duration float64
}
