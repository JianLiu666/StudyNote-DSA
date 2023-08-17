package p00543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := depth(root)
	return diameter
}

func depth(root *TreeNode) (maxDepth int, maxDiameter int) {
	if root == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := depth(root.Left)
	rightDepth, rightDiameter := depth(root.Right)

	maxDepth = max(leftDepth, rightDepth) + 1

	// 每個 diameter 都需要一個 node 當作左右邊的轉換點
	// 因此這個動作就是在比較用子節點當作轉換點比較大, 還是當前這個 node 當作轉換點比較大
	maxDiameter = max(max(leftDiameter, rightDiameter), leftDepth+rightDepth)

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
