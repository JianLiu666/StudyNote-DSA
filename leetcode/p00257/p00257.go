package p00257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n), where n is the depth of tree
// Space Complexity: O(n)
func binaryTreePaths(root *TreeNode) []string {
	path := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{path}
	}

	paths := []string{}
	if root.Left != nil {
		paths = append(paths, helper(root.Left, path)...)
	}
	if root.Right != nil {
		paths = append(paths, helper(root.Right, path)...)
	}

	return paths
}

func helper(node *TreeNode, path string) []string {
	path += "->" + strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		return []string{path}
	}

	paths := []string{}
	if node.Left != nil {
		paths = append(paths, helper(node.Left, path)...)
	}
	if node.Right != nil {
		paths = append(paths, helper(node.Right, path)...)
	}

	return paths
}
