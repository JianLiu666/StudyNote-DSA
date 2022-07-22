package p00205

// Time Complexity: O(n)
// Space Complexity: O(1)
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mappingTable := map[byte]byte{}
	seen := map[byte]bool{}

	for i := 0; i < len(s); i++ {
		if ch, exists := mappingTable[s[i]]; exists {
			if ch != t[i] {
				return false
			}

		} else {
			if _, exists := seen[t[i]]; exists {
				return false
			}

			mappingTable[s[i]] = t[i]
			seen[t[i]] = true
		}
	}

	return true
}
