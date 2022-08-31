package p00290

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(n)
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	mapping := map[byte]string{}
	visited := map[string]bool{}
	for i := 0; i < len(pattern); i++ {
		if word, exists := mapping[pattern[i]]; exists {
			if word != words[i] {
				return false
			}
		} else {
			if visited[words[i]] {
				return false
			}
			mapping[pattern[i]] = words[i]
			visited[words[i]] = true
		}
	}
	return true
}
