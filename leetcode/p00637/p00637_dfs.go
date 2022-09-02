package p00637

// Time Complexity: O(n)
// Space Complexity: O(n)
func averageOfLevels(root *TreeNode) []float64 {
	// 記錄每層的 value 總和
	total := make([]int, 0)
	// 紀錄每層的節點數量總和
	counter := make([]int, 0)

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if level == len(counter) {
			total = append(total, node.Val)
			counter = append(counter, 1)
		} else {
			total[level] += node.Val
			counter[level]++
		}

		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}

	dfs(root, 0)

	res := make([]float64, len(total))
	for i := 0; i < len(res); i++ {
		res[i] = float64(total[i]) / float64(counter[i])
	}

	return res
}
