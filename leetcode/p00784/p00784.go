package p00784

// Time Complexity: O(2^n), where n is the length of s
// Space Complexity: O(1)
func letterCasePermutation(s string) []string {
	res := []string{}
	dfs(s, 0, &res)
	return res
}

func dfs(s string, idx int, res *[]string) {
	if idx >= len(s) {
		*res = append(*res, s)
		return
	}

	if s[idx] < '0' || s[idx] > '9' {
		// 大小寫轉換
		if s[idx] >= 'a' && s[idx] <= 'z' {
			dfs(s[:idx]+string(s[idx]-'a'+'A')+s[idx+1:], idx+1, res)
		} else {
			dfs(s[:idx]+string(s[idx]-'A'+'a')+s[idx+1:], idx+1, res)
		}
	}
	dfs(s, idx+1, res)
}
