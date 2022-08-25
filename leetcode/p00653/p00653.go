package p00653

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func findTarget(root *TreeNode, k int) bool {
	memo := map[int]bool{}
	return helper(root, k, memo)
}

func helper(node *TreeNode, k int, memo map[int]bool) bool {
	if node == nil {
		return false
	}
	if memo[k-node.Val] {
		return true
	}

	memo[node.Val] = true

	return helper(node.Left, k, memo) || helper(node.Right, k, memo)
}
