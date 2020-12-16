package new_player

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
	Paused                       bool
	playing                      *search.YouTubeResult
	playlist                     *storage.Playlist
	Volume                       float64
	Duration                     float64
	ShowHelp, ShowURL, ShowLyric bool
	Lyric                        SongLyric
	Loop                         LoopStatus
}

func (s *PlayerState) IsPlaylist() bool {
	return s.playing == nil
}

func (s *PlayerState) GetPlaying() *search.YouTubeResult {
	if s.IsPlaylist() {
		return s.playlist.Songs[0]
	}
	return s.playing
}
