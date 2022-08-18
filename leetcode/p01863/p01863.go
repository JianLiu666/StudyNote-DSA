package p01863

// Time Complexity: O(2^n)
// Space Complexity: O(n)
func subsetXORSum(nums []int) int {
	return helper(nums, 0, 0)
}

func helper(nums []int, idx, currentXOR int) int {
	if idx >= len(nums) {
		return currentXOR
	}

	return helper(nums, idx+1, currentXOR^nums[idx]) + helper(nums, idx+1, currentXOR)
}
