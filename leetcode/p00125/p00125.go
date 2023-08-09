package p00125

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(1)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	head, tail := 0, len(s)-1

	for head < tail {
		for head < tail && !isValid(s[head]) {
			head++
		}
		for head < tail && !isValid(s[tail]) {
			tail--
		}

		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}

func isValid(a byte) bool {
	if a < '0' || (a > '9' && a < 'A') || (a > 'Z' && a < 'a') || a > 'z' {
		return false
	}
	return true
}
