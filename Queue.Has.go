package expiringQueue

// Has - search the queue for a given element.
func (q *Queue[T]) Has(value T) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	for node := q.head; node != nil; node = node.next {
		if q.equalsFunc(node.value, value) {
			return true
		}
	}

	return false

}
