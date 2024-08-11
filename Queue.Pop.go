package expiringQueue

// Pop - Pop a value from the head of the queue.
func (q *Queue[T]) Pop() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		var zero T
		return zero, false
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.count--
	q.pauseExpirationTask = q.count == 0
	return value, true
}
