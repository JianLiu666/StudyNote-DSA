[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reverse-integer/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 7. Reverse Integer

Given a signed 32-bit integer `x`, return `x` with its digits reversed. If reversing `x` causes the value to go outside the signed 32-bit integer range [-$2^{31}$, $2^{31}$ - 1], then return `0`.

**Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**

### Example 1:

```
Input: x = 123
Output: 321
```

### Example 2:

```
Input: x = -123
Output: -321
```

### Example 3:

```
Input: x = 120
Output: 21
```

### Constraints:

- -$2^{31}$ <= x <= $2^{31}$ - 1

### Related Topics

- Math
  
---

# 解題方向

### Solved using string concept

先將 `n` 轉為 `string` 後，反轉 `string` 再轉回 `integer`，結束。

### Sovled using recursion concept

不斷對 `n` 取 modulo 10 的餘數組成新的數字，直到 `n==0` 後結束。