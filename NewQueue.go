package expiringQueue

import "time"

// NewQueue - Create and configure the queue.
func NewQueue[T any](expiration time.Duration, equalsFunc func(a, b T) bool) *Queue[T] {
	q := &Queue[T]{
		expiration: expiration.Nanoseconds(),
		equalsFunc: equalsFunc,
	}
	return q
}
