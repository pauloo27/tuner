package style

import (
	"github.com/Pauloo27/tuner/internal/ui/components/progress"
	"github.com/Pauloo27/tuner/internal/utils"
)

type RawText struct{}

func NewRawText() RawText {
	return RawText{}
}

func (r RawText) Draw(p *progress.ProgressBar) {
	p.SetText(utils.Fmt("%0.f%%", p.GetProgress()*100))
}
