package p00020

// Time Complexity: O(n), where n is the length of s
// Space Complexity: O(n)
func isValid(s string) bool {
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	// tc: O(n)
	for _, symbol := range s {
		if left, exist := mapping[symbol]; exist {
			size := len(stack)
			// 命中任何一個右括號, 檢查是否存在對應的左括號
			if size > 0 && stack[size-1] == left {
				stack = stack[:size-1]
			} else {
				return false
			}

		} else {
			// 左括號直接加入 stack 即可
			stack = append(stack, symbol)
		}
	}

	// 最後檢查有沒有存在未配對到的 '(', '[' 或 '{'
	return len(stack) == 0
}
