package circulararrayloop

/**
A circular array has a cycle if:
	- The same set of indices is repeated when the sequence is traversed in accordance with the aforementioned rules.
 	- The length of the sequence is at least two.
	- The loop must be in a single direction, forward or backward.
*/

func circularArrayLoop(nums []int) bool {
	size := len(nums)
	for i := 0; i < size; i++ {
		direction := nums[i] > 0
		slow, fast := i, i
		if isCycle(nums, slow, fast, size, direction) {
			return true
		}
	}

	return false
}

func nextPosition(pointer, value, size int) int {
	nextPosition := (pointer + value) % size
	if nextPosition < 0 {
		nextPosition += size
	}
	return nextPosition
}

func isNotCycle(nums []int, previousDirection bool, pointer int) bool {
	currentDirection := nums[pointer] >= 0

	// If the current value would loop back to itself, or the previous and current directions don't match, an cycle is impossible.
	return (findAbsVal(nums[pointer])%len(nums) == 0) || (previousDirection != currentDirection)
}

func isCycle(nums []int, slow, fast, size int, direction bool) bool {
	for {
		slow = nextPosition(slow, nums[slow], size)
		if isNotCycle(nums, direction, slow) {
			break
		}

		fast = nextPosition(fast, nums[fast], size)
		if isNotCycle(nums, direction, fast) {
			break
		}

		fast = nextPosition(fast, nums[fast], size)
		if isNotCycle(nums, direction, fast) {
			break
		}

		if slow == fast {
			return true
		}
	}

	return false
}

func findAbsVal(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
