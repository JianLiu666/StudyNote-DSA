[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/climbing-stairs/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 70. Climbing Stairs

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

### Example 1:

```
Input: n = 2
Output: 2

Explanation:
 - There are two ways to climb to the top.
 - 1. 1 step + 1 step
 - 2. 2 steps
```

### Example 2:

```
Input: n = 3
Output: 3

Explanation:
 - There are three ways to climb to the top.
 - 1. 1 step + 1 step + 1 step
 - 2. 1 step + 2 steps
 - 3. 2 steps + 1 step
```

### Constraints:

- `1 <= n <= 45`

### Related Topics

- Math
- Dynamic Programming
- Memorization

---

# 解題方向

DP 經典入門題 : Bottom-up DP

首先試著將原問題分割成更小的子問題，即：

- 如果我們想要踏上第 `n` 階的話，只有兩種可能：
  - 從第 `n-1` 階踏一步上去
  - 從第 `n-2` 階踏兩步上去

- 換句話說，我們可以把題目轉換成這個表達式：`climbStairs(n) = climbStairs(n-1) + climbStairs(n-2)`

接著確認這些問題子問題的答案要如何計算出來，即：

```
       ┌ 1,               if n == 1
f(n) = ┤ 2,               if n == 2
       └ f(n-1) + f(n-2), if n >= 3
```

### Solved using recursion concept

```python
# Time Complexity: O(2^n)
# Space Complexity: O(n)
def climbStairs(n: int) -> int:
    if n == 1:
        return 1
    elif n == 2:
        return 2
    
    return climbStairs(n-1) + climbStairs(n-2)
```

### Solved using dynamic programming concept

```python
# Time Complexity: O(n)
# Space Complexity: O(n)
def climbStairs(n: int) -> int:
    memo = {}
    memo[1] = 1
    memo[2] = 2

    for i in range(3, n+1, 1):
        memeo[i] = memo[i-1] + memo[i-2]
    
    return memo[n]
```

觀察一下可以發現，其實不需要將所有結果都保存下來，只需紀錄第 `n-1` 與 `n-2` 的結果即可

```python
# Time Complexity: O(n)
# Space Complexity: O(1)
def climbStairs(n: int) -> int:
    nxt = 2
    pre = 1

    for i in range(3, n+1, 1):
        nxt, pre = nxt+pre, nxt
    
    return nxt
```