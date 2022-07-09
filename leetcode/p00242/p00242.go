package p00242

// Time Complexity: O(n)
// Space Complexity: O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	memo := map[rune]int{}

	for _, ch := range s {
		memo[ch]++
	}

	for _, ch := range t {
		if count, exists := memo[ch]; exists {
			if count == 0 {
				return false
			}
			memo[ch]--
		} else {
			return false
		}
	}

	return true
}
