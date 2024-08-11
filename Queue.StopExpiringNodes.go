package expiringQueue

func (q *Queue[T]) StopExpiringNodes() {
	q.pauseExpirationTask = true
}
