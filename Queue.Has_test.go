package expiringQueue

import (
	"testing"
	"time"
)

func TestQueue_Has(t *testing.T) {
	equalsInt := func(a, b int) bool {
		return a == b
	}
	q := NewQueue[int](10*time.Second, equalsInt)
	q.equalsFunc = equalsInt

	// Test with an empty queue
	if q.Has(10) {
		t.Errorf("Expected value 10 to be absent in the empty queue")
	}

	// Push some values into the queue
	q.Push(10)
	q.Push(20)
	q.Push(30)

	// Test existing values
	if !q.Has(10) {
		t.Errorf("Expected value 10 to be found in the queue")
	}
	if !q.Has(20) {
		t.Errorf("Expected value 20 to be found in the queue")
	}
	if !q.Has(30) {
		t.Errorf("Expected value 30 to be found in the queue")
	}

	// Test non-existing value
	if q.Has(40) {
		t.Errorf("Expected value 40 to be absent in the queue")
	}
}
