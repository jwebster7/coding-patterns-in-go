package removenthnode

func removeNthLastNode(head *LinkedListNode, n int) *LinkedListNode {
	// Initialize left and right pointers to head.
	left := head
	right := head

	// Move the right pointer "n" places to the right.
	for i := 0; i < n; i++ {
		right = right.next
	}

	// If the right pointer is nil it means the head is the one that needs to be removed.
	if right == nil || right.next == nil {
		return head.next
	}

	// Slide left and right until right hits nil / the end of the list.
	for right != nil && right.next != nil {
		left = left.next
		right = right.next
	}

	// At this point, left should be right before the node we need to return.
	// Copy a reference to it and return the node.
	left.next = left.next.next
	return head
}
