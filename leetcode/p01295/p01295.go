package p01295

import "math"

// Time Complexity: O(n)
// Space Complexity: O(1)
func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		count += int(math.Log10(float64(n))) & 1
	}

	return count
}
