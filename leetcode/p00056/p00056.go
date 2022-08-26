package p00056

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	res := [][]int{intervals[0]}

	cursor := 0
	for _, nums := range intervals {
		if !overlap(res[cursor], nums) {
			res = append(res, nums)
			cursor++
		}
	}

	return res
}

func overlap(nums1, nums2 []int) bool {
	// 因為 intervals 已經事先排序過, 所以只需判斷這個情境是否存在
	if nums2[0] <= nums1[1] {
		nums1[1] = max(nums1[1], nums2[1])
		return true
	}

	return false
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
