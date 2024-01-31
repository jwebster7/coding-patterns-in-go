package palindrome

const MAX_INPUT_LENGTH = 2*10 ^ 5

func isValidPalindrome(input string) bool {
	if len(input) >= MAX_INPUT_LENGTH {
		return false
	}

	// O(n) time complexity
	l, r := 0, len(input)-1
	for l < r {
		// One improvement would be to convert these elements to an ASCII code or unicode number then compare the those values.
		if input[l] != input[r] {
			return false
		}
		l += 1
		r -= 1
	}

	return true
}

/**
func AlternateIsValidPalindrome(input string) bool {
	if len(input) > MAX_INPUT_LENGTH {
		return false
	}

	n := len(input)
	for i := 0; i < n/2; i++ {
		// 'i' tracks the distance from the leftmost index is from the '0' index,
		// so can use that to dynamically know where our rightmost index should be.
		j := n - i - 1
		if input[i] != input[j] {
			return false
		}

	}

	return true
}
*/
