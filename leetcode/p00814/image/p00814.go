package image

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(n)
func pruneTree(root *TreeNode) *TreeNode {
	if !dfs(root) {
		return nil
	}
	return root

}

func dfs(root *TreeNode) bool {
	if root == nil {
		return false
	}

	left := dfs(root.Left)
	if !left {
		root.Left = nil
	}
	right := dfs(root.Right)
	if !right {
		root.Right = nil
	}

	return left || right || root.Val == 1
}
