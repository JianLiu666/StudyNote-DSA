package p01140

// Time Complexity: O()
// Space Complexity: O()
func stoneGameII(piles []int) int {
	// 初始化 dp 範圍
	dp := make([][]int, 101)
	for i := 0; i < 101; i++ {
		dp[i] = make([]int, 101)
	}

	suf := make([]int, 101)
	for i := len(piles) - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + piles[i]
	}

	return solve(piles, dp, suf, 0, 1)
}

func solve(piles []int, dp [][]int, suf []int, idx, m int) int {
	if idx == len(piles) {
		return 0
	}
	if dp[idx][m] != 0 {
		return dp[idx][m]
	}

	sum := 0
	for x := 1; x <= 2*m; x++ {
		if idx+x-1 == len(piles) {
			break
		}

		// 這次拿取的那堆石頭數量
		sum += piles[idx+x-1]

		dp[idx][m] = max(dp[idx][m], sum+suf[idx+x]-solve(piles, dp, suf, idx+x, max(x, m)))
	}
	return dp[idx][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
