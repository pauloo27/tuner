package new_player

import (
	"unsafe"

	"github.com/Pauloo27/tuner/new_player/mpv"
	"github.com/Pauloo27/tuner/utils"
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
				duration, err := MpvInstance.GetProperty("duration", mpv.FORMAT_DOUBLE)
				utils.HandleError(err, "Cannot get duration")
				State.Duration = duration.(float64)
				callHooks(HOOK_FILE_LOADED, err, duration)
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
