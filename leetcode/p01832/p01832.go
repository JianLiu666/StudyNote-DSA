package p01832

// Time Complexity: O(n)
// Space Complexity: O(1)
func checkIfPangram(sentence string) bool {
	memo := map[rune]bool{}

	for _, ch := range sentence {
		memo[ch] = true
		if len(memo) == 26 {
			return true
		}
	}

	return false
}
