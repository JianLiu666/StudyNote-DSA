package p00014

import "strings"

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	size := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < size {
			size = len(strs[i])
		}
	}

	var res strings.Builder

	for ch := 0; ch < size; ch++ {
		for i := 1; i < len(strs); i++ {
			if strs[i][ch] != strs[0][ch] {
				return res.String()
			}
		}
		res.WriteByte(strs[0][ch])
	}

	return res.String()
}
