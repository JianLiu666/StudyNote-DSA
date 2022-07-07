package p00198

// Time Complexity: O(n)
// Space Complexity: O(1)
func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	previous := nums[0]
	next := max(nums[1], nums[0])

	for i := 2; i < len(nums); i++ {
		next, previous = max(nums[i]+previous, next), next
	}

	return max(next, previous)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
