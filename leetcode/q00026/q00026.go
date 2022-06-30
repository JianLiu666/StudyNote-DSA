package q00026

// Time Complexity: O(n)
// Space Complexity: O(1)
func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]
		}
	}

	return count + 1
}
