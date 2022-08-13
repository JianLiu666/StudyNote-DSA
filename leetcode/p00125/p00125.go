package p00125

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(n)
func isPalindrome(s string) bool {
	var builder strings.Builder
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			builder.WriteRune(ch)
		} else if ch >= 'A' && ch <= 'Z' {
			builder.WriteRune(ch + 32)
		}
	}

	str := builder.String()
	head := 0
	tail := len(str) - 1
	for head < tail {
		if str[head] != str[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}
