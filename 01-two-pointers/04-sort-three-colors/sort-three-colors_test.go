package sortthreecolors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortThreeColors(t *testing.T) {
	tt := map[string]struct {
		input    []int
		expected []int
	}{
		"with 1 element": {
			input:    []int{0},
			expected: []int{0},
		},
		"with 3 elements": {
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		"with 4 elements": {
			input:    []int{2, 1, 1, 0},
			expected: []int{0, 1, 1, 2},
		},
		"with many elements": {
			input:    []int{2, 1, 1, 0, 1, 0, 2},
			expected: []int{0, 0, 1, 1, 1, 2, 2},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.EqualValues(t, test.expected, sortThreeColors(test.input))
		})
	}
}
