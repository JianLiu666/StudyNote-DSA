package p00124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	leftSum := dfs(root.Left, result)
	rightSum := dfs(root.Right, result)

	maximum := root.Val
	if leftSum >= 0 {
		maximum += leftSum
	}
	if rightSum >= 0 {
		maximum += rightSum
	}
	*result = max(*result, maximum)

	// 左右的加總結果都是負數的話就直接回傳自己本身就好, 否則就是左右挑一個取大往上丟
	if leftSum < 0 && rightSum < 0 {
		return root.Val
	}
	return max(leftSum, rightSum) + root.Val

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
