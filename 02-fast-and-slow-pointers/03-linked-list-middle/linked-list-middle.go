package linkedlistmiddle

// getMiddleNode finds the middle node of a linked list.
// If the list has an even count of nodes, it may have 2 middle nodes and chooses rightmost middle node.
// It's time complexity is O(n), with a space complexity of O(1).
func getMiddleNode(head *LinkedListNode) *LinkedListNode {
	// If there is only one node in the linked list, then immediately return it.
	if head.Next() == nil {
		return head
	}

	// Initialize pointers for slow and fast to head.
	slow, fast := head, head

	// Move both pointers until the fast pointer hits the end.
	for fast != nil && fast.Next() != nil {
		slow = slow.Next()
		fast = fast.Next().Next()
	}

	return slow
}
