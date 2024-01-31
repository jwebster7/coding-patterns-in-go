package removenthnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthNode(t *testing.T) {
	tt := map[string]struct {
		data []int
		n    int
	}{
		"should remove 2nd from last node": {
			data: []int{23, 28, 10, 5, 67, 39, 70, 28},
			n:    2,
		},
		"should remove 3rd from last node": {
			data: []int{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
			n:    3,
		},
		"should remove 4th from last node": {
			data: []int{288, 224, 275, 390, 4, 383, 330, 60, 193},
			n:    4,
		},
		"should remove head": {
			data: []int{69, 8, 49, 106, 116, 112},
			n:    6,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ll := LinkedList{}
			ll.CreateLinkedList(test.data)
			ll.DisplayLinkedList()

			node := removeNthLastNode(ll.head, test.n)
			ll.DisplayLinkedList()

			assert.NotNil(t, node)
		})
	}
}
