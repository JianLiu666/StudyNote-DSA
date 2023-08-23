package p00153

// Time Complexity: O(logn)
// Space Complexity: O(1)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[0]
	}

	for left != right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			// 屬於遞增區間, 最小值會在最左邊, 也有可能是自己本身
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}
