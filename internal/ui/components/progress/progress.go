package progress

import (
	"github.com/rivo/tview"
)

type ProgressBar struct {
	*tview.TextView
	progress float64
	style    ProgressBarStyle
}

func NewProgressBar(style ProgressBarStyle) *ProgressBar {
	return &ProgressBar{
		style:    style,
		TextView: tview.NewTextView(),
		progress: 0,
	}
}

func (p *ProgressBar) SetRelativeProgress(duration, position float64) *ProgressBar {
	return p.SetProgress(position / duration)
}

func (p *ProgressBar) GetProgress() float64 {
	return p.progress
}

func (p *ProgressBar) SetProgress(percentage float64) *ProgressBar {
	if p.progress == percentage {
		return p
	}
	p.progress = percentage
	p.style.Draw(p)
	return p
}
