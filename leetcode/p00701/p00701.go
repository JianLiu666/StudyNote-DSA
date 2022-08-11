package p00701

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}

	current := root
	for {
		if current.Val > node.Val {
			if current.Left == nil {
				current.Left = node
				break
			} else {
				current = current.Left
			}
		} else {
			if current.Right == nil {
				current.Right = node
				break
			} else {
				current = current.Right
			}
		}
	}
	return root
}
