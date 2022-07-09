[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/valid-parentheses/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 20. Valid Parentheses

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.

### Example 1:

```
Input: s = "()"
Output: true
```

### Example 2:

```
Input: s = "()[]{}"
Output: true
```

### Example 3:

```
Input: s = "(]"
Output: false
```

### Constraints:

- 1 <= s.length <= $10^4$
- `s` consists of parentheses only `'()[]{}'`.

### Related Topics

- String
- Stack

---

# 解題方向

基本 `Stack` 觀念練習 (i.e., 先進後出)