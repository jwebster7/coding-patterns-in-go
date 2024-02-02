package linkedlistcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// introduceCycle attempts to introduce a cycle to a linked list.
func introduceCycle(t *testing.T, head *LinkedListNode, val int) {
	t.Helper()

	// Make a temporary pointer to track the node at the head.
	cycle := head

	// Move the head pointer to the last node in the list.
	for head.next != nil {
		head = head.next
	}

	for cycle != nil {
		// If the current node cycle is the value we're looking for,
		if cycle.data == val {
			// Introduce a cycle and return.
			head.next = cycle
			return
		}

		cycle = cycle.next
	}
}

func TestDetectCycle(t *testing.T) {
	tt := map[string]struct {
		data       []int
		cyclePoint int
		expected   bool
	}{
		"should detect a cycle exists": {
			data:       []int{23, 28, 10, 5, 67, 39, 70, 28},
			cyclePoint: 39,
			expected:   true,
		},
		"should detect a cycle exists at the beginning": {
			data:       []int{23, 28, 10, 5, 67, 39, 70, 28},
			cyclePoint: 23,
			expected:   true,
		},
		"should detect a cycle exists at the end": {
			data:       []int{23, 27, 10, 5, 67, 39, 70, 28},
			cyclePoint: 28,
			expected:   true,
		},
		"should not detect a cycle when all values are same": {
			data:     []int{4, 4, 4, 4, 4, 4, 4},
			expected: false,
		},
		"should not detect a cycle exists": {
			data:     []int{288, 224, 275, 390, 4, 383, 330, 60, 193},
			expected: false,
		},
		"should not detect a cycle exists with single node list": {
			data:     []int{288},
			expected: false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ll := LinkedList{}
			ll.CreateLinkedList(test.data)
			ll.DisplayLinkedList()

			if test.expected {
				introduceCycle(t, ll.head, test.cyclePoint)
			}

			assert.Equal(t, test.expected, detectCycle(ll.head))
		})
	}
}
