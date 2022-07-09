package p00070

// Time Complexity: O(n)
// Space Complexity: O(1)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	next := 2
	previous := 1

	for i := 3; i <= n; i++ {
		next, previous = next+previous, next
	}

	return next
}
