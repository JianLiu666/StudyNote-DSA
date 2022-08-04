package p00719

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high {
		mid := low + (high-low)/2
		if count(nums, mid, k) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func count(nums []int, distance, k int) bool {
	acc, left, right, size := 0, 0, 0, len(nums)
	for left < size {
		for right < size && nums[right]-nums[left] <= distance {
			right++
		}
		acc += right - left - 1
		if acc >= k {
			return true
		}

		left++
	}

	return false
}
