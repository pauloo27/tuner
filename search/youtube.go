package search

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Pauloo27/tuner/utils"
	"github.com/buger/jsonparser"
)

func SearchYouTube(searchTerm string) {
	url := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", url.QueryEscape(searchTerm))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Language", "en")
	res, err := client.Do(req)

	utils.HandleError(err, "Cannot get youtube page")

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	buffer, err := ioutil.ReadAll(res.Body)
	utils.HandleError(err, "Cannot read body")

	body := string(buffer)
	splittedScript := strings.Split(body, `window["ytInitialData"] = `)
	if len(splittedScript) != 2 {
		utils.HandleError(errors.New("Too much splitted scripts"), "Cannot split script")
	}
	splittedScript = strings.Split(splittedScript[1], `window["ytInitialPlayerResponse"] = null;`)
	jsonData := []byte(splittedScript[0])

	contents, _, _, _ := jsonparser.Get(jsonData, "contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents", "[0]", "itemSectionRenderer", "contents")
	jsonparser.ArrayEach(contents, func(value []byte, t jsonparser.ValueType, i int, err error) {
		id, err := jsonparser.GetString(value, "videoRenderer", "videoId")
		title, err := jsonparser.GetString(value, "videoRenderer", "title", "runs", "[0]", "text")

		fmt.Printf(" -> %s: %s\n", title, id)
	})
}
