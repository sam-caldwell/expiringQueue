package expiringQueue

import (
	"testing"
	"time"
)

func TestNewQueue(t *testing.T) {
	equalsInt := func(a, b int) bool {
		return a == b
	}

	// Define test values
	expiration := 2 * time.Second
	equalFunc := equalsInt

	// Create a new queue using NewQueue
	q := NewQueue(expiration, equalFunc)

	// Check expiration is set correctly
	if q.expiration != expiration.Nanoseconds() {
		t.Errorf("Expected expiration %d, got %d", expiration.Nanoseconds(), q.expiration)
	}

	// Check if equalsFunc is set correctly
	if q.equalsFunc == nil {
		t.Errorf("Expected equalsFunc to be set, but it is nil")
	}

	// Test that the equalsFunc works as expected
	if !q.equalsFunc(10, 10) {
		t.Errorf("Expected equalsFunc to return true for equal values")
	}
	if q.equalsFunc(10, 20) {
		t.Errorf("Expected equalsFunc to return false for non-equal values")
	}
}
