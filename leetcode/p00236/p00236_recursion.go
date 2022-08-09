package p00236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func lowestCommonAncestor_recursion(root, p, q *TreeNode) *TreeNode {
	ancestor, _ := recursive(root, p, q)
	return ancestor
}

func recursive(node, p, q *TreeNode) (*TreeNode, bool) {
	if node == nil {
		return nil, false
	}

	ancestor, left := recursive(node.Left, p, q)
	if ancestor != nil && left {
		return ancestor, true
	}
	ancestor, right := recursive(node.Right, p, q)
	if ancestor != nil && right {
		return ancestor, true
	}

	if ((node == p || node == q) && (left || right)) || (left && right) {
		return node, true
	}

	return nil, node == p || node == q || left || right
}
