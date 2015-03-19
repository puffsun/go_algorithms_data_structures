package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	queue  []interface{}
	length int
	lock   *sync.Mutex
}

func New() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.length = 0
	queue.lock = new(sync.Mutex)
	return queue
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Enqueue(value interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, value)
	q.length += 1
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.length == 0 {
		return nil, errors.New("Empty queue")
	}

	result := q.queue[0]
	q.queue = q.queue[1:]
	q.length -= 1
	return result, nil
}
