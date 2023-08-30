package p00216

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func combinationSum3(k int, n int) [][]int {
	result := [][]int{}

	dfs(k, 1, n, []int{}, &result)

	return result
}

func dfs(k, idx, n int, path []int, result *[][]int) {
	if len(path) == k && n == 0 {
		dst := make([]int, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}
	if (len(path) < k && n == 0) || len(path) > k || n < 0 || idx > 9 {
		return
	}

	for ; idx <= 9; idx++ {
		path = append(path, idx)
		dfs(k, idx+1, n-idx, path, result)
		path = path[:len(path)-1]
	}
}

/**
分析題目:
單純 backtracking 窮舉問題, 一路到 sum > n 之後 pruning
因為是一路遞增上去的, 所以不會有重複的 number 拿過兩次的問題
*/
