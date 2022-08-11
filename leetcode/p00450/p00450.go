package p00450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 完全沒有命中目標
	if root == nil {
		return root
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		// 只有右子節點或是左右子節點都沒有的情況
		if root.Left == nil {
			return root.Right
		}
		// 只有左子節點的情況
		if root.Right == nil {
			return root.Left
		}
		// 左右子節點都存在的情況, 選擇用 right subtree 的最小節點替換
		minNode := findMinNode(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}

	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
