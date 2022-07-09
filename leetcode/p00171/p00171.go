package p00171

import "math"

// Time Complexity: O(n)
// Space Complexity: O(1)
func titleToNumber(columnTitle string) int {
	idx := len(columnTitle) - 1

	num := 0
	for _, ch := range columnTitle {
		num += (int(ch) - 64) * int(math.Pow(26, float64(idx)))
		idx--
	}

	return num
}
