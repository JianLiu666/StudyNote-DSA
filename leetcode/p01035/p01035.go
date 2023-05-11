package p01035

// Time Complexity: O(mn), where m is the length of nums1, and n is the length of nums2
// Space Complexity: O(mn)
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	row := len(nums1)
	col := len(nums2)

	// init array to store dp results
	dp := make([][]int, row+1)
	for r := 0; r < row+1; r++ {
		dp[r] = make([]int, col+1)
	}

	for r := 1; r < row+1; r++ {
		for c := 1; c < col+1; c++ {
			if nums1[r-1] == nums2[c-1] {
				dp[r][c] = dp[r-1][c-1] + 1
			} else {
				dp[r][c] = max(dp[r-1][c], dp[r][c-1])
			}
		}
	}

	return dp[row][col]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
