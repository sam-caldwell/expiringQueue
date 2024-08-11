package expiringQueue

func (q *Queue[T]) StartExpiringNodes() {
	q.pauseExpirationTask = false
	go q.expireNodes()
}
