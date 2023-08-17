package p00572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	if isSame(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}

func isSame(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}
	if l.Val != r.Val {
		return false
	}

	return isSame(l.Left, r.Left) && isSame(l.Right, r.Right)
}
