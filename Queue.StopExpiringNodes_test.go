package expiringQueue

import "testing"

func TestQueue_StopExpiringNodes(t *testing.T) {
	q := Queue[int]{
		pauseExpirationTask: false,
	}
	q.StopExpiringNodes()
	if !q.pauseExpirationTask {
		t.Fatal("expect true")
	}
}
