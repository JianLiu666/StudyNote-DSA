package p00530

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func getMinimumDifference(root *TreeNode) int {
	res := []int{}

	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}

	min := res[1] - res[0]
	for i := 2; i < len(res); i++ {
		if res[i]-res[i-1] < min {
			min = res[i] - res[i-1]
		}
	}

	return min
}
