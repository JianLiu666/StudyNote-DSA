package p00518

// Time Complexity: O(mn), where m is the amount, n is the length of coins
// Space Complexity: O(m)
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= len(coins); i++ {
		for j := coins[i-1]; j <= amount; j++ {
			dp[j] += dp[j-coins[i-1]]
		}
	}

	return dp[amount]
}

/**
公式簡化:
 1. dp[i][j] =          dp[i-1][j-0*coins[i]]          + dp[i-1][j-1*coins[i]]          + ... + dp[i-1][j-k*coins[i]]
 2. dp[i][j-coins[i]] = dp[i-1][j-coins[i]-0*coins[i]] + dp[i-1][j-coins[i]-1*coins[i]] + ... + dp[i-1][j-coins[i]-(k-1)*coins[i]]
                                                                                                 ^ 如果在 i,j 都一樣的情況下, 這裡是 k-1
    dp[i][j-coins[i]] = dp[i-1][j-coins[i]] + dp[i-1][j-2*coins[i]] + ... + dp[i-1][j-k*coins[i]]

 則:
 dp[i][j] - dp[i][j-coins[i]] = dp[i-1][j]
 整理一下:
 dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]

 dp[coins][amount] 可以被簡化成一維的 dp[amount]
*/
