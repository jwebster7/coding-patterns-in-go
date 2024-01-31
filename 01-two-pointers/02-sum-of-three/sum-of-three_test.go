package sumofthree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfThreeFails(t *testing.T) {
	tt := map[string]struct {
		nums     []int
		target   int
		expected bool
	}{
		"should fail with less than expected amount of elements": {
			nums:   []int{1},
			target: 0,
		},
		"should fail with invalid positive target": {
			nums:   []int{-1, 2, 1, 4},
			target: 1001,
		},
		"should fail with invalid negative target": {
			nums:   []int{-1, 2, 1, 4},
			target: -1000,
		},
		"should fail with invalid elements": {
			nums:   []int{-1, 2, 1, 4},
			target: 1,
		},
		"should pass with valid but unsorted elements and target": {
			nums:     []int{1, 2, -1, 4, -2},
			target:   1,
			expected: true,
		},
		"should pass with more valid elements and target": {
			nums:     []int{3, 7, 1, 2, 8, 4, 5},
			target:   20,
			expected: true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, findSumOfThree(test.nums, test.target))
		})
	}
}
