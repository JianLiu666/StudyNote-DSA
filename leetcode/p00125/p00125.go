package p00125

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(n)
func isPalindrome(s string) bool {
	// 過濾非法字元
	var builder strings.Builder
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			builder.WriteRune(ch)
		} else if ch >= 'A' && ch <= 'Z' {
			builder.WriteRune(ch + 32)
		}
	}
	s = builder.String()

	// check by two pointers
	head := 0
	tail := len(s) - 1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}
