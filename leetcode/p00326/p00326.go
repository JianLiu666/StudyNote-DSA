package p00326

// Time Complexity: O(log3(n))
// Space Complexity: O(1)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}
