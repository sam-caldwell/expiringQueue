package expiringQueue

import (
	"testing"
	"time"
)

func TestQueue_struct(t *testing.T) {
	q := Queue[int]{
		head:                &(Node[int]{}),
		tail:                &(Node[int]{}),
		count:               0,
		pauseExpirationTask: true,
		expiration:          time.Now().UnixNano(),
		equalsFunc:          func(a, b int) bool { return true },
	}
	q.lock.Lock()
	defer q.lock.Unlock()
}
