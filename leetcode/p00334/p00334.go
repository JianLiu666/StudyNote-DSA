package p00334

import "math"

// Time Complexity: O(n)
// Space Complexity: O(1)
func increasingTriplet(nums []int) bool {
	n1, n2 := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if n1 >= n {
			n1 = n
		} else if n2 >= n {
			n2 = n
		} else {
			return true
		}
	}
	return false
}
