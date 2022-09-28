package p00461

// Time Complexity: O(n), where n is the length of bits
// Space Complexity: O(1)
func hammingDistance(x int, y int) int {
	xor := x ^ y
	count := 0
	for xor > 0 {
		if xor&1 == 1 {
			count++
		}
		xor >>= 1
	}
	return count
}
