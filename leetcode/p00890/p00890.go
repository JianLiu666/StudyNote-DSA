package p00890

// Time Complexity: O(nk)
// Space Complexity: O(2k+n)
func findAndReplacePattern(words []string, pattern string) []string {
	if len(pattern) == 1 {
		return words
	}

	result := []string{}
	for _, word := range words {
		if match(pattern, word) {
			result = append(result, word)
		}
	}

	return result
}

func match(pattern, word string) bool {
	mapping := map[byte]byte{}
	visited := map[byte]bool{}

	for i := 0; i < len(pattern); i++ {
		if !visited[pattern[i]] {
			if _, exists := mapping[word[i]]; exists {
				return false
			}
			visited[pattern[i]] = true
			mapping[word[i]] = pattern[i]

		} else {
			if mapping[word[i]] != pattern[i] {
				return false
			}
		}
	}

	return true
}
