package oneoffpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {

	tt := map[string]bool{
		"aba":                    true,
		"abca":                   true,
		"acba":                   true,
		"madame":                 true,
		"dead":                   true,
		"abcda":                  false,
		"tebbem":                 false,
		"eeccccbebaeeabebccceea": false,
	}

	for input, expected := range tt {
		assert.Equal(t, expected, validPalindrome(input))
	}
}
