package p00687

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(1)
func longestUnivaluePath(root *TreeNode) int {
	maximum := 0
	dfs(root, &maximum)
	return maximum
}

func dfs(root *TreeNode, maximum *int) int {
	if root == nil {
		return 0
	}

	// 如果本身節點跟子節點的 value 相同時, 這條邊的 count 就可以認列並繼續累加
	left := dfs(root.Left, maximum)
	if root.Left != nil && root.Val == root.Left.Val {
		left++
	} else {
		left = 0
	}

	// 如果本身節點跟子節點的 value 相同時, 這條邊的 count 就可以認列並繼續累加
	right := dfs(root.Right, maximum)
	if root.Right != nil && root.Val == root.Right.Val {
		right++
	} else {
		right = 0
	}

	// 很有可能最長路徑不在 root 節點, 所以在處理完子節點後還要記得更新最長路徑
	*maximum = max(*maximum, left+right)

	// 往上回傳左右子節點中的最長路徑
	return max(left, right)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
