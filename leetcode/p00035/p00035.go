package p00035

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func searchInsert(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		mid := (head + tail) >> 1
		// 命中目標, 直接回傳index
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			tail = mid
		} else {
			head = mid + 1
		}
	}

	mid := (head + tail) / 2
	if nums[mid] >= target {
		return head
	}
	return head + 1
}
