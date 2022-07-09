package p00344

// Time Complexity: O(n)
// Space Complexity: O(1)
func reverseString(s []byte) {
	head := 0
	tail := len(s) - 1

	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
}
