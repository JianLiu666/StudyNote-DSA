package p00105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	memo := map[int]int{}
	for idx, val := range inorder {
		memo[val] = idx
	}

	root, _ := helper(preorder, inorder, memo, 0, 0, len(preorder)-1)
	return root
}

func helper(preorder, inorder []int, memo map[int]int, rootIdx, startIdx, endIdx int) (*TreeNode, int) {
	if startIdx > endIdx {
		return nil, rootIdx
	}

	root := &TreeNode{Val: preorder[rootIdx]}
	rootIdx++
	root.Left, rootIdx = helper(preorder, inorder, memo, rootIdx, startIdx, memo[root.Val]-1)
	root.Right, rootIdx = helper(preorder, inorder, memo, rootIdx, memo[root.Val]+1, endIdx)

	return root, rootIdx
}
