[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reverse-string/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 344. Reverse String

Write a function that reverses a string. The input string is given as an array of characters `s`.

You must do this by modifying the input array `in-place` with `O(1)` extra memory.

### Example 1:

```
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

### Example 2:

```
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
```

### Constraints:

- 1 <= s.length <= $10^5$
- `s[i]` is a `printable ascii character`.

### Related Topics

- Two Pointers
- String
- Recursion
  
---

# 解題方向

Two Pointers (head and tail pointer)

頭尾兩兩交換