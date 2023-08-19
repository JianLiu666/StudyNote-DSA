package p00230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n)
// Space Complexity: O(n)
func kthSmallest_inorder(root *TreeNode, k int) int {
	monotonicQueue := []int{}
	dfs(root, k, &monotonicQueue)
	return monotonicQueue[k-1]
}

func dfs(root *TreeNode, k int, q *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, k, q)

	*q = append(*q, root.Val)
	if len(*q) >= k {
		return
	}

	dfs(root.Right, k, q)
}
