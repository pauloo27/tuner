package player

import (
	"unsafe"

	"github.com/Pauloo27/tuner/player/mpv"
)

func handlePropertyChange(data *mpv.EventProperty) {
	switch data.Name {
	case "volume":
		volume := *(*float64)(data.Data.(unsafe.Pointer))
		callHooks(HookVolumeChanged, volume)
	case "loop-file":
		callHooks(HookLoopTrackChanged)
	case "loop-playlist":
		callHooks(HookLoopPlaylistChanged)
	}
}

func startEventHandler() {
	go func() {
		for {
			event := MpvInstance.WaitEvent(60)
			switch event.Event_Id {
			case mpv.EVENT_NONE:
				continue
			case mpv.EVENT_PROPERTY_CHANGE:
				data := event.Data.(*mpv.EventProperty)
				handlePropertyChange(data)
			case mpv.EVENT_FILE_LOADED:
				callHooks(HookFileLoaded)
			case mpv.EVENT_PAUSE:
				State.Paused = true
				callHooks(HookPlaybackPaused)
			case mpv.EVENT_UNPAUSE:
				State.Paused = false
				callHooks(HookPlaybackResumed)
			case mpv.EVENT_END_FILE:
				callHooks(HookFileEnded)
			case mpv.EVENT_IDLE:
				callHooks(HookIdle)
			case mpv.EVENT_SEEK:
				callHooks(HookSeek)
			}
		}
	}()
}
