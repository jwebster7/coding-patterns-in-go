package oneoffpalindrome

func validPalindrome(input string) bool {

	left, right := 0, len(input)-1

	for left <= right {
		if input[left] != input[right] {
			// Try skipping a character on the right first. If it's not a palindrome,
			// we can try skipping a character on the left.
			// If both substrings are not palindromes, then we know the input cannot
			// possibly be a palindrome by removing one character.
			if !isPalindrome(input, left, right-1) && !isPalindrome(input, left+1, right) {
				return false
			}
		}
		left++
		right--
	}

	return true
}

func isPalindrome(input string, start, end int) bool {
	for start < end {
		if input[start] != input[end] {
			return false
		}
		start++
		end--
	}
	return true
}
