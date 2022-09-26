package p00122

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProfit(prices []int) int {
	sum := 0
	maxPrice := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		} else {
			sum += maxPrice - prices[i]
			maxPrice = prices[i]
		}
	}

	return sum
}
