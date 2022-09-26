package p00121

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i]-minPrice)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
