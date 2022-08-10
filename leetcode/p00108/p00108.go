package p00108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func sortedArrayToBST(nums []int) *TreeNode {
	mid := len(nums) / 2

	root := &TreeNode{Val: nums[mid]}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if mid+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}
