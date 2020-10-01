package lyric

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

type SongLyric struct {
	Lyric string
}

var lyricUrlRe = regexp.MustCompile(`https:\/\/genius.com/[^/]+-lyrics`)

func FormatArgument(str string) string {
	str = strings.ToLower(str)
	// TODO: Improve
	str = strings.ReplaceAll(str, " ", "-")
	str = strings.ReplaceAll(str, "/", "-")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, "ç", "c")
	str = strings.ReplaceAll(str, "ã", "a")
	str = strings.ReplaceAll(str, "é", "e")
	str = strings.ReplaceAll(str, "í", "i")
	str = strings.ReplaceAll(str, "ô", "o")
	str = strings.ReplaceAll(str, "ú", "u")
	str = strings.ReplaceAll(str, "ü", "u")

	return str
}

func FetchLyricFor(artist, song string) (SongLyric, error) {
	return Fetch(fmt.Sprintf("https://genius.com/%s-%s-lyrics", FormatArgument(artist), FormatArgument(song)))
}

func Fetch(path string) (lyric SongLyric, err error) {
	res, err := http.Get(path)
	if err != nil {
		return
	}

	if res.StatusCode == 404 {
		err = fmt.Errorf("Cannot find lyrics (status code 404)")
		return
	}

	defer res.Body.Close()

	bodyB, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	doc := soup.HTMLParse(string(bodyB))

	lyricDiv := doc.Find("div", "class", "lyrics")

	if lyricDiv.Error != nil {

		for _, div := range doc.FindAll("div", "class", "jgQsqn") {

			html := strings.ReplaceAll(div.HTML(), "<br/>", "<br/>\n")

			lyric.Lyric += soup.HTMLParse(html).FullText()
		}
	} else {
		lyric.Lyric = lyricDiv.FullText()
	}

	return
}

func SearchFor(query string) (string, error) {
	path := fmt.Sprintf("https://html.duckduckgo.com/html/?q=site:genius.com+%s", url.QueryEscape(query))

	res, err := soup.Get(path)
	if err != nil {
		return "", err
	}

	doc := soup.HTMLParse(res)
	if err != nil {
		return "", err
	}

	for _, result := range doc.FindAll("a", "class", "result__url") {
		r := fmt.Sprintf("https://%s", strings.TrimSpace(result.Text()))

		if lyricUrlRe.MatchString(r) {
			return r, nil
		}
	}

	return "", fmt.Errorf("No results found")
}
