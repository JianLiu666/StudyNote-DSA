package p00022

// Time Complexity: O(2^2n)
// Space Complexity: O(2n)
func generateParenthesis(n int) []string {
	res := []string{}
	dfs(&res, "", 0, 0, n)
	return res
}

func dfs(res *[]string, path string, l, r, max int) {
	if l+r == max*2 {
		*res = append(*res, path)
		return
	}

	if l < max {
		dfs(res, path+"(", l+1, r, max)
	}
	if r < l {
		dfs(res, path+")", l, r+1, max)
	}
}
