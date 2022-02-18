package queue

import (
	"math/rand"

	"github.com/Pauloo27/tuner/internal/core/event"
	"github.com/Pauloo27/tuner/internal/core/track"
)

type Queue struct {
	*event.EventEmitter

	queue []track.Track
}

func NewQueue() *Queue {
	return &Queue{
		EventEmitter: event.NewEventEmitter(),
		queue:        []track.Track{},
	}
}

func (q *Queue) Append(item track.Track) {
	q.queue = append(q.queue, item)
	q.Emit(EventAppend, EventAppendData{
		Queue:  q,
		Index:  q.Size() - 1,
		IsMany: false,
		Items:  []track.Track{item},
	})
}

func (q *Queue) AppendMany(items ...track.Track) {
	q.queue = append(q.queue, items...)
	q.Emit(EventAppend, EventAppendData{
		Queue:  q,
		Index:  q.Size() - 1,
		IsMany: true,
		Items:  items,
	})
}

func (q *Queue) AppendAt(index int, item track.Track) {
	var tmp []track.Track
	tmp = append(tmp, q.queue[:index]...)
	tmp = append(tmp, item)
	tmp = append(tmp, q.queue[index:]...)
	q.queue = tmp
	q.Emit(EventAppend, EventAppendData{
		Queue:  q,
		Index:  index,
		IsMany: false,
		Items:  []track.Track{item},
	})
}

func (q *Queue) Remove(index int) {
	if q.Size() == 0 {
		return
	}
	var tmp []track.Track
	tmp = append(tmp, q.queue[:index]...)
	tmp = append(tmp, q.queue[index+1:]...)
	q.queue = tmp
	q.Emit(EventRemove, EventRemoveData{Queue: q, Index: index})
}

func (q *Queue) First() track.Track {
	if q.Size() == 0 {
		return nil
	}
	return q.queue[0]
}

func (q *Queue) All() []track.Track {
	return q.queue
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) Clear() {
	var tmp []track.Track
	q.queue = tmp
}

func (q *Queue) ItemAt(index int) track.Track {
	if index >= q.Size() {
		return nil
	}
	return q.queue[index]
}

func (q *Queue) Shuffle() {
	for i := range q.queue {
		j := rand.Intn(i + 1) //#nosec G404
		// the top element on the list is the one being played,
		// there's no need to move it
		if i == 0 || j == 0 {
			continue
		}
		q.queue[i], q.queue[j] = q.queue[j], q.queue[i]
	}
	q.Emit(EventShuffle)
}
