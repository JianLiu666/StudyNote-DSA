package p01372

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	result := 0

	// 給定初始值
	memo := map[*TreeNode][2]int{}
	memo[nil] = [2]int{-1, -1}

	dfs(root, memo, &result)

	return result
}

func dfs(root *TreeNode, memo map[*TreeNode][2]int, result *int) {
	if root == nil {
		return
	}

	dfs(root.Left, memo, result)
	dfs(root.Right, memo, result)

	// 把這個節點能找到的 zigzag 記錄下來
	count := [2]int{memo[root.Left][1] + 1, memo[root.Right][0] + 1}
	memo[root] = count

	*result = max(*result, max(memo[root][0], memo[root][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
分析題目:
需要一個 direction 讓我知道自己現在要往右還是往左走
每個 subtree 找到的最大 zigzag 是固定的, 可以被 hash table 記下來不用重複尋訪
用後序遍歷就可以先走到底 從底往上處理回來
對於每個 leaf node 來說, 就是初始值0
每個 node 應該都要記以自己為主如果是 left or right 時的 zigzag 數量是多少
*/
