package p00771

// Time Complexity: O(n)
// Space Complexity: O(n)
func numJewelsInStones(jewels string, stones string) int {
	mapping := map[rune]bool{}

	for _, ch := range jewels {
		mapping[ch] = true
	}

	count := 0
	for _, ch := range stones {
		if _, exists := mapping[ch]; exists {
			count++
		}
	}

	return count
}
