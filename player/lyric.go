package player

import (
	"fmt"
	"strings"

	"github.com/Pauloo27/lyric"
)

func FetchLyric() {
	playing := State.GetPlaying()
	path, err := lyric.SearchDDG(fmt.Sprintf("%s %s", playing.Title, playing.Uploader))
	if err != nil {
		State.Lyric.Lines = []string{"Cannot get lyric"}
		ForceUpdate()
		return
	}

	l, err := lyric.Fetch(path)
	if err != nil {
		State.Lyric.Lines = []string{"Cannot get lyric"}
		ForceUpdate()
		return
	}

	State.Lyric.Lines = strings.Split(l, "\n")
	ForceUpdate()
}
