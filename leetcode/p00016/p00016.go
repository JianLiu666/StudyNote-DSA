package p00016

import (
	"math"
	"sort"
)

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt

	sort.Ints(nums)

	for i := range nums {
		head, tail := i+1, len(nums)-1
		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			result = closest(result, sum, target)
			if sum < target {
				head++
			} else {
				tail--
			}
		}
	}

	return result
}

func closest(a, b, target int) int {
	if math.Abs(float64(target-a)) < math.Abs(float64(target-b)) {
		return a
	}
	return b
}
