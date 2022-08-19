package p00654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func constructMaximumBinaryTree(nums []int) *TreeNode {
	root := &TreeNode{}
	helper(root, nums)
	return root
}

func helper(node *TreeNode, nums []int) {
	if len(nums) == 1 {
		node.Val = nums[0]
		return
	}

	max, idx := getMaxInArray(nums)
	node.Val = max
	if idx != 0 {
		node.Left = &TreeNode{}
		helper(node.Left, nums[:idx])
	}
	if idx != len(nums)-1 {
		node.Right = &TreeNode{}
		helper(node.Right, nums[idx+1:])
	}
}

func getMaxInArray(nums []int) (int, int) {
	max, idx := 0, -1
	for i, n := range nums {
		if n > max {
			max = n
			idx = i
		}
	}
	return max, idx
}
