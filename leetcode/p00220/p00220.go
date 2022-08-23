package p00220

import "sort"

// Time Complexity: O(nlogk)
// Space Complexity: O(k)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	window := []int{}
	for i := 0; i < len(nums); i++ {
		if i > k {
			idx := sort.SearchInts(window, nums[i-k-1]) // tc: O(logk)
			window = append(window[:idx], window[idx+1:]...)
		}

		// tc: O(logk)
		idx := sort.Search(len(window), func(j int) bool {
			return window[j] > nums[i]
		})
		window = append(window[:idx], append([]int{nums[i]}, window[idx:]...)...)

		if idx != 0 && nums[i]-window[idx-1] <= t {
			return true
		} else if idx < len(window)-1 && window[idx+1]-nums[i] <= t {
			return true
		}
	}

	return false
}
