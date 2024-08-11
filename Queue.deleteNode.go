package expiringQueue

// deleteNode - delete the head node
func (q *Queue[T]) deleteNode(n, p *Node[T]) bool {
	if p == nil {
		// If p is nil, we're deleting the head node
		q.head = n.next
	} else {
		p.next = n.next
	}

	if n == q.tail {
		// If n is the tail node, update the tail pointer
		q.tail = p
	}

	n.next = nil // Clear the next pointer of the deleted node
	q.count--
	q.pauseExpirationTask = q.count == 0
	return q.pauseExpirationTask
}
