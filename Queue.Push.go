package expiringQueue

import "time"

// Push - Push a new value to the queue.
func (q *Queue[T]) Push(element T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	newNode := &Node[T]{
		value:     element,
		timestamp: time.Now().UnixNano(),
	}

	if q.tail != nil {
		q.tail.next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
	q.count++
	q.pauseExpirationTask = false
	go q.expireNodes(q.expiration)
}
