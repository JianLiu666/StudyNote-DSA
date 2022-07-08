package p00387

// Time Complexity: O(n)
// Space Complexity: O(n)
func firstUniqChar(s string) int {
	memo := map[rune]int{}

	for _, ch := range s {
		memo[ch]++
	}

	for idx, ch := range s {
		if memo[ch] == 1 {
			return idx
		}
	}

	return -1
}
