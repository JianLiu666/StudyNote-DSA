package p00101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func isSymmetric_recursion(root *TreeNode) bool {
	return recursive(root.Left, root.Right)
}

func recursive(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return recursive(left.Left, right.Right) && recursive(left.Right, right.Left)
}
