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

var lyricUrlRe = regexp.MustCompile(`https:\/\/genius.com/[^/]+-lyrics`)

func Fetch(path string) (lyric string, err error) {
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

			lyric += strings.TrimSpace(soup.HTMLParse(html).FullText()) + "\n"
		}
		lyric = strings.TrimSpace(lyric)
	} else {
		lyric = strings.TrimSpace(lyricDiv.FullText())
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
