package happynumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	tt := map[int]bool{
		1:          true,
		1000000000: true,
		19:         true,
		7:          true,
		8:          false,
		2147483646: false,
	}

	for input, expected := range tt {
		t.Run(fmt.Sprintf("Test num: %v", input), func(t *testing.T) {
			assert.Equal(t, expected, isHappy(input))
		})
	}
}
