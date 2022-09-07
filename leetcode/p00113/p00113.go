package p00113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(n)
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	dfs(root, &res, []int{}, targetSum)

	return res
}

func dfs(root *TreeNode, res *[][]int, path []int, targetSum int) {
	if root == nil {
		return
	}

	targetSum -= root.Val
	path = append(path, root.Val)

	if targetSum == 0 && root.Left == nil && root.Right == nil {
		dst := make([]int, len(path))
		copy(dst, path)
		*res = append(*res, dst)
		return
	}

	if root.Left != nil {
		dfs(root.Left, res, path, targetSum)
	}
	if root.Right != nil {
		dfs(root.Right, res, path, targetSum)
	}
}
