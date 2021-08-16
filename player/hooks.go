package player

const (
	HookPlayerInitialized = iota
	HookFileLoadStarted
	HookPlaybackPaused
	HookPlaybackResumed
	HookVolumeChanged
	HookFileLoaded
	HookFileEnded
	HookGenericUpdate
	HookLoopTrackChanged
	HookPlaylistSongChanged
	HookLoopPlaylistChanged
	HookIdle
	HookSavingTrackToPlaylist
	HookSeek
)

type HookCallback func(params ...interface{})

var hooks = make(map[int][]*HookCallback)

func RegisterHook(cb HookCallback, hookType int) {
	if currentHooks, ok := hooks[hookType]; ok {
		hooks[hookType] = append(currentHooks, &cb)
	} else {
		hooks[hookType] = []*HookCallback{&cb}
	}
}

func RegisterHooks(cb HookCallback, hookTypes ...int) {
	for _, hookType := range hookTypes {
		RegisterHook(cb, hookType)
	}
}

func callHooks(hookType int, params ...interface{}) {
	if hooks, ok := hooks[hookType]; ok {
		for _, hook := range hooks {
			(*hook)(params...)
		}
	}
}
