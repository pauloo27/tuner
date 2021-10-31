package style

import (
	"math"
	"strings"

	"github.com/Pauloo27/tuner/internal/ui/components/progress"
	"github.com/Pauloo27/tuner/internal/utils"
)

type SimpleBar struct {
	body []string
}

func NewSimpleBar(body ...string) SimpleBar {
	return SimpleBar{body}
}

func (r SimpleBar) Draw(p *progress.ProgressBar) {
	percentage := math.Min(1, float64(p.GetProgress()))
	x, _, w, _ := p.GetRect()

	bodyLen := len(r.body)
	lastBodyItem := r.body[bodyLen-1]

	lineWidth := (w - x) * bodyLen
	usedLine := int(math.Round(percentage * float64(lineWidth)))

	fullBlocks := strings.Repeat(lastBodyItem, usedLine/bodyLen)

	p.SetText(utils.Fmt("%s%s", fullBlocks, r.body[usedLine%bodyLen]))
}
