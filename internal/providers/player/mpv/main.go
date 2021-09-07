package mpv

import (
	"math"
	"strconv"

	"github.com/Pauloo27/logger"

	mpv "github.com/Pauloo27/tuner/internal/providers/player/mpv/libmpv"
)

const (
	maxVolume = 150.0
)

var MpvInstance *mpv.Mpv
var State *PlayerState

func Initialize() error {
	var err error

	// create a mpv instance
	MpvInstance = mpv.Create()

	initialVolume := 50.0

	// set options
	// disable video
	err = MpvInstance.SetOptionString("video", "no")
	if err != nil {
		return err
	}

	// set default volume value
	err = MpvInstance.SetOption("volume", mpv.FORMAT_DOUBLE, initialVolume)
	if err != nil {
		return err
	}

	// set default volume value
	err = MpvInstance.SetOption("volume-max", mpv.FORMAT_DOUBLE, maxVolume)
	if err != nil {
		return err
	}

	// set quality to worst
	err = MpvInstance.SetOptionString("ytdl-format", "worst")
	if err != nil {
		return err
	}

	// add observers
	err = MpvInstance.ObserveProperty(0, "volume", mpv.FORMAT_DOUBLE)
	if err != nil {
		return err
	}

	err = MpvInstance.ObserveProperty(0, "loop-file", mpv.FORMAT_FLAG)
	if err != nil {
		return err
	}

	err = MpvInstance.ObserveProperty(0, "loop-playlist", mpv.FORMAT_FLAG)
	if err != nil {
		return err
	}

	// TODO: start event listener
	//startEventHandler()

	// create the state
	State = &PlayerState{
		Paused:   false,
		Idle:     false,
		Result:   nil,
		Volume:   initialVolume,
		Duration: 0.0,
		Loop:     StatusLoopNone,
	}

	registerInternalHooks()

	// start the player
	err = MpvInstance.Initialize()
	if err != nil {
		return err
	}

	callHooks(HookPlayerInitialized)
	return nil
}

func registerInternalHooks() {
	RegisterHook(func(param ...interface{}) {
		if State.GetPlaying().IsLive {
			return
		}
		duration, err := MpvInstance.GetProperty("duration", mpv.FORMAT_DOUBLE)
		if err != nil {
			logger.Errorf("Cannot get duration: %v", err)
		}

		State.Duration = duration.(float64)
	}, HookFileLoaded)

	RegisterHook(func(params ...interface{}) {
		State.Volume = params[0].(float64)
	}, HookVolumeChanged)

	RegisterHook(func(params ...interface{}) {
		if State.Loop == StatusLoopNone {
			State.Loop = StatusLoopTrack
			return
		}
		State.Loop = StatusLoopNone
	}, HookLoopTrackChanged)

	RegisterHook(func(params ...interface{}) {
		if State.Loop == StatusLoopNone {
			State.Loop = StatusLoopPlaylist
		}
		State.Loop = StatusLoopNone
	}, HookLoopPlaylistChanged)

	RegisterHook(func(param ...interface{}) {
		State.Idle = true
	}, HookIdle)

	RegisterHook(func(param ...interface{}) {
		State.Idle = false
	}, HookFileLoadStarted)
}

func ClearPlaylist() error {
	return MpvInstance.Command([]string{"playlist-clear"})
}

func RemoveCurrentFromPlaylist() error {
	return MpvInstance.Command([]string{"playlist-remove", "current"})
}

func Stop() error {
	return MpvInstance.Command([]string{"stop"})
}

func PlaylistNext() error {
	return MpvInstance.Command([]string{"playlist-next"})
}

func PlaylistPrevious() error {
	return MpvInstance.Command([]string{"playlist-prev"})
}

func ForceUpdate() {
	callHooks(HookGenericUpdate)
}

func LoadFile(filePath string) error {
	err := MpvInstance.Command([]string{"loadfile", filePath})
	callHooks(HookFileLoadStarted, err, filePath)
	return err
}

func AppendFile(filePath string) error {
	err := MpvInstance.Command([]string{"loadfile", filePath, "append"})
	callHooks(HookFileLoadStarted, err, filePath)
	return err
}

func PlayPause() error {
	if State.Paused {
		return Play()
	}
	return Pause()
}

func Seek(seconds int) error {
	return MpvInstance.Command([]string{"seek", strconv.Itoa(seconds)})
}

func Pause() error {
	return MpvInstance.SetProperty("pause", mpv.FORMAT_FLAG, true)
}

func Play() error {
	return MpvInstance.SetProperty("pause", mpv.FORMAT_FLAG, false)
}

func LoopNone() error {
	err := MpvInstance.SetProperty("loop-file", mpv.FORMAT_FLAG, false)
	if err != nil {
		return err
	}
	err = MpvInstance.SetProperty("loop-playlist", mpv.FORMAT_FLAG, false)
	return err
}

func LoopTrack() error {
	err := MpvInstance.SetProperty("loop-file", mpv.FORMAT_FLAG, true)
	if err != nil {
		return err
	}
	err = MpvInstance.SetProperty("loop-playlist", mpv.FORMAT_FLAG, false)
	return err
}

func LoopPlaylist() error {
	err := MpvInstance.SetProperty("loop-file", mpv.FORMAT_FLAG, false)
	if err != nil {
		return err
	}
	err = MpvInstance.SetProperty("loop-playlist", mpv.FORMAT_FLAG, true)
	return err
}

func SetVolume(volume float64) error {
	volume = math.Min(maxVolume, volume)
	err := MpvInstance.SetProperty("volume", mpv.FORMAT_DOUBLE, volume)
	return err
}

func GetPosition() (float64, error) {
	position, err := MpvInstance.GetProperty("time-pos", mpv.FORMAT_DOUBLE)
	if err != nil {
		return 0.0, err
	}
	return position.(float64), err
}

func SetPosition(pos float64) error {
	return MpvInstance.SetProperty("time-pos", mpv.FORMAT_DOUBLE, pos)
}
