package player

import (
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/storage"
)

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
	Data                         *storage.TunerData
	Paused                       bool
	Idle                         bool
	Result                       *search.SearchResult
	Playlist                     *storage.Playlist
	PlaylistIndex                int
	Volume                       float64
	Duration                     float64
	ShowHelp, ShowURL, ShowLyric bool
	Lyric                        SongLyric
	Loop                         LoopStatus
	SavingToPlaylist             bool
}

func (s *PlayerState) IsPlaylist() bool {
	return s.Playlist != nil
}

func (s *PlayerState) GetPlaying() *search.SearchResult {
	if s.IsPlaylist() {
		return s.Playlist.Songs[s.PlaylistIndex]
	}
	return s.Result
}
