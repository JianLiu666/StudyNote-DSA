package p00567

// Time Complexity: O(n)
// Space Complexity: O(1)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// 將 s1 轉為 hash table 儲存
	memo := [26]int{}
	for _, ch := range s1 {
		memo[ch-'a']++
	}

	// 將 s2 的 substring 轉成 hash table 儲存 (rolling hash)
	tmp := [26]int{}
	for i := 0; i < len(s1); i++ {
		tmp[s2[i]-'a']++
	}

	if helper(memo, tmp) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		// rolling hash
		tmp[s2[i-len(s1)]-'a']--
		tmp[s2[i]-'a']++
		if helper(memo, tmp) {
			return true
		}
	}

	return false
}

func helper(arr1, arr2 [26]int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
