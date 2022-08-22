package p00098

func isValidBST_recursion(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {
		if !isValidBST_recursion(root.Left) || root.Left.Val >= root.Val || max(root.Left) >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		if !isValidBST_recursion(root.Right) || root.Right.Val <= root.Val || min(root.Right) <= root.Val {
			return false
		}
	}

	return true
}

func min(node *TreeNode) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}

func max(node *TreeNode) int {
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}
