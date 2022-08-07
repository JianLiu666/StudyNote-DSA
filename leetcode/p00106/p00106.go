package p00106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
	memo := make(map[int]int, len(inorder))
	for idx, val := range inorder {
		memo[val] = idx
	}

	root, _ := helper(inorder, postorder, memo, len(postorder)-1, 0, len(postorder)-1)
	return root
}

func helper(inorder, postorder []int, memo map[int]int, rootIdx, startIdx, endIdx int) (*TreeNode, int) {
	if startIdx > endIdx {
		return nil, rootIdx
	}

	root := &TreeNode{Val: postorder[rootIdx]}
	rootIdx--
	root.Right, rootIdx = helper(inorder, postorder, memo, rootIdx, memo[root.Val]+1, endIdx)
	root.Left, rootIdx = helper(inorder, postorder, memo, rootIdx, startIdx, memo[root.Val]-1)

	return root, rootIdx
}
