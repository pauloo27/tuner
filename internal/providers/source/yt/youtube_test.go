package yt_test

import (
	"testing"

	"github.com/pauloo27/tuner/internal/providers/source/yt"
	"github.com/stretchr/testify/assert"
)

var youtubeProvider yt.YouTubeSource

func TestYouTubeSearch(t *testing.T) {
	assert.Equal(t, "YouTube", youtubeProvider.Name())
	results, err := youtubeProvider.SearchFor("no copyright songs")
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.NotEmpty(t, results)
	assert.Len(t, results, 10)
}

func TestSearchResult(t *testing.T) {
	results, err := youtubeProvider.SearchFor("https://www.youtube.com/watch?v=yJg-Y5byMMw")
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.NotEmpty(t, results)
	assert.GreaterOrEqual(t, len(results), 1)
	result := results[0]
	assert.Equal(t, "NoCopyrightSounds", result.Artist)
	assert.Equal(t, "Warriyo - Mortals (feat. Laura Brehm) | Future Trap | NCS - Copyright Free Music", result.Title)
	assert.False(t, result.IsLive)
	assert.Equal(t, "3m50s", result.Length.String())
}
