package expiringQueue

import "testing"

func TestQueue_Push(t *testing.T) {
	queue := Queue[int]{}
	if queue.count != 0 {
		t.Fatal("count expects 0")
	}
	if queue.head != nil {
		t.Fatal("head expects nil")
	}
	if queue.tail != nil {
		t.Fatal("tail expects nil")
	}
	if queue.pauseExpirationTask {
		t.Fatal("pauseExpirationTask expects false")
	}
	if queue.expiration != 0 {
		t.Fatal("expiration expects 0")
	}
	queue.Push(1)
	if queue.count != 1 {
		t.Fatal("count expects 1")
	}
	if v := queue.head.value; v != 1 {
		t.Fatalf("head expects 1. got: %d", v)
	}
	if v := queue.tail.value; v != 1 {
		t.Fatalf("tail expects 1. got: %d", v)
	}
	if queue.pauseExpirationTask {
		t.Fatal("pauseExpirationTask expects false")
	}
	if queue.expiration != 0 {
		t.Fatal("expiration expects 0")
	}
	if queue.head.timestamp == 0 {
		t.Fatal("timestamp expects not 0")
	}
	queue.Push(2)
	if queue.count != 2 {
		t.Fatal("count expects 2")
	}
	if v := queue.head.value; v != 1 {
		t.Fatalf("head expects 1. got: %d", v)
	}
	if v := queue.tail.value; v != 2 {
		t.Fatalf("tail expects 2. got: %d", v)
	}
	if queue.pauseExpirationTask {
		t.Fatal("pauseExpirationTask expects false")
	}
	if queue.expiration != 0 {
		t.Fatal("expiration expects 0")
	}
	if queue.head.timestamp == 0 {
		t.Fatal("timestamp expects not 0")
	}
}
