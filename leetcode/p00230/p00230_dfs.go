package p00230

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(n)
func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	helper(root, k, &res, &count)
	return res
}

func helper(root *TreeNode, k int, res, count *int) {
	if root == nil {
		return
	}

	helper(root.Left, k, res, count)
	*count++

	if *count == k {
		*res = root.Val
		return
	}

	helper(root.Right, k, res, count)
}
