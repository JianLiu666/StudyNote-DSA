package p00557

import "strings"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func reverseWords(s string) string {
	var res strings.Builder

	strs := strings.Split(s, " ")
	for i, str := range strs {
		for j := len(str) - 1; j > -1; j-- {
			res.WriteByte(str[j])
		}

		if i != len(strs)-1 {
			res.WriteRune(' ')
		}
	}

	return res.String()
}
