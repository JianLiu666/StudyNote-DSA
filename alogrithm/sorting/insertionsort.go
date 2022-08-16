package sorting

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func insertionsort(nums []int) {
	end := len(nums)
	for i := end - 2; i >= 0; i-- {
		for j := i; j < end-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break
			}
		}
	}
}
