[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/perfect-squares/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 279. Perfect Squares

Given an integer `n`, return the least number of perfect square numbers that sum to `n`.

A **perfect square** is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, `1`, `4`, `9`, and 16 are perfect squares while `3` and `11` are not.

### Example 1:

```
Input: n = 12
Output: 3

Explanation:
 - 12 = 4 + 4 + 4.
```

### Example 2:

```
Input: n = 13
Output: 2

Explanation:
 -  13 = 4 + 9.
```

### Constraints:

- 1 <= n <= $10^4$

### Related Topics

- Math
- Dynamic Programming
- Breadth-First Search
  
---

# 解題方向

### Solved using Breadth-First Search Concept

Perfect Squares 的組成肯定是要考慮最接近 `n` 的平方數開始，換句話說 BFS 切入時最好是由大到小的方向開始才能有效降低處理次數

Example

```
先定義一下表達形式: 1(11) 為本次的 square=1, remainder=11

N: 12, 由小到大 

                              12
                              |
                    +---------+----+
                    |         |    |
  round.1          1(11)     4(9) 9(4)
                    |         |    |
              +-----+----+    |    |
              |     |    |    |    |
  round.2    1(10) 4(7) 9(2) 4(5) 9(0)

一樣 N:12 時, 由大到小處理

                   12
                   |
              +----+----+
              |    |    |
  round.1    9(4) 4(9) 1(11)
              |
  round.2    9(0)

```

