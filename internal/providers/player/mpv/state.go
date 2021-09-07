package mpv

import "github.com/Pauloo27/tuner/internal/providers/search"

type SongLyric struct {
	Lines []string
	Index int
}

type LoopStatus int

const (
	StatusLoopNone = iota
	StatusLoopTrack
	StatusLoopPlaylist
)

type PlayerState struct {
	Paused   bool
	Idle     bool
	Result   *search.SearchResult
	Volume   float64
	Duration float64
	Loop     LoopStatus
}

func (s *PlayerState) GetPlaying() *search.SearchResult {
	return s.Result
}
