package p00220

import "sort"

// Time Complexity: O(n)
// Space Complexity: O(k)
func containsNearbyAlmostDuplicate_bucketsort(nums []int, k int, t int) bool {
	buckets := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(buckets, bucketIdx(nums[i-k-1], t))
		}

		idx := bucketIdx(nums[i], t)
		if _, exists := buckets[idx]; exists {
			return true
		}
		if num, exists := buckets[idx-1]; exists && nums[i]-num <= t {
			return true
		}
		if num, exists := buckets[idx+1]; exists && num-nums[i] <= t {
			return true
		}

		buckets[idx] = nums[i]
	}

	return false
}

func bucketIdx(num, t int) int {
	if num == 0 || t == 0 {
		return num
	}

	if num > 0 {
		return num / (t + 1)
	}

	// 確保正負數不會被放進同一個格子裡
	// e.g. nums = [2147483647,-1,2147483647], k = 1, t = 2147483647
	return num/t - 1
}

// Time Complexity: O(nlogk)
// Space Complexity: O(k)
func containsNearbyAlmostDuplicate_binarysearch(nums []int, k int, t int) bool {
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
