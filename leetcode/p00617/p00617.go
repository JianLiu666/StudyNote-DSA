package p00617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n), where n is the depth of tree
// Space Complexity: O(1)
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	head := &TreeNode{}
	dfs(head, root1)
	dfs(head, root2)

	return head
}

func dfs(dst, src *TreeNode) {
	if src == nil {
		return
	}

	dst.Val += src.Val
	if dst.Left == nil && src.Left != nil {
		dst.Left = &TreeNode{}
	}
	dfs(dst.Left, src.Left)

	if dst.Right == nil && src.Right != nil {
		dst.Right = &TreeNode{}
	}
	dfs(dst.Right, src.Right)
}
