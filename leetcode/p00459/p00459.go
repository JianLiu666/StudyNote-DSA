package p00459

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func repeatedSubstringPattern(s string) bool {
	size := len(s)

	// 從一半開始逐步縮小 substring 的範圍
	for i := (size / 2) - 1; i >= 0; i-- {
		substring := i + 1

		// 如果無法被剛好切成 n 個 substring 就不需要檢查
		if size%substring != 0 {
			continue
		}

		found := true
		numOfSubstrings := size/substring - 1

		// 逐一比較每個 substring 是否相同
		for j := 0; j < numOfSubstrings; j++ {
			if s[j*substring:(j+1)*substring] != s[(j+1)*substring:(j+2)*substring] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}

	return false
}
