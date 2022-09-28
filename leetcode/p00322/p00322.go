package p00322

// Time Complexity: O(nm), where n is the number of coins, m is the amount
// Space Complexity: O(m)
func coinChange(coins []int, amount int) int {
	// 開出一個跟 amount 一樣大的空間來紀錄過去的計算結果, 空間換時間
	dp := make([]int, amount+1)
	// 初始化為 -1 表示從 [1, amount] 的面額目前還沒有辦法被找開
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for i := 0; i <= amount; i++ {
		// 沒辦法靠 coins 找開的面額就直接跳過
		if dp[i] == -1 {
			continue
		}
		// 將現在的面額(i)加上 coin, 紀錄每個 dp[i]+coin 找開時需要用到的最少零錢數量是多少
		for _, coin := range coins {
			if i+coin > amount {
				continue
			}
			if dp[i+coin] == -1 {
				dp[i+coin] = dp[i] + 1
			} else {
				dp[i+coin] = min(dp[i+coin], dp[i]+1)
			}
		}
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
