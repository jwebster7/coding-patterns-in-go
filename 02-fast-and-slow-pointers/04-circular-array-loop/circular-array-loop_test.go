package circulararrayloop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularArrayLoop(t *testing.T) {
	tt := map[string]struct {
		nums     []int
		expected bool
	}{
		"should detect a loop one": {
			nums:     []int{1, 3, -2, -4, 1},
			expected: true,
		},
		"should not detect a loop one": {
			nums:     []int{5, -1, 1, 1, -7, -9},
			expected: false,
		},
		"should not detect a loop two": {
			nums:     []int{2, 1, -1, -2},
			expected: false,
		},
		"should not detect a loop three": {
			nums:     []int{2, 5, -4, 3, -1, 4},
			expected: false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, circularArrayLoop(test.nums))
		})
	}
}

func TestNextPosition(t *testing.T) {
	tt := map[string]struct {
		nums          []int
		pointer       int
		expectedIndex int
		expectedValue int
	}{
		"should detect next position without wrapping": {
			nums:          []int{3, 1, -2, -4, 1},
			pointer:       0,
			expectedIndex: 3,
			expectedValue: -4,
		},
		"should detect next position by wrapping backwards": {
			nums:          []int{3, 3, -3, -4, 1},
			pointer:       2,
			expectedIndex: 4,
			expectedValue: 1,
		},
		"should detect next position by wrapping forwards": {
			nums:          []int{3, 3, -2, 4, 1},
			pointer:       3,
			expectedIndex: 2,
			expectedValue: -2,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			position := nextPosition(test.pointer, test.nums[test.pointer], len(test.nums))
			assert.Equal(t, test.expectedIndex, position)
			assert.Equal(t, test.expectedValue, test.nums[position])
		})
	}
}

func TestIsNotCycle(t *testing.T) {
	tt := map[string]struct {
		nums              []int
		pointer           int
		previousDirection bool
		expected          bool
	}{
		"should detect next position without wrapping": {
			nums:    []int{3, 1, -2, -4, 1},
			pointer: 0,
		},
		"should detect next position by wrapping backwards": {
			nums:    []int{3, 3, -3, -4, 1},
			pointer: 2,
		},
		"should detect next position by wrapping forwards": {
			nums:    []int{3, 3, -2, 4, 1},
			pointer: 3,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			position := nextPosition(test.pointer, test.nums[test.pointer], len(test.nums))
			assert.Equal(t, test.expected, position)
		})
	}
}

func TestFindAbsVal(t *testing.T) {
	tt := map[string]struct {
		num      int
		expected int
	}{
		"detect abs value of positive": {
			num:      1,
			expected: 1,
		},
		"detect abs value of negative": {
			num:      -1,
			expected: 1,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, findAbsVal(test.expected))
		})
	}
}
