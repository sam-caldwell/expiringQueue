package expiringQueue

// Count - return the number of elements in the queue
func (q *Queue[T]) Count() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.count
}
