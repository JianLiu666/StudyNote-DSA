package p00977

import "math"

// Time Complexity: O(n)
// Space Complexity: O(1)
func sortedSquares(nums []int) []int {
	head := 0
	tail := len(nums) - 1

	result := make([]int, len(nums))
	current := len(result) - 1
	for current >= 0 {
		if math.Abs(float64(nums[head])) >= math.Abs(float64(nums[tail])) {
			result[current] = int(math.Pow(float64(nums[head]), 2))
			head++
		} else {
			result[current] = int(math.Pow(float64(nums[tail]), 2))
			tail--
		}
		current--
	}

	return result
}
