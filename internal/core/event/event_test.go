package event_test

import (
	"testing"

	"github.com/pauloo27/tuner/internal/core/event"
	"github.com/stretchr/testify/require"
)

type EventType string

func TestEventEmitter(t *testing.T) {
	var (
		helloEvent = EventType("HELLO")
		byeEvent   = EventType("BYE")
	)

	t.Run("create new event emitter", func(t *testing.T) {
		emitter := event.NewEventEmitter[EventType]()
		require.NotNil(t, emitter)
	})

	t.Run("add listener", func(t *testing.T) {
		emitter := event.NewEventEmitter[EventType]()
		require.NotNil(t, emitter)

		emitter.On(helloEvent, func(...interface{}) {})
	})

	t.Run("add listener and call it", func(t *testing.T) {
		emitter := event.NewEventEmitter[EventType]()
		require.NotNil(t, emitter)

		ch := make(chan string, 1)
		data := "expected data"

		emitter.On(helloEvent, func(params ...interface{}) {
			require.Len(t, params, 1)
			strData, ok := params[0].(string)
			require.True(t, ok)
			require.Equal(t, data, strData)
			ch <- strData
		})

		emitter.Emit(helloEvent, data)
		require.Equal(t, data, <-ch)
	})

	t.Run("add 4 listeners and call them", func(t *testing.T) {
		emitter := event.NewEventEmitter[EventType]()
		require.NotNil(t, emitter)

		ch := make(chan string, 4)
		data := "expected data"

		listener := func(params ...interface{}) {
			require.Len(t, params, 1)
			strData, ok := params[0].(string)
			require.True(t, ok)
			require.Equal(t, data, strData)
			ch <- strData
		}

		for i := 0; i < 4; i++ {
			emitter.On(helloEvent, listener)
		}

		emitter.Emit(helloEvent, data)
		for i := 0; i < 4; i++ {
			require.Equal(t, data, <-ch)
		}
	})

	t.Run("add listeners to 2 events and call only one", func(t *testing.T) {
		emitter := event.NewEventEmitter[EventType]()
		require.NotNil(t, emitter)

		ch := make(chan string, 1)
		data := "expected data"

		emitter.On(byeEvent, func(params ...interface{}) {
			t.FailNow()
		})

		emitter.On(helloEvent, func(params ...interface{}) {
			require.Len(t, params, 1)
			strData, ok := params[0].(string)
			require.True(t, ok)
			require.Equal(t, data, strData)
			ch <- strData
		})

		emitter.Emit(helloEvent, data)
	})
}
