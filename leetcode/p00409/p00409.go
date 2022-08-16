package p00409

// Time Complexity: O(n)
// Space Complexity: O(n)
func longestPalindrome(s string) int {
	memo := map[rune]int{}
	for _, ch := range s {
		memo[ch]++
	}

	longest := 0
	hasCenter := false
	for _, count := range memo {
		if count%2 == 0 {
			longest += count
		} else {
			longest += count - 1
			if !hasCenter {
				longest++
				hasCenter = true
			}
		}
	}

	return longest
}
