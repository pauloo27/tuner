package event

type EventHandler = func(...any)

type EventEmitter[EventType comparable] struct {
	events map[EventType][]*EventHandler
}

func NewEventEmitter[EventType comparable]() *EventEmitter[EventType] {
	return &EventEmitter[EventType]{
		events: make(map[EventType][]*EventHandler),
	}
}

func (e *EventEmitter[EventType]) Emit(eventType EventType, params ...any) {
	listeners, found := e.events[eventType]

	if !found {
		return
	}
	for _, listener := range listeners {
		(*listener)(params...)
	}
}

func (e *EventEmitter[EventType]) On(event EventType, listener EventHandler) {
	e.events[event] = append(e.events[event], &listener)
}
