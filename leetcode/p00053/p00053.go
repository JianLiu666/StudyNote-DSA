package p00053

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxSubArray(nums []int) int {
	// 預設選擇第一個 element 做為初始值
	result := nums[0]

	dp := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		// 每次都可以選擇要繼續累加前面的, 或者是歸零重來
		dp[i] = max(dp[i-1], 0) + nums[i-1]
		result = max(result, dp[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
