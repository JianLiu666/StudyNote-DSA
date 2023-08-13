package p00076

// Time Complexity: O(n+k), where n is the length of s, k is the length of t
// Space Complexity: O(n+k)
func minWindow(s string, t string) string {
	start, end := 0, len(s)

	sMap := map[byte]int{}
	tMap := map[byte]int{}

	// 紀錄 t 字串的字母出現次數
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	// sliding window 查找 s
	right := 0
	for left := 0; left < len(s); left++ {
		// 持續向右擴大 window range
		for right < len(s) {
			sMap[s[right]]++
			// 一旦 window 包含 t 的所有字母, 就比較 window 大小是否更小
			if include(sMap, tMap) {
				if right-left < end-start {
					start = left
					end = right
				}
				// 如果這裡不扣掉的話, 就會無限遞增上去
				sMap[s[right]]--
				break
			}
			right++
		}
		// 持續向右縮小 window range
		sMap[s[left]]--
	}

	if end == len(s) {
		return ""
	}

	return s[start : end+1]
}

func include(sMap, tMap map[byte]int) bool {
	for k, v := range tMap {
		if sMap[k] < v {
			return false
		}
	}
	return true
}
