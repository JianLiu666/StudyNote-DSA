package p00342

func isPowerOfFour_bitmanipulation(n int) bool {
	return (n > 0) && (n&(n-1) == 0) && (n&0x55555555 == n)
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
func isPowerOfFour_recursion(n int) bool {
	if n == 0 {
		return false
	}
	for n != 1 {
		if n&3 != 0 {
			return false
		}
		n >>= 2
	}
	return true
}
