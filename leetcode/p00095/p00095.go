package p00095

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n^3)
// Space Complexity: O(n^2)
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	memo := map[int]map[int][]*TreeNode{}
	return helper(1, n, memo)
}

func helper(start, end int, memo map[int]map[int][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	if _, exists := memo[start][end]; exists {
		return memo[start][end]
	}

	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftSubTrees := helper(start, i-1, memo)
		rightSubTrees := helper(i+1, end, memo)
		for _, leftNode := range leftSubTrees {
			for _, rightNode := range rightSubTrees {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  leftNode,
					Right: rightNode,
				})
			}
		}
	}

	if memo[start] == nil {
		memo[start] = map[int][]*TreeNode{}
	}
	memo[start][end] = res

	return res
}
