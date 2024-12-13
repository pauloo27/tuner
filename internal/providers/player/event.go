package player

type PlayerEvent string

const (
	PlayerEventPlay  PlayerEvent = "PLAY"
	PlayerEventPause PlayerEvent = "PAUSE"
)

type PlayerEventCallback = func(...any)
