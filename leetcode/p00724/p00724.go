package p00724

// Time Complexity: O(n)
// Space Complexity: O(1)
func pivotIndex(nums []int) int {
	left := 0
	right := 0

	for _, n := range nums {
		right += n
	}

	for i, n := range nums {
		right -= n
		if left == right {
			return i
		}

		left += n
	}

	return -1
}
