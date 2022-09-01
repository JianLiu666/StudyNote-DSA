package p00005

// Time Complexity: O(n^3)
// Space Complexity: O(1)
func longestPalindrome_bruteforce(s string) string {
	if len(s) == 1 {
		return s
	}

	maxLen := 1
	start := 0
	// 遍歷整個 s, i 表示回文字串的開頭
	for i := 0; i < len(s)-1; i++ {
		// j 表示回文字串的結尾, 不斷判斷 s[i:j+1] 是否構成回文
		for j := i + 1; j < len(s); j++ {
			// j-i+1 > maxLen 作為剪枝條件, 如果取出來的長度比已知的最大回文長度還低時就不處理了
			if j-i+1 > maxLen && isPalindrome(s, i, j) {
				maxLen = j - i + 1
				start = i
			}
		}
	}

	return s[start : start+maxLen]
}

func isPalindrome(s string, head, tail int) bool {
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}
