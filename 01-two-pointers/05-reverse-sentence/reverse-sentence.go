package reversesentence

import (
	"regexp"
	"strings"
)

// reverseSentence reverses the order of words in sentence, but not the words themselves.
func reverseSentence(sentence string) string {
	// Trim the white consecutive and trailing white spaces within the string.
	// ex. " hello there   friend "
	trimmedSentence := trimString(sentence)

	// Reverse the trimmed sentence.
	// ex. "dneirf ereht olleh"
	reversedSentence := reverseArrayInRange([]byte(trimmedSentence), 0, len(trimmedSentence)-1)

	// Reverse each word (separated by a " ") in the reversed input.
	// ex. "dneirf ereht olleh" -> "friend there hello"
	start, end := 0, 0
	for start < len(reversedSentence) {

		// Move the right pointer until we find the end of the sentence or a the end of a word (signified by " ").
		for end < len(reversedSentence) && reversedSentence[end] != ' ' {
			end += 1
		}

		// Reverse only the section of the array where the current word is.
		reversedSentence = reverseArrayInRange(reversedSentence, start, end-1)

		// Move the "start" or left pointer to the start of the next word.
		start = end + 1

		// End currently points to " ". Move the end forward by 1 to start on the next word.
		end += 1
	}

	return string(reversedSentence)
}

// trimString removes trailing and consecutive white space from a supplied string.
func trimString(str string) string {
	regex := regexp.MustCompile(`\s+`)
	return regex.ReplaceAllString(strings.TrimSpace(str), " ")
}

// reverseArrayInRange reverses a certain range of an array and returns a reference to the changed array.
func reverseArrayInRange(input []byte, left, right int) []byte {
	for left < right {
		input[left], input[right] = input[right], input[left]
		left += 1
		right -= 1
	}

	return input
}
