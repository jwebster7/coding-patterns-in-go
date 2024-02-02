package linkedlistcycle

func detectCycle(head *LinkedListNode) bool {

	// Replace this placeholder return statement with your code

	// Initialize a slow and fast pointer at the head.
	slow, fast := head, head

	// Loop until fast reaches the end of the linked list.
	for fast != nil && fast.next != nil {
		// Move the slow pointer forward by one, and the fast pointer forward by two
		slow = slow.next
		fast = fast.next.next

		// Check if slow equals fast. If so, a cycle exists.
		if slow == fast {
			return true
		}
	}

	// If we get here, it's because no cycle exists.
	return false
}
