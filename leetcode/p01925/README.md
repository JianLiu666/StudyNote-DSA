[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/count-square-sum-triples/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 1925. Count Square Sum Triples

A **square triple** `(a,b,c)` is a triple where `a`, `b`, and `c` are integers and $a^2$ + $b^2$ = $c^2$.

Given an integer `n`, return the number of **square triples** such that `1 <= a, b, c <= n`.

### Example 1:

```
Input: n = 5
Output: 2

Explanation:
 - The square triples are (3,4,5) and (4,3,5).
```

### Example 2:

```
Input: n = 10
Output: 4

Explanation:
 - The square triples are (3,4,5), (4,3,5), (6,8,10), and (8,6,10).
```

### Constraints:

- `1 <= n <= 250`

### Related Topics

- Math
- Enumeration
  
---

# 解題方向

先用 `Hash Table` 把 `[1, n]` 之間的平方數記錄下來，接著再透過 Enumeration 把所有 $a^2$ + $b^2$ 的可能枚舉出來
