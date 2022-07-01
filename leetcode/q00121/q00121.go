package q00121

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProfit(prices []int) int {
	maximum_profit := 0
	min_price := prices[0]

	for _, price := range prices {
		profit := price - min_price
		if profit > maximum_profit {
			maximum_profit = profit
		}

		if price < min_price {
			min_price = price
		}
	}

	return maximum_profit
}
