[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 122. Best Time to Buy and Sell Stock II

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the $i^{th}$ day.

On each day, you may decide to buy and/or sell the stock. You can only hold **at most one** share of the stock at any time. However, you can buy it then immediately sell it on the **same day**.

Find and return the **maximum** profit you can achieve.

### Example 1:

```
Input: prices = [7,1,5,3,6,4]
Output: 7

Explanation:
 - Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
 - Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
 - Total profit is 4 + 3 = 7.
```

### Example 2:

```
Input: prices = [1,2,3,4,5]
Output: 4

Explanation:
 - Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
 - Total profit is 4.
```

### Example 3:

```
Input: prices = [7,6,4,3,1]
Output: 0

Explanation:
 - There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
```

### Constraints:

- `1` <= `prices.length` <= 3 * $10^4$
- `0` <= `prices[i]` <= $10^4$

### Related Topics

- Array
- Dynamic Programming
- Greedy
  
---

# 解題方向

跟[上一題](./../p00121/README.md)差別在於可以不斷的買賣股票，但同時間只能持有一次交易的股票數量

由於我們沒辦法預測股市未來發展，因此我們不知道何時買入才是最佳時機，反過來說如果我們能夠從已知點位往回看的話事情就簡單許多，只要將所有波段的差值加總起來就是最好的交易操作 ~~*(如果現實也可以這麼簡單就好了...)*~~