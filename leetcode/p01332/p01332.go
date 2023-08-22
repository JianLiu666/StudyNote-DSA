package p01332

import "strings"

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s == reverse(s) {
		return 1
	}
	return 2
}

func reverse(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}

	return result.String()
}
