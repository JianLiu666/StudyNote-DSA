package sorting

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func mergesort(nums []int) {
	mergesort_inner(nums, 0, len(nums)-1)
}

func mergesort_inner(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	// devide and conquer
	mergesort_inner(nums, start, mid)
	mergesort_inner(nums, mid+1, end)

	// 臨時紀錄用的陣列
	tmp := make([]int, end-start+1)

	left, right, cursor := start, mid+1, 0
	for left <= mid && right <= end {
		if nums[left] <= nums[right] {
			tmp[cursor] = nums[left]
			left++
		} else {
			tmp[cursor] = nums[right]
			right++
		}
		cursor++
	}
	for left <= mid {
		tmp[cursor] = nums[left]
		left++
		cursor++
	}
	for right <= end {
		tmp[cursor] = nums[right]
		right++
		cursor++
	}

	// 覆蓋回原始陣列
	for i, n := range tmp {
		nums[start+i] = n
	}
}
