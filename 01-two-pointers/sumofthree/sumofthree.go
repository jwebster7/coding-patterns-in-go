package sumofthree

import (
	"slices"
)

const (
	MAX_INPUT_LENGTH = 500
	MIN_INPUT_LENGTH = 3
	VALID_VALUE      = 1000
)

// findSumOfThree accepts an array of numbers and a target sum to search for.
func findSumOfThree(nums []int, target int) bool {
	// Check the constraints.
	if len(nums) < MIN_INPUT_LENGTH || len(nums) > MAX_INPUT_LENGTH {
		return false
	} else if target > VALID_VALUE || target < VALID_VALUE*-1 {
		return false
	}

	// If nums isn't sorted, the solution below won't work.
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i
		runner := i + 1
		right := len(nums) - 1
		for runner < right {
			currentSum := nums[left] + nums[right] + nums[runner]
			if currentSum == target {
				return true
			} else if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			}
			runner += 1
		}
	}
	return false
}

/**
// altFindSumOfThree accepts an array of numbers and a target sum to search for.
func findSumOfThree(nums []int, target int) bool {
	// Check the constraints.
	if len(nums) < MIN_INPUT_LENGTH || len(nums) > MAX_INPUT_LENGTH {
		return false
	} else if target > VALID_VALUE || target < VALID_VALUE*-1 {
		return false
	}

	// If nums isn't sorted, the solution below won't work.
	slices.Sort(nums)

	left := 0
	right := len(nums) - 1
	for left <= right {
		runner := left + 1
		currentSum := nums[left] + nums[right]
		for runner < right {
			if nums[runner]+currentSum == target {
				return true
			}
			runner += 1
		}

		// If currentSum is less than the target, move the left pointer to a higher value.
		// If currentSum is greater than the target, move the right pointer to a lower value.
		if currentSum < target {
			left++
		} else {
			right--
		}
	}
	return false
}
*/
