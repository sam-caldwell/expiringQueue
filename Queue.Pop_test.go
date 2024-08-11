package expiringQueue

import (
	"testing"
	"time"
)

func TestQueue_Pop(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		// Create a new queue with a reasonable expiration time (e.g., 1 second)
		q := Queue[int]{}

		// Test popping from an empty queue
		value, ok := q.Pop()
		if ok || value != 0 {
			t.Errorf("Expected empty queue to return zero value and false, got value %v, ok %v", value, ok)
		}

		// Push some values into the queue
		q.Push(10)
		q.Push(20)
		q.Push(30)

		// Test popping from a non-empty queue
		value, ok = q.Pop()
		if !ok || value != 10 {
			t.Errorf("Expected value 10, got %v", value)
		}

		// Test popping remaining values
		value, ok = q.Pop()
		if !ok || value != 20 {
			t.Errorf("Expected value 20, got %v", value)
		}

		value, ok = q.Pop()
		if !ok || value != 30 {
			t.Errorf("Expected value 30, got %v", value)
		}

		// Test popping from an empty queue again
		value, ok = q.Pop()
		if ok || value != 0 {
			t.Errorf("Expected empty queue to return zero value and false, got value %v, ok %v", value, ok)
		}
	})
	t.Run("pop after expiration", func(t *testing.T) {
		// Create a new queue with a short expiration time (e.g., 1 Nanoseconds)
		q := Queue[int]{
			expiration: 500 * time.Millisecond.Nanoseconds(),
			equalsFunc: func(a, b int) bool {
				return a == b
			},
		}

		// Push some values into the queue
		q.Push(10)
		q.Push(20)
		q.Push(30)

		// Wait for expiration
		time.Sleep(5 * time.Second)

		// Test popping after expiration expect
		//    ok false because there should be no node.
		//    value should be zero (0) because nothing is there.
		if value, ok := q.Pop(); ok || value != 0 {
			t.Errorf("Expected empty queue to return zero value and false after expiration.\n"+
				"got: value %v,\n"+
				" ok: %v", value, ok)
		}
	})
}
