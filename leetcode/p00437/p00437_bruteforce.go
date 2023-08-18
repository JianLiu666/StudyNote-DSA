package p00437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n^2)
// Space Complexity: O(logn)
func pathSum_bruteforce(root *TreeNode, targetSum int) int {
	result := 0
	dfs_bruteforce(root, targetSum, &result)
	return result
}

func dfs_bruteforce(root *TreeNode, sum int, result *int) {
	if root == nil {
		return
	}

	try_bruteforce(root, sum, result)
	dfs_bruteforce(root.Left, sum, result)
	dfs_bruteforce(root.Right, sum, result)
}

func try_bruteforce(root *TreeNode, sum int, result *int) {
	if root == nil {
		return
	}
	if root.Val == sum {
		*result++
	}

	sum -= root.Val
	try_bruteforce(root.Left, sum, result)
	try_bruteforce(root.Right, sum, result)
}
