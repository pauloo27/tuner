package display

import (
	"fmt"

	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
)

func ListResults(results []*search.SearchResult) {
	for i, result := range results {
		bold := ""
		if i%2 == 0 {
			bold = utils.ColorBold
		}
		defaultColor := bold + utils.ColorWhite
		altColor := bold + utils.ColorGreen

		duration := result.Duration

		if duration == "" {
			duration = utils.ColorRed + "LIVE"
		}

		fmt.Printf("  %s%d: %s %sfrom %s - %s%s\n",
			defaultColor, i+1,
			altColor+result.Title,
			defaultColor,
			altColor+result.Uploader,
			defaultColor+duration,
			utils.ColorReset,
		)
	}
}
