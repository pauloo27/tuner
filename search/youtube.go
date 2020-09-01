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

type YouTubeResult struct {
	Title, Uploader, Duration, ID string
	Live                          bool
}

func SearchYouTube(searchTerm string, limit int) (results []YouTubeResult) {
	url := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", url.QueryEscape(searchTerm))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err, "Cannot create GET request")
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

		results = append(results, YouTubeResult{
			Title:    title,
			Uploader: uploader,
			Duration: duration,
			ID:       id,
			Live:     live,
		})
	})

	return
}
