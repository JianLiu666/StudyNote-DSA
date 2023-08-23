package p00046

// Time Complexity: O(n!)
// Space Complexity: O(n)
func permute(nums []int) [][]int {
	result := [][]int{}

	visited := map[int]bool{}
	for _, n := range nums {
		visited[n] = false
	}
	dfs(visited, 0, len(nums), []int{}, &result)

	return result
}

func dfs(visited map[int]bool, i, n int, path []int, result *[][]int) {
	if i == n {
		dst := make([]int, n)
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	for k, v := range visited {
		if v {
			continue
		}

		visited[k] = true
		path = append(path, k)

		dfs(visited, i+1, n, path, result)

		visited[k] = false
		path = path[:len(path)-1]
	}
}
