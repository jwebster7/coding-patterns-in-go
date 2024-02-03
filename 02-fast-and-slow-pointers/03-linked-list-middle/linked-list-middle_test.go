package linkedlistmiddle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMiddleNode(t *testing.T) {
	tt := map[string]struct {
		data     []int
		expected int
	}{
		"should detect middle with odd number of elements": {
			data:     []int{1, 2, 3, 4, 5},
			expected: 3,
		},
		"should detect middle with even number of elements": {
			data:     []int{23, 28, 10, 5, 67, 39, 70, 28},
			expected: 67,
		},
		"should detect middle with two elements": {
			data:     []int{4, 5},
			expected: 5,
		},
		"should detect middle with single element": {
			data:     []int{23},
			expected: 23,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ll := LinkedList{}
			ll.CreateLinkedList(test.data)
			ll.DisplayLinkedList()

			if ll.GetLength(ll.head) == 0 {
				t.FailNow()
			}

			assert.Equal(t, test.expected, getMiddleNode(ll.head).data)
		})
	}
}
