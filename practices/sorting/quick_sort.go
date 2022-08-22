package sorting

// Time Complexity: O(nlogn), where n is the length of nums, and the worst case's TC is n^2
// Space Complexity: O(n), where n is the implicit call stack
func quicksort(nums []int) {
	quicksort_inner(nums, 0, len(nums)-1)
}

func quicksort_inner(nums []int, start, end int) {
	if start >= end {
		return
	}

	// 從 nums 中挑選出一個指標位置作為 pivot
	pivot := (start + end) / 2
	// 將 pivot 先移動到 nums 的尾巴
	nums[pivot], nums[end] = nums[end], nums[pivot]
	pivot = end

	// 示意圖
	//
	//                                         pivot
	//         start                           end
	//          v                               v
	//        +---+---+---+---+---+---+---+---+---+
	// nums = | 9 | 8 | 7 | 6 | 1 | 4 | 3 | 2 | 5 |
	//        +---+---+---+---+---+---+---+---+---+
	//          ^                           ^
	//         left                        right
	left, right := start, end-1

	// 直到 left == right 時表示本階段的快排結束
	for left < right {
		for left < right && nums[left] <= nums[end] {
			left++
		}
		for left < right && nums[right] >= nums[end] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}

	if nums[left] > nums[end] {
		nums[left], nums[end] = nums[end], nums[left]
		pivot = left
	}

	// 以 pivot 為中心點左右切分，即 divide and conquer
	quicksort_inner(nums, start, pivot-1)
	quicksort_inner(nums, pivot+1, end)
}
