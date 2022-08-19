package sorting

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func bubblesort(nums []int) {
	end := len(nums)
	for i := 0; i < end; i++ {
		for j := 1; j < end-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
