package source_test

import (
	"testing"

	"github.com/Pauloo27/tuner/internal/providers/source"
	"github.com/stretchr/testify/assert"
)

var youtubeProvider source.YouTubeSearch

func TestYouTubeSearch(t *testing.T) {
	assert.Equal(t, "YouTube", youtubeProvider.GetName())
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
	assert.Equal(t, "Warriyo - Mortals (feat. Laura Brehm) [NCS Release]", result.Title)
	assert.False(t, result.IsLive)
	assert.Equal(t, "3m50s", result.Length.String())
}
