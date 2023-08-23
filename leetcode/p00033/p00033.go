package p00033

// Time Complexity: O(logn)
// Space Complexity: O(1)
func search(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	for head <= tail {
		mid := (head + tail) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[tail] {
			// 如果右邊有序, 且 target 也在這個範圍內
			if nums[mid] < target && nums[tail] >= target {
				head = mid + 1
			} else {
				// 不然就丟到另外一邊
				tail = mid - 1
			}
		} else {
			// 左邊有序, 且 target 也在這個範圍
			if nums[head] <= target && nums[mid] > target {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		}
	}
	return -1
}
