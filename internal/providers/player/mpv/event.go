package mpv

import (
	"unsafe"

	"github.com/pauloo27/libmpv"
	"github.com/pauloo27/tuner/internal/providers/player"
)

func (p *MpvPlayer) listenToEvents() error {
	err := p.instance.ObserveProperty(0, "pause", libmpv.FORMAT_FLAG)
	if err != nil {
		return err
	}

	for {
		event := p.instance.WaitEvent(60)
		p.logger.Debug("Got event!", "event", event)
		switch event.Event_Id {
		case libmpv.EVENT_NONE:
			continue
		case libmpv.EVENT_PROPERTY_CHANGE:
			data := event.Data.(*libmpv.EventProperty)
			p.handlePropertyChange(data)
		}
	}
}

func (p *MpvPlayer) handlePropertyChange(data *libmpv.EventProperty) {
	switch data.Name {
	case "pause":
		value := *(*int)(data.Data.(unsafe.Pointer))
		if value == 0 {
			p.logger.Info("Play event!")
			p.Emit(player.PlayerEventPlay)
		} else {
			p.logger.Info("Pause event!")
			p.Emit(player.PlayerEventPause)
		}
	}
}
