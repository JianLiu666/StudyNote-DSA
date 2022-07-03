package q00283

// Time Complexity: O(n)
// Space Complexity: O(1)
func moveZeroes(nums []int) {
	nonzero_ptr := 0

	for idx, num := range nums {
		if num != 0 {
			nums[idx], nums[nonzero_ptr] = nums[nonzero_ptr], nums[idx]
			nonzero_ptr++
		}
	}
}
