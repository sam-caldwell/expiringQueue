package expiringQueue

// Node - Time-stamped queue node.
type Node[T any] struct {
	value     T
	timestamp int64
	next      *Node[T]
}
