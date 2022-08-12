package p00110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func isBalanced(root *TreeNode) bool {
	return recursive(root) != -1
}

func recursive(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	ld := recursive(node.Left)
	rd := recursive(node.Right)
	if ld == -1 || rd == -1 || math.Abs(ld-rd) > 1 {
		return -1
	}

	return max(ld, rd) + 1
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
