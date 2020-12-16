package new_player

const (
	HOOK_PLAYER_INITIALIZED = iota
	HOOK_RESULT_FETCH_STARTED
	HOOK_FILE_LOAD_STARTED
	HOOK_PLAYBACK_PAUSED
	HOOK_PLAYBACK_RESUMED
	HOOK_VOLUME_CHANGED
	HOOK_FILE_LOADED
	HOOK_POSITION_CHANGED
	HOOK_RESULT_DOWNLOAD_STARTED
	HOOK_FILE_ENDED
	HOOK_GENERIC_UPDATE
	HOOK_LOOP_TRACK_CHANGED
	HOOK_PLAYLIST_SONG_CHANGED
	HOOK_LOOP_PLAYLIST_CHANGED
	HOOK_IDLE
	HOOK_SAVING_TRACK_TO_PLAYLIST
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
