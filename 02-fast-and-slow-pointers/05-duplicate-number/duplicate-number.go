package duplicatenumber

// Notes:
//   - This problem requires understanding of Linked List Cycles and Floyd's Algorithm.
//   - the length of 'num' is 'n+1', but every value within the array is between [1, n].
//   - The key mental leaps with this type of problem are recognizing that you need to use
//     each value in the array as a pointer to another element in the array, and that
//     you can use Floyd's Algorithm.
//   - i.e. [1, 3, 4, 2, 2];
//     -- nums[0] = 1
//     -- nums[nums[0]] = 4
//     -- nums[nums[nums[0]]] = 2
func findDuplicate(nums []int) int {
	return findEntryPoint(nums, findIntersectionPoint(nums))
}

func findIntersectionPoint(nums []int) int {
	// Start fast and slow at the first element.
	slow, fast := nums[0], nums[0]
	for {
		slow, fast = nums[slow], nums[nums[fast]]

		// Slow and fast meet at the intersection point.
		if slow == fast {
			break
		}
	}

	return fast
}

func findEntryPoint(nums []int, intersectionPoint int) int {
	// Loop until slow meets fast again, but this time only move 1 at a time.
	// Reset slow to the beginning but start slowTwo at the intersection point.
	// The distance between the first element and intersection point from the entry point to the cycle should be the same.
	// When fast and slow meet, they are at the first node that the cycle starts (entry point).
	slow := nums[0]
	slowTwo := intersectionPoint
	for slow != slowTwo {
		slow = nums[slow]
		slowTwo = nums[slowTwo]
	}

	return slowTwo
}
