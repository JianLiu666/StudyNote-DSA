package p00077

// Time Complexity: O(C(n,k)*k)
// Space Complexity: O(n)
func combine(n int, k int) [][]int {
	res := [][]int{}
	dfs([]int{}, 1, n, k, &res)
	return res
}

func dfs(path []int, pre, n, k int, res *[][]int) {
	if len(path) == k {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := pre; i <= n; i++ {
		path = append(path, i)
		dfs(path, i+1, n, k, res)
		path = path[:len(path)-1]
	}
}
