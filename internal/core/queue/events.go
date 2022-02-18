package queue

import (
	"github.com/Pauloo27/tuner/internal/core/event"
	"github.com/Pauloo27/tuner/internal/core/track"
)

const (
	EventAppend  = event.EventType("APPEND")
	EventRemove  = event.EventType("REMOVE")
	EventShuffle = event.EventType("SHUFFLE")
)

type EventAppendData struct {
	Items  []track.Track
	Queue  *Queue
	Index  int
	IsMany bool
}

type EventRemoveData struct {
	Queue *Queue
	Index int
}
