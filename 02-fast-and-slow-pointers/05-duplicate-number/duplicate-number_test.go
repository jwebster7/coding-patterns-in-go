package duplicatenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	tt := map[string]struct {
		nums      []int
		duplicate int
	}{
		"should find duplicate": {
			nums:      []int{3, 4, 4, 4, 2},
			duplicate: 4,
		},
		"should find duplicate in array with 2 elements": {
			nums:      []int{1, 1},
			duplicate: 1,
		},
		"should find duplicate in array when duplicate is at end": {
			nums:      []int{1, 3, 4, 2, 2},
			duplicate: 2,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.duplicate, findDuplicate(test.nums))
		})
	}
}
