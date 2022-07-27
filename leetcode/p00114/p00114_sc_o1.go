package p00114

// Time Complexity: O(n)
// Space Complexity: O(1)
func flatten_sc_o1(root *TreeNode) {
	if root != nil {
		dfs_o1(root)
	}
}

func dfs_o1(node *TreeNode) {
	if node.Right != nil {
		dfs_o1(node.Right)
	}
	if node.Left != nil {
		dfs_o1(node.Left)
		current := node.Left
		for current.Right != nil {
			current = current.Right
		}
		current.Right = node.Right
		node.Right = node.Left
		node.Left = nil
	}
}
