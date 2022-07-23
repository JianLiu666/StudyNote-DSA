package p00652

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	seen := map[string]int{}
	res := []*TreeNode{}

	_, res = dfs(root, "", seen, res)
	return res
}

func dfs(node *TreeNode, s1 string, seen map[string]int, res []*TreeNode) (string, []*TreeNode) {
	if node.Left != nil {
		s1, res = dfs(node.Left, s1, seen, res)
	} else {
		s1 += "n"
	}

	s2 := ""
	if node.Right != nil {
		s2, res = dfs(node.Right, s2, seen, res)
	} else {
		s2 += "n"
	}

	s := fmt.Sprintf("%s%s0%d", s1, s2, node.Val)
	if count, exists := seen[s]; exists {
		if count == 0 {
			seen[s]++
			res = append(res, node)
		}
	} else {
		seen[s] = 0
	}

	return s, res
}
