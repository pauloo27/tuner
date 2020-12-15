package new_player

import "github.com/Pauloo27/tuner/search"

type SongLyric struct {
	Lines []string
	Index int
}
type PlayerState struct {
	Paused                       bool
	playing                      *search.YouTubeResult
	Volume                       float64
	Duration                     float64
	ShowHelp, ShowURL, ShowLyric bool
	Lyric                        SongLyric
}

func (s *PlayerState) GetPlaying() *search.YouTubeResult {
	return s.playing
}
