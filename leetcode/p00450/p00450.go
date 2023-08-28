package p00450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(logn)
// Space Complexity: O(logn)
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 前序遍歷
	if root.Val == key {
		// case.1 要刪除的節點沒有任何子節點
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// case.2 要刪除的節點只有一邊有子節點
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}

		// case.3 兩邊都有子節點, 挑最小的子節點來交換
		// 這邊挑的是右子樹最小的 (或是要挑左子樹最大的也可以)
		// 取到最小的子節點後, 一定要先把他從原本的位置刪除後, 才可以提上來重新接好左右子樹
		minNode := findMinNode(root.Right)
		leftSubtree := root.Left
		rightSubtree := deleteNode(root.Right, minNode.Val)
		minNode.Left = leftSubtree
		minNode.Right = rightSubtree
		return minNode
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func findMinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return findMinNode(root.Left)
}
