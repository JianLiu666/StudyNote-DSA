package p00606

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(n)
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	left := tree2str(root.Left)
	right := tree2str(root.Right)

	// pattern: root(left-subtree)(right-subtree)

	var res strings.Builder
	res.WriteString(strconv.Itoa(root.Val))

	if left == "" && right == "" {
		return res.String()
	}

	res.WriteByte('(')
	res.WriteString(left)
	res.WriteByte(')')

	if right != "" {
		res.WriteByte('(')
		res.WriteString(right)
		res.WriteByte(')')
	}

	return res.String()
}
