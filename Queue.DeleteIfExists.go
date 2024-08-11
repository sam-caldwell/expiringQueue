package expiringQueue

// DeleteIfExists - if a given claimedToken exists, return true and delete the item from the GenericQueue.
func (q *Queue[T]) DeleteIfExists(claimedToken T) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.equalsFunc == nil {
		panic("missing/nil equalsFunction in expiringQueue")
	}
	if q.head == nil {
		return false
	}

	var prevRecord *Node[T] = nil
	currentRecord := q.head
	for currentRecord != nil {
		if q.equalsFunc(currentRecord.value, claimedToken) {
			if prevRecord == nil { // Deleting the head node
				q.head = currentRecord.next
				if q.head == nil { // Queue is now empty
					q.tail = nil
				}
			} else {
				prevRecord.next = currentRecord.next
				if currentRecord.next == nil { // Deleting the tail node
					q.tail = prevRecord
				}
			}
			currentRecord.next = nil
			q.count--
			return true
		}
		prevRecord = currentRecord
		currentRecord = currentRecord.next
	}

	return false
}
