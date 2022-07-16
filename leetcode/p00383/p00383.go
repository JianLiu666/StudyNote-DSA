package p00383

// Time Complexity: O(n)
// Space Complexity: O(n)
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	memo := make(map[rune]int, len(magazine))

	for _, ch := range magazine {
		memo[ch]++
	}

	for _, ch := range ransomNote {
		if count, exists := memo[ch]; exists && count > 0 {
			memo[ch]--
		} else {
			return false
		}
	}

	return true
}
