package p00153

// Time Complexity: O(logn)
// Space Complexity: O(1)
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	// it's a sorted array without any rotated
	if nums[right] > nums[left] {
		return nums[0]
	}

	for left != right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
