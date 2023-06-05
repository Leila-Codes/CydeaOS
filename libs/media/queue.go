package media

import (
	"fmt"
	"github.com/google/uuid"
)

var (
	ErrDuplicateMedia = fmt.Errorf("media already in queue")
	ErrQueueEmpty     = fmt.Errorf("queue is empty")
)

//type UUID string

type Queue struct {
	Mood
	Queue map[uuid.UUID]struct{}
}

func NewMediaQueue(m Mood) *Queue {
	return &Queue{
		Mood:  m,
		Queue: make(map[uuid.UUID]struct{}),
	}
}

func (q *Queue) Enqueue(id uuid.UUID) error {
	if _, exists := q.Queue[id]; exists {
		return ErrDuplicateMedia
	}

	q.Queue[id] = struct{}{}

	return nil
}

func (q *Queue) Dequeue() (uuid.UUID, error) {
	// TODO: optimise by using an object pool
	// 	items are never removed and just "shuffled" keeping track of visited items

	for id := range q.Queue {
		delete(q.Queue, id)
		return id, nil
	}

	return uuid.UUID{}, ErrQueueEmpty
}

func (q *Queue) HasNext() bool {
	return len(q.Queue) > 0
}
