package expiringQueue

// Has checks if a given value exists in the queue.
func (q *Queue[T]) Has(value T) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.equalsFunc == nil {
		panic("nil equalsFunc pointer in expiringQueue Has() method")
	}

	for node := q.head; node != nil; node = node.next {
		if q.equalsFunc(node.value, value) {
			return true
		}
	}

	return false
}
