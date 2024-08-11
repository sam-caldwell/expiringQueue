package expiringQueue

import (
	"time"
)

// expireNodes - run in the background to expire nodes by age.
func (q *Queue[T]) expireNodes() {
	if q.expiration == 0 {
		return
	}
	for {
		var prev *Node[T]
		time.Sleep(time.Duration(q.expiration) * time.Nanosecond)
		q.lock.Lock()
		if q.pauseExpirationTask {
			q.lock.Unlock()
			return
		}
		now := time.Now().UnixNano()
		for node := q.head; node != nil; node = node.next {
			age := node.timestamp - now
			if (age - q.expiration) < 0 {
				if q.deleteNode(node, prev) {
					q.lock.Unlock()
					return
				}
			}
			prev = node
		}
		q.lock.Unlock()
	}
}
