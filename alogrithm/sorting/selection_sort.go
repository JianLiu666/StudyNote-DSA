package sorting

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func selectionsort(nums []int) {
	end := len(nums)
	for i := 0; i < end; i++ {
		min := i
		for j := i + 1; j < end; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
