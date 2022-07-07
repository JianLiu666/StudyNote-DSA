package p00414

import "math"

// Time Complexity: O(n)
// Space Complexity: O(1)
func thirdMax(nums []int) int {
	first := math.MinInt64
	second := math.MinInt64
	third := math.MinInt64

	for _, num := range nums {
		if num > first {
			first, second, third = num, first, second
		} else if num > second && num != first {
			second, third = num, second
		} else if num > third && num != second && num != first {
			third = num
		}
	}

	if third != math.MinInt64 {
		return third
	}
	return first
}
