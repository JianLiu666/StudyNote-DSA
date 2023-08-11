package p00424

// Time Complexity: O(n)
// Space Complexity: O(1)
func characterReplacement(s string, k int) int {
	result := 0

	// 用來記錄 sliding window 內的字母出現次數
	count := [26]int{}

	j := 0
	for i := 0; i < len(s); i++ {
		// 不斷驗證當前的範圍是否能滿足 replacement 的條件
		// 如果可以就繼續擴大 sliding window 的右邊範圍直到無法滿足
		for j < len(s) && canChange(&count, s[j], k, j-i+1) {
			j++
		}
		result = max(result, j-i)

		// 扣除 sliding window 最左邊的字母
		count[s[i]-'A']--
	}

	return result
}

func canChange(count *[26]int, char byte, k, length int) bool {
	count[char-'A']++

	// 找出 sliding window 中出現頻率最高的字母出現次數
	maxCount := 0
	for i := 0; i < len(count); i++ {
		maxCount = max(maxCount, count[i])
	}

	// 確認 replacement 的替換次數是否夠用
	if length-maxCount <= k {
		return true
	}

	count[char-'A']--
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
