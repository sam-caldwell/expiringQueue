package expiringQueue

import (
	"sync"
)

// Queue - Expiring Queue.
type Queue[T any] struct {
	head                *Node[T]
	tail                *Node[T]
	count               int
	pauseExpirationTask bool
	expiration          int64 //in nanoseconds
	equalsFunc          func(a, b T) bool
	lock                sync.Mutex
}
