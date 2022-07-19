package p00151

import (
	"strings"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func reverseWords(s string) string {
	var res strings.Builder

	tail := 0

	// 0: search word tail, 1: search word head
	state := 0
	for i := len(s) - 1; i > -1; i-- {
		if state == 0 && s[i] != ' ' && (i == len(s)-1 || s[i+1] == ' ') {
			tail = i
			state = 1
		}
		if state == 1 && s[i] != ' ' && (i == 0 || s[i-1] == ' ') {
			if res.Len() != 0 {
				res.WriteRune(' ')
			}
			res.WriteString(s[i : tail+1])
			state = 0
		}
	}

	return res.String()
}
