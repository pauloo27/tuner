package new_player

import (
	"fmt"
	"strings"

	"github.com/Pauloo27/tuner/lyric"
)

func FetchLyric() {
	playing := State.GetPlaying()
	path, err := lyric.SearchFor(fmt.Sprintf("%s %s", playing.Title, playing.Uploader))
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
