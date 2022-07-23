package p00003

// Time Complexity: O(n)
// Space Complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	seen := map[rune]int{}

	max_distance := 0
	start := 0
	for cursor, ch := range s {
		if index, exists := seen[ch]; exists && index >= start {
			start = index + 1
		} else if cursor-start > max_distance {
			max_distance = cursor - start
		}

		seen[ch] = cursor
	}

	return max_distance + 1
}
