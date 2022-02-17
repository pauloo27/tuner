package style

import (
	"fmt"

	"github.com/Pauloo27/tuner/internal/ui/components/progress"
)

type RawText struct{}

func NewRawText() RawText {
	return RawText{}
}

func (r RawText) Draw(p *progress.ProgressBar) {
	p.SetText(fmt.Sprintf("%0.f%%", p.GetProgress()*100))
}
