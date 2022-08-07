package p00116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Time Complexity: O(n)
// Space Complexity: O(n), where n is the implicit call stack
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	helper(root.Left, root.Right)
	return root
}

func helper(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	left.Next = right
	helper(left.Left, left.Right)
	helper(left.Right, right.Left)
	helper(right.Left, right.Right)
}
