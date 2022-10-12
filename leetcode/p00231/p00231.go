package p00231

// Time Complexity: O(1)
// Space Complexity: O(1)
func isPowerOfTwo(n int) bool {
	return n != 0 && n&(n-1) == 0
}
