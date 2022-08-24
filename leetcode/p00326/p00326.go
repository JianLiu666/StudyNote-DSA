package p00326

// Time Complexity: O(log3(n))
// Space Complexity: O(1)
func isPowerOfThree_recursion(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func isPowerOfThree_limitation(n int) bool {
	return n > 0 && 0b1000101010001101011001111011011%n == 0
}
