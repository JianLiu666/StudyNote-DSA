package p00735

import "math"

// Time Complexity: O(n)
// Space Complexity: O(n)
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, n := range asteroids {
		if len(stack) == 0 || n > 0 {
			stack = append(stack, n)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 && math.Abs(float64(n)) > math.Abs(float64(stack[len(stack)-1])) {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, n)
			} else if math.Abs(float64(n)) == math.Abs(float64(stack[len(stack)-1])) {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return stack
}
