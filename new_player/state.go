package new_player

import "github.com/Pauloo27/tuner/search"

type SongLyric struct {
	Lines []string
	Index int
}

type LoopStatus int

const (
	LOOP_NONE = iota
	LOOP_TRACK
	LOOP_PLAYLIST
)

type PlayerState struct {
	Paused                       bool
	playing                      *search.YouTubeResult
	Volume                       float64
	Duration                     float64
	ShowHelp, ShowURL, ShowLyric bool
	Lyric                        SongLyric
	Loop                         LoopStatus
}

func (s *PlayerState) GetPlaying() *search.YouTubeResult {
	return s.playing
}
