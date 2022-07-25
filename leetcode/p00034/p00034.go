package p00034

// Time Complexity: O(logn)
// Space Complexity: O(1)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	anchor := binary_search(nums, 0, len(nums)-1, target)
	if anchor == -1 {
		return []int{-1, -1}
	}

	res := []int{}

	// find left edge of target
	edge := anchor
	for edge != 0 && nums[edge-1] == target {
		edge = binary_search(nums, 0, edge, target)
	}
	res = append(res, edge)

	// find right edge of target
	edge = anchor
	for edge != len(nums)-1 && nums[edge+1] == target {
		edge = binary_search(nums, edge+1, len(nums)-1, target)
	}
	res = append(res, edge)

	return res
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
func binary_search(nums []int, head, tail, target int) int {
	for head <= tail {
		mid := head + (tail-head)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}

	return -1
}
