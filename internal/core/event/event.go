package event

type EventType string

type EventHandler func(params ...interface{})

type EventEmitter struct {
	events map[EventType][]*EventHandler
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		events: make(map[EventType][]*EventHandler),
	}
}

func (e *EventEmitter) Emit(eventType EventType, params ...interface{}) {
	listeners, found := e.events[eventType]
	if !found {
		return
	}
	for _, listener := range listeners {
		(*listener)(params...)
	}
}

func (e *EventEmitter) On(event EventType, listener EventHandler) {
	e.events[event] = append(e.events[event], &listener)
}
