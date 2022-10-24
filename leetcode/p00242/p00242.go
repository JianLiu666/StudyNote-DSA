package p00242

// Time Complexity: O(n)
// Space Complexity: O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// rune 可以保存 unicode 編碼
	memo := map[rune]int{}
	for _, ch := range s {
		memo[ch]++
	}

	for _, ch := range t {
		memo[ch]--
		if memo[ch] < 0 {
			return false
		}
	}

	return true
}
