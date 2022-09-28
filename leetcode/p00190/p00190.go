package p00190

// Time Complexity: O(1)
// Space Complexity: O(1)
func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	for i := 31; i >= 0; i-- {
		if num&1 == 1 {
			ans += 1 << i
		}
		num >>= 1
	}
	return ans
}
