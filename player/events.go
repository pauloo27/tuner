package player

import (
	"unsafe"

	"github.com/Pauloo27/tuner/player/mpv"
)

func handlePropertyChange(data *mpv.EventProperty) {
	switch data.Name {
	case "volume":
		volume := *(*float64)(data.Data.(unsafe.Pointer))
		callHooks(HOOK_VOLUME_CHANGED, volume)
	case "loop-file":
		callHooks(HOOK_LOOP_TRACK_CHANGED)
	case "loop-playlist":
		callHooks(HOOK_LOOP_PLAYLIST_CHANGED)
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
				callHooks(HOOK_FILE_LOADED)
			case mpv.EVENT_PAUSE:
				State.Paused = true
				callHooks(HOOK_PLAYBACK_PAUSED)
			case mpv.EVENT_UNPAUSE:
				State.Paused = false
				callHooks(HOOK_PLAYBACK_RESUMED)
			case mpv.EVENT_END_FILE:
				callHooks(HOOK_FILE_ENDED)
			case mpv.EVENT_IDLE:
				callHooks(HOOK_IDLE)
			}
		}
	}()
}
