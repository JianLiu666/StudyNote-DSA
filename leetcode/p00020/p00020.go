package p00020

// Time: Complexity: O(n)
// Space: Complexity: O(1)
func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if (ch == ')' || ch == ']' || ch == '}') && len(stack) == 0 {
			return false
		} else if ch == ')' && stack[len(stack)-1] != '(' {
			return false
		} else if ch == ']' && stack[len(stack)-1] != '[' {
			return false
		} else if ch == '}' && stack[len(stack)-1] != '{' {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
