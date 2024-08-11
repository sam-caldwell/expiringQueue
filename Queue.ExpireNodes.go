package expiringQueue

import "time"

// expireNodes - run in the background to expire nodes by age.
func (q *Queue[T]) expireNodes(expiration int64) {
	q.expiration = expiration
	for {
		var prev *Node[T]
		time.Sleep(time.Duration(q.expiration) * time.Nanosecond)
		q.lock.Lock()
		now := time.Now().UnixNano()
		for node := q.head; node != nil; node = node.next {
			if age := node.timestamp - now; age > q.expiration && q.deleteNode(node, prev) {
				q.lock.Unlock()
				return
			}
			prev = node
		}
		q.lock.Unlock()
	}
}
