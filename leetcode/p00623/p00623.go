package p00623

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n), where n is equal to depth
// Space Complexity: O(1)
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}

	dfs(root, val, depth)
	return root
}

func dfs(root *TreeNode, val, depth int) {
	if root == nil {
		return
	}

	if depth == 2 {
		root.Left = &TreeNode{
			Val:   val,
			Left:  root.Left,
			Right: nil,
		}
		root.Right = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: root.Right,
		}

	} else {
		dfs(root.Left, val, depth-1)
		dfs(root.Right, val, depth-1)
	}
}
