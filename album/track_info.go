package album

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
	"github.com/buger/jsonparser"
)

type ArtistInfo struct {
	Name, MBID string
}

type AlbumInfo struct {
	Title, MBID, ImageURL string
}

type TrackInfo struct {
	Title, MBID string
	Tags        []string
	Artist      *ArtistInfo
	Album       *AlbumInfo
}

const (
	APIKEY   = "12dec50313f885d407cf8132697b8712"
	EndPoint = "https://ws.audioscrobbler.com/2.0"
)

var (
	infoCache = make(map[string]*TrackInfo)
)

func FetchTrackInfo(artist, track string) (*TrackInfo, error) {
	// we use the hash to cache the track info
	// i know, md5 is shitty,
	hash := fmt.Sprintf("%x", md5.Sum([]byte(artist+"|"+track)))

	info, found := infoCache[hash]
	if found {
		return info, nil
	}

	// escape params
	artist = url.QueryEscape(artist)
	track = url.QueryEscape(track)

	reqPath := utils.Fmt(
		"%s/?method=track.getInfo&api_key=%s&artist=%s&track=%s&format=json",
		EndPoint, APIKEY, artist, track,
	)

	res, err := http.Get(reqPath)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buffer, err := ioutil.ReadAll(res.Body)

	// artist info
	artistName, err := jsonparser.GetString(buffer, "track", "artist", "name")
	if err != nil {
		return nil, fmt.Errorf("Cannot get artist name: %v", err)
	}

	artistMBID, err := jsonparser.GetString(buffer, "track", "artist", "mbid")
	if err != nil {
		artistMBID = "*" + artistName
	}

	// album info
	albumTitle, err := jsonparser.GetString(buffer, "track", "album", "title")
	if err != nil {
		return nil, fmt.Errorf("Cannot get album title: %v", err)
	}

	albumMBID, err := jsonparser.GetString(buffer, "track", "album", "mbid")
	if err != nil {
		albumMBID = artistMBID + "|" + albumTitle
	}

	albumImageURL, err := jsonparser.GetString(buffer, "track", "album", "image", "[3]", "#text")
	if err != nil {
		return nil, fmt.Errorf("Cannot get album image: %v", err)
	}

	// track info
	trackTitle, err := jsonparser.GetString(buffer, "track", "name")
	if err != nil {
		return nil, fmt.Errorf("Cannot get track title: %v", err)
	}

	trackMBID, err := jsonparser.GetString(buffer, "track", "mbid")
	if err != nil {
		trackMBID = albumMBID + "|" + trackTitle
	}

	var trackTags []string
	tagsArr, _, _, err := jsonparser.Get(buffer, "track", "toptags", "tag")

	_, err = jsonparser.ArrayEach(tagsArr, func(data []byte, t jsonparser.ValueType, i int, err error) {
		tagName, err := jsonparser.GetString(data, "name")
		trackTags = append(trackTags, tagName)
	})
	if err != nil {
		return nil, fmt.Errorf("Cannot get track tags: %v", err)
	}

	info = &TrackInfo{
		Title: trackTitle,
		MBID:  trackMBID,
		Tags:  trackTags,
		Album: &AlbumInfo{
			Title:    albumTitle,
			MBID:     albumMBID,
			ImageURL: albumImageURL,
		},
		Artist: &ArtistInfo{
			Name: artistName,
			MBID: artistMBID,
		},
	}
	infoCache[hash] = info
	return info, nil
}

func GetAlbumURL(result *search.SearchResult) string {
	var albumURL string

	if result.SourceName == "soundcloud" {
		albumURL = result.Extra[0]
	} else {
		videoInfo, err := FetchVideoInfo(result)
		if err != nil {
			return ""
		}
		albumURL = utils.Fmt("https://i1.ytimg.com/vi/%s/hqdefault.jpg", videoInfo.ID)
		if videoInfo.Artist != "" && videoInfo.Track != "" {
			trackInfo, err := FetchTrackInfo(videoInfo.Artist, videoInfo.Track)
			if err == nil {
				albumURL = trackInfo.Album.ImageURL
			}
		}
	}
	return albumURL
}
