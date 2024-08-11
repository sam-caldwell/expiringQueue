package expiringQueue

import (
	"testing"
	"time"
)

func TestNode_struct(t *testing.T) {
	_ = Node[int]{
		value:     0,
		timestamp: time.Now().UnixNano(),
		next:      &(Node[int]{}),
	}
}
