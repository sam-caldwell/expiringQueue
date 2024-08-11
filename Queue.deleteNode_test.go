package expiringQueue

import (
	"testing"
	"time"
)

func TestDeleteNode(t *testing.T) {
	q := NewQueue[int](time.Second, func(a, b int) bool {
		return a == b
	})

	// Add nodes to the queue
	n1 := &Node[int]{value: 10}
	n2 := &Node[int]{value: 20}
	n3 := &Node[int]{value: 30}
	q.head = n1
	n1.next = n2
	n2.next = n3
	q.tail = n3
	q.count = 3

	// Test deleting head node
	if pause := q.deleteNode(n1, nil); pause {
		t.Errorf("Expected pauseExpirationTask to be false, got true")
	}
	if q.head != n2 {
		t.Errorf("Expected head to be n2, got %v", q.head)
	}
	if q.count != 2 {
		t.Errorf("Expected count to be 2, got %d", q.count)
	}

	// Test deleting tail node
	if pause := q.deleteNode(n3, n2); pause {
		t.Errorf("Expected pauseExpirationTask to be false, got true")
	}
	if q.tail != n2 {
		t.Errorf("Expected tail to be n2, got %v", q.tail)
	}
	if q.count != 1 {
		t.Errorf("Expected count to be 1, got %d", q.count)
	}

	// Test deleting remaining node
	if pause := q.deleteNode(n2, nil); !pause {
		t.Errorf("Expected pauseExpirationTask to be false, got true")
	}
	if q.head != nil || q.tail != nil {
		t.Errorf("Expected head and tail to be nil, got head %v, tail %v", q.head, q.tail)
	}
	if q.count != 0 {
		t.Errorf("Expected count to be 0, got %d", q.count)
	}
}
