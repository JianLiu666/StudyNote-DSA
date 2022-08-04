package p00700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
func searchBST(root *TreeNode, val int) *TreeNode {
	current := root
	for current != nil {
		if current.Val == val {
			return current
		} else if current.Val < val {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	return nil
}
