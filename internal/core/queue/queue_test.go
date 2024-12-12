package queue_test

import (
	"testing"

	"github.com/pauloo27/tuner/internal/core/queue"
	"github.com/pauloo27/tuner/internal/core/track"
	"github.com/stretchr/testify/require"
)

var (
	helloTrack = track.DummyTrack{
		TrackTitle:  "hello",
		TrackArtist: "john doe",
	}
	supTrack = track.DummyTrack{
		TrackTitle:  "sup",
		TrackArtist: "john doe",
	}
	byeTrack = track.DummyTrack{
		TrackTitle:  "bye",
		TrackArtist: "john doe",
	}
)

func TestAppend(t *testing.T) {
	t.Run("append 1 track", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.Append(helloTrack)
		require.Equal(t, 1, q.Size())
	})

	t.Run("append 2 tracks", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.Append(helloTrack)
		q.Append(byeTrack)
		require.Equal(t, 2, q.Size())
	})

	t.Run("append 2 tracks with AppendMany", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendMany(helloTrack, byeTrack)
		require.Equal(t, 2, q.Size())
	})

	t.Run("append 2 with AppendAt", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendAt(0, byeTrack)
		q.AppendAt(0, helloTrack)
		require.Equal(t, 2, q.Size())
		require.Equal(t, helloTrack, q.ItemAt(0))
		require.Equal(t, byeTrack, q.ItemAt(1))
	})
}

func TestRemove(t *testing.T) {
	t.Run("remove when queue is empty", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.Remove(0)
		require.Equal(t, 0, q.Size())
	})

	t.Run("remove item with invalid index", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.Remove(0)
		require.Equal(t, 0, q.Size())
	})

	t.Run("remove 1 item", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendMany(helloTrack, byeTrack)
		require.Equal(t, 2, q.Size())

		q.Remove(0)
		require.Equal(t, 1, q.Size())
		require.Equal(t, byeTrack, q.ItemAt(0))
	})

	t.Run("remove 1 item from the middle of the queue", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendMany(byeTrack, helloTrack, byeTrack)

		require.Equal(t, 3, q.Size())
		require.Equal(t, byeTrack, q.ItemAt(0))
		require.Equal(t, helloTrack, q.ItemAt(1))
		require.Equal(t, byeTrack, q.ItemAt(2))

		q.Remove(1)
		require.Equal(t, 2, q.Size())
		require.Equal(t, byeTrack, q.ItemAt(0))
		require.Equal(t, byeTrack, q.ItemAt(1))
	})
}

func TestSelectors(t *testing.T) {
	t.Run("select first", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendMany(byeTrack, helloTrack, byeTrack)

		require.Equal(t, 3, q.Size())
		require.Equal(t, byeTrack, q.First())
	})

	t.Run("select first when empty", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		require.Nil(t, q.First())
	})

	t.Run("select item at invalid index", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		require.Nil(t, q.ItemAt(0))
	})

	t.Run("select all", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendMany(byeTrack, helloTrack, byeTrack)

		require.Equal(t, 3, q.Size())
		require.Equal(t, byeTrack, q.ItemAt(0))
		require.Equal(t, helloTrack, q.ItemAt(1))
		require.Equal(t, byeTrack, q.ItemAt(2))

		all := q.All()
		require.NotNil(t, all)
		require.Len(t, all, 3)

		require.Equal(t, byeTrack, all[0])
		require.Equal(t, helloTrack, all[1])
		require.Equal(t, byeTrack, all[2])
	})
}

func TestClear(t *testing.T) {
	t.Run("clear empty queue", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())
		q.Clear()
		require.Equal(t, 0, q.Size())
	})

	t.Run("clear queue with tracks", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendMany(byeTrack, helloTrack, byeTrack)
		require.Equal(t, 3, q.Size())

		q.Clear()
		require.Equal(t, 0, q.Size())
	})
}

func TestShuffle(t *testing.T) {
	t.Run("shuffle empty queue", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())
		q.Shuffle()
		require.Equal(t, 0, q.Size())
	})

	// TODO: run more than once, since it's "random" and by bad luck a
	// bad code might pass? idk =/
	t.Run("shuffle queue with tracks", func(t *testing.T) {
		q := queue.NewQueue()
		require.Equal(t, 0, q.Size())

		q.AppendMany(helloTrack, supTrack, byeTrack)
		require.Equal(t, 3, q.Size())

		q.Shuffle()
		require.Equal(t, 3, q.Size())

		require.Equal(t, helloTrack, q.First()) // should not change the first track
	})
}
