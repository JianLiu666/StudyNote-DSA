package p00410

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func splitArray(nums []int, m int) int {
	low, high := 0, 0
	for _, num := range nums {
		if low < num {
			low = num
		}
		high += num
	}

	for low < high {
		mid := low + (high-low)/2
		if cut(nums, m-1, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

// @param nums the given array
// @param cuts how how many cuts should we did (m - 1)
// @param max the maximum of the sum of the subarrays
//
// @return bool
func cut(nums []int, cuts, max int) bool {
	acc := 0
	for _, num := range nums {
		if acc+num <= max {
			acc += num
		} else {
			cuts--
			if cuts < 0 {
				return false
			}
			acc = num
		}
	}

	return true
}
