package debug

import "github.com/pauloo27/tuner/internal/ui/view"

type incMsg struct{}

func (incMsg) ForwardTo() view.ViewName {
	return view.DebugViewName
}
