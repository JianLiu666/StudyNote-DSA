package p00652

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(2^n)
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo := map[string]bool{}
	duplicates := map[string]*TreeNode{}

	dfs(memo, duplicates, root)

	res := []*TreeNode{}
	for _, node := range duplicates {
		res = append(res, node)
	}
	return res
}

func dfs(memo map[string]bool, duplicates map[string]*TreeNode, root *TreeNode) string {
	if root == nil {
		return "x"
	}

	key := fmt.Sprintf("%v,%v,%v",
		root.Val,
		dfs(memo, duplicates, root.Left),
		dfs(memo, duplicates, root.Right),
	)

	if memo[key] {
		duplicates[key] = root
		return key
	}

	memo[key] = true
	return key
}
