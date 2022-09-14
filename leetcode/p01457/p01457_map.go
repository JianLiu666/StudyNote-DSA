package p01457

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(mlogn), where m is the number of nodes, n is the depth of tree
// Space Complexity: O(m)
func pseudoPalindromicPaths_map(root *TreeNode) int {
	res := 0
	dfs_map(map[int]int{}, root, &res)

	return res
}

func dfs_map(memo map[int]int, root *TreeNode, res *int) {
	if root == nil {
		return
	}

	memo[root.Val]++
	if root.Left == root.Right {
		*res += pseudoPalindromic(memo)
	} else {
		dfs_map(memo, root.Left, res)
		dfs_map(memo, root.Right, res)
	}
	memo[root.Val]--
}

func pseudoPalindromic(memo map[int]int) int {
	odd := false
	for _, count := range memo {
		if count&1 == 1 {
			if !odd {
				odd = true
			} else {
				return 0
			}
		}
	}
	return 1
}
