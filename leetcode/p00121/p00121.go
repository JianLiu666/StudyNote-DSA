package p00121

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProfit(prices []int) int {
	min := prices[0] // 紀錄目前為止最低的價格
	max := prices[0] // 紀錄目前為止最高的價格
	profit := 0      // 紀錄目前為止最好的獲利

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			// 找到更小的買進時機
			min = prices[i]
			max = prices[i]

		} else if prices[i] > max {
			// 找到更好的賣出時機
			max = prices[i]
			// 只有當遇到更好的賣出時機時在比較 profit, 可以減少需要比較的次數
			if max-min > profit {
				profit = max - min
			}
		}
	}

	return profit
}
