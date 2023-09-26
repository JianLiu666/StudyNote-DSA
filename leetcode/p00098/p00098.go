package p00098

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n), where n is the number of tree nodes
// Space Complexity: O(logk), where k is the max-depth of tree
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return dfs(root.Left, lower, root.Val) && dfs(root.Right, root.Val, upper)
}

/**
BST 特性, left subtree 的最大值一定會比 root 還小, right subtree 的最小值一定會比最大值還大
所以可以用 dfs 把 lower bound and upper bound 定好之後往下傳遞, 只要不合法就提早結束
*/
