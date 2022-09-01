package p00005

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func longestPalindrome_bruteforce_optimize(s string) string {
	if len(s) == 1 {
		return s
	}

	maxLen := 1
	start := 0
	// 遍歷整個 s, i 表示回文字串的中心點
	for i := 0; i < len(s)-1; i++ {
		// 回文字串的可能有兩種, 例如: aba, abba
		// 每一次作為中心點時都要考慮回文字串屬於奇數類型還是偶數類型
		tmpLen := max(expandAroundCenter(s, i, i), expandAroundCenter(s, i, i+1))
		if tmpLen > maxLen {
			maxLen = tmpLen
			// 更新回文字串的起始位置
			// maxLen-1 的原因是當回文字串是偶數類型時, 我們要取的中心位置是中間點的左邊
			start = i - (maxLen-1)/2
		}
	}

	return s[start : start+maxLen]
}

func expandAroundCenter(s string, head, tail int) int {
	for head >= 0 && tail < len(s) {
		if s[head] != s[tail] {
			return tail - head - 1
		}
		head--
		tail++
	}
	return tail - head - 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
