package p00205

// Time Complexity: O(n)
// Space Complexity: O(1)
func isIsomorphic(s string, t string) bool {
	sPattern, tPattern := map[byte]int{}, map[byte]int{}

	for i := range s {
		if sPattern[s[i]] != tPattern[t[i]] {
			// 兩邊上一次遇到的 index 不一樣, 表示順序不對
			return false
		} else {
			// 將各自的字母紀錄相同的 index
			sPattern[s[i]] = i + 1
			tPattern[t[i]] = i + 1
		}
	}
	return true
}
