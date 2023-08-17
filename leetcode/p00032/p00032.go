package p00032

// Time Complexity: O(n)
// Space Complexity: O(n)
func longestValidParentheses(s string) int {
	result := 0

	stack := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				j := -1
				if len(stack) > 0 {
					j = stack[len(stack)-1]
				}
				result = max(result, i-j)
			} else {
				stack = append(stack, i)
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
