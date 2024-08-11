package expiringQueue

import "testing"

func TestQueue_StartExpiringNodes(t *testing.T) {
	q := Queue[int]{
		pauseExpirationTask: true,
	}
	q.StartExpiringNodes()
	if q.pauseExpirationTask {
		t.Fatal("expect false")
	}
}
