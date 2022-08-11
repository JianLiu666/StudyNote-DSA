package p00235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n), where n is the height of bst
// Space Complexity: O(1)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
}
