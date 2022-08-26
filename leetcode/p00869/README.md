[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reordered-power-of-2/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 869. Reordered Power of 2

You are given an integer `n`. We reorder the digits in any order (including the original order) such that the leading digit is not zero.

Return `true` if and only if we can do this so that the resulting number is a power of two.

### Example 1:

```
Input: n = 1
Output: true
```

### Example 2:

```
Input: n = 10
Output: false
```

### Constraints:

- 1 <= n <= $10^9$

### Related Topics

- Math
- Sorting
- Counting
- Enumeration
  
---

# 解題方向

將 `n` 內的每個數字分別計數後轉成 `string`，接著只要對符合題目範圍的 $2^29$ 內的每個數字逐一比對，只要符合表示可以重新排序成 `2` 的次方

e.g.

- `0010000000 = 2`
- `0000100000 = 4`
- `0000000010 = 8`
- `0100001000 = 16`
- `0011000000 = 32`
- ...