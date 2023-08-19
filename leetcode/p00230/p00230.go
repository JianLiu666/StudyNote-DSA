package p00230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(n)
func kthSmallest(root *TreeNode, k int) int {
	index := 0
	result := 0
	dfs(root, k, &index, &result)
	return result
}

func dfs(root *TreeNode, k int, index, result *int) {
	if root == nil {
		return
	}

	dfs(root.Left, k, index, result)

	*index++
	if *index == k {
		*result = root.Val
		return
	}

	dfs(root.Right, k, index, result)
}
