package p00150

import "strconv"

// Time Complexity: O(n)
// Space Complexity: O(n)
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			last := len(stack) - 1
			if s == "+" {
				stack[last-1] += stack[last]
			}
			if s == "-" {
				stack[last-1] -= stack[last]
			}
			if s == "*" {
				stack[last-1] *= stack[last]
			}
			if s == "/" {
				stack[last-1] /= stack[last]
			}
			stack = stack[:last]
		} else {
			n, _ := strconv.Atoi(s)
			stack = append(stack, n)
		}
	}

	return stack[0]
}
