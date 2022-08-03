package p00154

// Time Complexity: average is O(logn), the worst case is O(n)
// Space Complexity: O(1)
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	if nums[high] > nums[0] {
		return nums[0]
	}

	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] == nums[high] {
			for mid < high && nums[high] == nums[mid] {
				high--
			}
		} else {
			high = mid
		}
	}

	return nums[low]
}
