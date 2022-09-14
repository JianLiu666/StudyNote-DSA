package p01457

// Time Complexity: O(m), where m is the number of nodes
// Space Complexity: O(n), where n is the depth of tree
func pseudoPalindromicPaths(root *TreeNode) int {
	res := 0
	dfs_bit(root, 0, &res)

	return res
}

func dfs_bit(root *TreeNode, palindromic int, res *int) {
	if root == nil {
		return
	}

	palindromic ^= 1 << root.Val
	if root.Left == root.Right && palindromic&(palindromic-1) == 0 {
		*res++
	} else {
		dfs_bit(root.Left, palindromic, res)
		dfs_bit(root.Right, palindromic, res)
	}
}
