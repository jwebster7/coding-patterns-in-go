package reversesentence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseSentence(t *testing.T) {
	tt := map[string]string{
		"hello world":                        "world hello",
		"hello   world ":                     "world hello",
		" hello":                             "hello",
		" this sentence starts with a space": "space a with starts sentence this",
		"1 2 3 4  5 ":                        "5 4 3 2 1",
	}

	for input, expected := range tt {
		t.Run(input, func(t *testing.T) {
			actual := reverseSentence(input)
			assert.EqualValues(t, expected, actual)
		})
	}
}
