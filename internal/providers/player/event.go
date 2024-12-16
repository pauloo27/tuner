package player

type PlayerEvent string

const (
	PlayerEventPlay          PlayerEvent = "PLAY"
	PlayerEventPause         PlayerEvent = "PAUSE"
	PlayerEventVolumeChanged PlayerEvent = "VOLUME_CHANGED"
	PlayerEventIdle          PlayerEvent = "IDLE"
)

type PlayerEventCallback = func(...any)
