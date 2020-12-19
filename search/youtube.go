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

func getContent(data []byte, index int) []byte {
	id := fmt.Sprintf("[%d]", index)
	contents, _, _, _ := jsonparser.Get(data, "contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents", id, "itemSectionRenderer", "contents")
	return contents
}

var httpClient = &http.Client{}

func SearchYouTube(searchTerm string, limit int) (results []*SearchResult) {
	url := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", url.QueryEscape(searchTerm))

	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err, "Cannot create GET request")
	req.Header.Add("Accept-Language", "en")
	res, err := httpClient.Do(req)
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
		splittedScript = strings.Split(body, `var ytInitialData = `)
	}

	if len(splittedScript) != 2 {
		utils.HandleError(errors.New("Too much splitted scripts"), "Cannot split script")
	}
	splittedScript = strings.Split(splittedScript[1], `window["ytInitialPlayerResponse"] = null;`)
	jsonData := []byte(splittedScript[0])

	index := 0
	var contents []byte

	for {
		contents = getContent(jsonData, index)
		_, _, _, err = jsonparser.Get(contents, "[0]", "carouselAdRenderer")

		if err == nil {
			index++
		} else {
			break
		}
	}

	_, err = jsonparser.ArrayEach(contents, func(value []byte, t jsonparser.ValueType, i int, err error) {
		utils.HandleError(err, "Cannot parse result contents")
		if limit > 0 && len(results) >= limit {
			return
		}

		id, err := jsonparser.GetString(value, "videoRenderer", "videoId")
		if err != nil {
			return
		}

		title, err := jsonparser.GetString(value, "videoRenderer", "title", "runs", "[0]", "text")
		if err != nil {
			return
		}

		uploader, err := jsonparser.GetString(value, "videoRenderer", "ownerText", "runs", "[0]", "text")
		if err != nil {
			return
		}

		live := false
		duration, err := jsonparser.GetString(value, "videoRenderer", "lengthText", "simpleText")

		if err != nil {
			duration = ""
			live = true
		}

		results = append(results, &SearchResult{
			Title:    title,
			Uploader: uploader,
			Duration: duration,
			ID:       id,
			URL:      fmt.Sprintf("https://youtube.com/watch?v=%s", id),
			Live:     live,
		})
	})

	if err != nil {
		utils.HandleError(err, "Cannot parse result")
	}

	return
}
