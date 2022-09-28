[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/coin-change/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 322. Coin Change

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

You may assume that you have an infinite number of each kind of coin.

### Example 1:

```
Input: coins = [1,2,5], amount = 11
Output: 3

Explanation:
 11 = 5 + 5 + 1
```

### Example 2:

```
Input: coins = [2], amount = 3
Output: -1
```

### Example 3:

```
Input: coins = [1], amount = 0
Output: 0
```

### Constraints:

- `1 <= coins.length <= 12`
- 1 <= `coins[i]` <= $2^{31}$ - 1
- 0 <= `amount` <= $10^4$

### Related Topics

- Array
- Dynamic Programming
- Breadth-First Search
  
---

# 解題方向

### Solved using Dynamic Programming concept

**Bottom-up** 方式切入，既然我們的目標是把指定的 `amount` 換成最少數量的零錢，我們就從零開始一步一步計算 `[0, amount]` 這些區間的每一筆錢能夠換出多少的零錢

過程中將前一次的結果用 `Array` 記錄下來，省下不必要的重複計算開銷

剩下的部分直接看程式碼吧