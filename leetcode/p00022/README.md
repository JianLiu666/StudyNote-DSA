[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/generate-parentheses/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 22. Generate Parentheses

Given `n` pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

### Example 1:

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

### Example 2:

```
Input: n = 1
Output: ["()"]
```

### Constraints:

- `1 <= n <= 8`

### Related Topics

- Array
- Dynamic Programming
- Backtracking
  
---

# 解題方向

### Solved using Backtracking concept

可以 pruning 的地方在於遇到 `)` 加進去會超過 `(` 的數量時，後面的情境就不用再考慮了，直接看程式碼吧

### Solved using Dynamic Programming concept 

施工中 ...