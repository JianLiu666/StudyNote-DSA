[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/power-of-three/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 326. Power of Three

Given an integer `n`, return `true` if it is a power of three. Otherwise, return `false`.

An integer `n` is a power of three, if there exists an integer `x` such that n == $3^x$.

### Example 1:

```
Input: n = 27
Output: true
```

### Example 2:

```
Input: n = 0
Output: false
```

### Example 3:

```
Input: n = 9
Output: true
```

### Constraints:

- -$2^{31}$ <= n <= $2^{31}$ - 1
 

#### Follow up:

Could you solve it without loops/recursion?

### Related Topics

- Math
- Recursion
  
---

# 解題方向

### Solved using Recursion concept

遞迴判斷解決

### Solved using Integer Limitation concept

知道整數範圍後，直接取餘數處理

- **Int32** : `retrun 0b1000101010001101011001111011011 % n == 0`
- **Int64** : `return (1.2157665459056929e+19) % n == 0`