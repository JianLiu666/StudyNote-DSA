package p00098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func isValidBST(root *TreeNode) bool {
	monotonicQueue := []int{}
	return dfs(root, &monotonicQueue)
}

func dfs(root *TreeNode, q *[]int) bool {
	if root == nil {
		return true
	}

	if !dfs(root.Left, q) {
		return false
	}

	*q = append(*q, root.Val)
	if len(*q) >= 2 && (*q)[len(*q)-1] <= (*q)[len(*q)-2] {
		return false
	}

	if !dfs(root.Right, q) {
		return false
	}

	return true
}
