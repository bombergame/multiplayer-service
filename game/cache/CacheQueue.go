package cache

import (
	"github.com/bombergame/multiplayer-service/game/errs"
)

type Queue struct {
	items    []interface{}
	index    int
	fullErr  *errs.FullQueueError
	emptyErr *errs.EmptyQueueError
}

func NewQueue() *Queue {
	return &Queue{
		items:    make([]interface{}, 0),
		index:    -1,
		fullErr:  errs.NewFullQueueError(),
		emptyErr: errs.NewEmptyQueueError(),
	}
}

func (q *Queue) Add(item interface{}) {
	q.items = append(q.items, item)
	q.index++
}

func (q *Queue) Enqueue(item interface{}) *errs.FullQueueError {
	if q.index == len(q.items) {
		return q.fullErr
	}
	q.index++
	q.items[q.index] = item
	return nil
}

func (q *Queue) Dequeue() (interface{}, *errs.EmptyQueueError) {
	if q.index < 0 {
		return nil, q.emptyErr
	}
	item := q.items[q.index]
	q.index--
	return item, nil
}
