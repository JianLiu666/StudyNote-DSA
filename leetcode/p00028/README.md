[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/implement-strstr/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 28. Implement strStr()

Implement `strStr()`.

Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

#### Clarification:

What should we return when `needle` is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when `needle` is an empty string. This is consistent to C's `strstr()` and Java's `indexOf()`.

### Example 1:

```
Input: haystack = "hello", needle = "ll"
Output: 2
```

### Example 2:

```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

### Constraints:

- 1 <= haystack.length, needle.length <= $10^4$
- `haystack` and `needle` consist of only lowercase English characters.

### Related Topics

- Two Pointers
- String
- String Matching
  
---

# 解題方向

基本的 `string` 操作練習