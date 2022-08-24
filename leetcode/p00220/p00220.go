package p00220

import "sort"

// Time Complexity: O(nlogk)
// Space Complexity: O(k)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	window := []int{}
	for i := 0; i < len(nums); i++ {
		if i > k {
			// tc: O(logk)
			idx := sort.SearchInts(window, nums[i-k-1])
			window = append(window[:idx], window[idx+1:]...)
		}

		// note: 如果是空陣列會回傳 0
		// tc: O(logk)
		idx := sort.Search(len(window), func(j int) bool {
			return window[j] > nums[i]
		})
		// 確保新資料以 ascending order 排序插入
		window = append(window[:idx], append([]int{nums[i]}, window[idx:]...)...)

		// 跟插入位置的鄰近兩筆資料做比較
		if idx != 0 && nums[i]-window[idx-1] <= t {
			return true
		}
		if idx < len(window)-1 && window[idx+1]-nums[i] <= t {
			return true
		}
	}

	return false
}
