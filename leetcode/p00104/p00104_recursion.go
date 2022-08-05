package p00104

// Time Complexity: O(n)
// Space Complexity: O(n)
func maxDepth_topDown(root *TreeNode) int {
	maximum := 0
	topDown(root, 1, &maximum)

	return maximum
}

func topDown(node *TreeNode, depth int, maximum *int) {
	if node == nil {
		return
	}

	if depth > *maximum {
		*maximum = depth
	}

	topDown(node.Left, depth+1, maximum)
	topDown(node.Right, depth+1, maximum)
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func maxDepth_bottomUp(root *TreeNode) int {
	return bottomUp(root, 0)
}

func bottomUp(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	return max(
		bottomUp(node.Left, depth+1),
		bottomUp(node.Right, depth+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
