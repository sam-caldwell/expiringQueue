package expiringQueue

import "testing"

func TestQueue_Count(t *testing.T) {
	queue := Queue[int]{
		count: 1337,
	}
	if count := queue.Count(); count != 1337 {
		t.Fatal("count mismatch")
	}
}
