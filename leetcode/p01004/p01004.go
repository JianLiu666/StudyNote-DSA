package p01004

// Time Complexity: O(n)
// Space Complexity: O(1)
func longestOnes(nums []int, k int) int {
	result := 0

	l := 0
	used := 0

	for r, n := range nums {
		if n == 0 {
			used++
		}
		for l <= r && used > k {
			if nums[l] == 0 {
				used--
			}
			l++
		}

		result = max(result, r-l+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
這題是不固定長度的 sliding window

*/
