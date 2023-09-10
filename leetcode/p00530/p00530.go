package p00530

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(logn)
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	dfs(root, nil, &result)
	return result
}

func dfs(root, prev *TreeNode, result *int) *TreeNode {
	if root == nil {
		return prev
	}

	// in-order traversal
	// left
	prev = dfs(root.Left, prev, result)

	// root
	if prev != nil {
		*result = min(*result, abs(root.Val-prev.Val))
	}
	prev = root

	// right
	prev = dfs(root.Right, prev, result)

	return prev
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
