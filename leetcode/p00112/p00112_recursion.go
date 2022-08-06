package p00112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func hasPathSum_recursion(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum_recursion(root.Left, targetSum-root.Val) ||
		hasPathSum_recursion(root.Right, targetSum-root.Val)
}
