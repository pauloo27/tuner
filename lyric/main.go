package lyric

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/Pauloo27/tuner/utils"
	"github.com/anaskhan96/soup"
	"github.com/buger/jsonparser"
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

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	doc := soup.HTMLParse(string(body))

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
	path := fmt.Sprintf("https://searx.lnode.net/search?format=json&q=site:genius.com+%s", url.QueryEscape(query))

	res, err := http.Get(path)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	buffer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	for i := 0; i < 5; i++ {
		url, err := jsonparser.GetString(buffer, "results", utils.Fmt("[%d]", i), "pretty_url")
		if err != nil {
			return "", err
		}
		if lyricUrlRe.MatchString(url) {
			return url, nil
		}
	}

	return "", fmt.Errorf("No results found")
}
