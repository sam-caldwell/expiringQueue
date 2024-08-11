package expiringQueue

import "time"

// NewQueue - Create and configure the queue.
func NewQueue[T any](expiration time.Duration) *Queue[T] {
	q := &Queue[T]{}
	return q
}
