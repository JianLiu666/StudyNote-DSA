package p00448

// Time Complexity: O(n)
// Space Complexity: O(1)
func findDisappearedNumbers(nums []int) []int {
	result := []int{}

	size := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		nums[(nums[i]%size)-1] += size
	}

	for i := 0; i < len(nums); i++ {
		if nums[i]/size == 0 {
			result = append(result, i+1)
		}
	}

	return result
}
