package p02390

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(n)
func removeStars_bruteforce(s string) string {
	var str strings.Builder

	cursor, skip := len(s)-1, 0
	for cursor >= 0 {
		if s[cursor] != '*' {
			if skip == 0 {
				str.WriteByte(s[cursor])
			} else {
				skip--
			}
		} else {
			skip++
		}
		cursor--
	}

	reversed := str.String()
	str.Reset()
	for i := len(reversed) - 1; i >= 0; i-- {
		str.WriteByte(reversed[i])
	}

	return str.String()
}

func removeStars_stack(s string) string {
	stack := []rune{}
	size := 0

	for _, ch := range s {
		if ch != '*' {
			stack = append(stack, ch)
			size++
		} else {
			size--
			stack = stack[:size]
		}
	}

	var str strings.Builder
	for _, ch := range stack {
		str.WriteRune(ch)
	}

	return str.String()
}
