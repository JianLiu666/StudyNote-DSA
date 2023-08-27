package p01493

// Time Complexity: O(n)
// Space Complexity: O(1)
func longestSubarray(nums []int) int {
	result := 0

	l := 0
	used := 0

	for r, n := range nums {
		if n == 0 {
			used++
		}
		for l <= r && used > 1 {
			if nums[l] == 0 {
				used--
			}
			l++
		}
		result = max(result, r-l)
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
[now0+1, right]
next0
*/
