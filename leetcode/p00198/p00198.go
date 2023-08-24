package p00198

// Time Complexity: O(n)
// Space Complexity: O(n)
func rob(nums []int) int {
	// 為了後面方便跟 dp 的 index 對齊, 所以補一個用不到的值
	nums = append([]int{0}, nums...)

	// 每一輪只有可能有兩種狀態: 0:搶, 1:不搶
	// dp 的長度 len(nums)+1 的原因是, dp[0] 是我們一開始先給定的初始值
	dp := make([][2]int, len(nums))

	for i := 1; i < len(nums); i++ {
		// 如果我這輪要搶, 我只有一種選擇: 上輪不搶+這輪搶
		dp[i][0] = dp[i-1][1] + nums[i]

		// 如果我這輪不搶, 那我就有兩種選擇:
		// 1. 上輪搶, 2. 上輪不搶
		// 因此我需要相比較找到上輪搶或不搶時的最大值
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
第 i-1 輪的房子有兩種選擇: 搶 or 不搶
                          \   /|
                           \ / |
                            x  |
                           / \ |
                          /   \|
                         /     |
第 i 輪的房子有兩種選擇:   搶 or 不搶
                        v     v -> 不搶就是 +0
                    這輪搶的話至少會增加 nums[i] 的金額
*/
