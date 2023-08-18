package p00662

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func widthOfBinaryTree(root *TreeNode) int {
	result := 1

	root.Val = 1
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		result = max(result, q[size-1].Val-q[0].Val+1)
		for i := 0; i < size; i++ {
			node := q[0]
			if node.Left != nil {
				node.Left.Val = node.Val * 2
				q = append(q, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*2 + 1
				q = append(q, node.Right)
			}
			q = q[1:]
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
