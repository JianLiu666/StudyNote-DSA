package p01448

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(n)
func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, maximum int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maximum {
		count++
		maximum = node.Val
	}

	count += dfs(node.Left, maximum) + dfs(node.Right, maximum)

	return count
}
