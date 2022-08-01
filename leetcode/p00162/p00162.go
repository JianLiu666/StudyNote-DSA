package p00162

// Time Complexity: O(logn)
// Space Complexity: O(1)
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right-left)/2
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			} else {
				left = mid + 1
			}
		} else if mid == len(nums)-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			} else {
				right = mid - 1
			}
		} else if nums[mid] < nums[mid-1] {
			right = mid - 1
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
}
