from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def maxProfit(self, prices: List[int]) -> int:
        maximum_profit = 0
        min_price = prices[0]

        for price in prices:
            profit = price - min_price
            if profit > maximum_profit:
                maximum_profit = profit
            
            if price < min_price:
                min_price = price

        return maximum_profit

if __name__ == '__main__':
    s = Solution()

    assert s.maxProfit([7,1,5,3,6,4]) == 5
    assert s.maxProfit([7,6,4,3,1]) == 0