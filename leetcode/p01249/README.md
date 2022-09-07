[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1249. Minimum Remove to Make Valid Parentheses

Given a string s of `'('` , `')'` and lowercase English characters.

Your task is to remove the minimum number of parentheses ( `'('` or `')'`, in any positions ) so that the resulting parentheses string is valid and return **any** valid string.

Formally, a parentheses string is valid if and only if:

- It is the empty string, contains only lowercase characters, or
- It can be written as `AB` (`A` concatenated with `B`), where `A` and `B` are valid strings, or
- It can be written as `(A)`, where `A` is a valid string.

### Example 1:

```
Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"

Explanation:
 - "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
```

### Example 2:

```
Input: s = "a)b(c)d"
Output: "ab(c)d"
```

### Example 3:

```
Input: s = "))(("
Output: ""

Explanation:
 - An empty string is also valid.
```

### Constraints:

- 1 <= s.length <= $10^5$
- `s[i]` is either `'('` , `')'`, or lowercase English letter.

### Related Topics

- String
- Stack
  
---

# 解題方向

Stack 暖身題，自己第一次寫的時候是直接用 `Stack` 紀錄 substring

看到別人的寫法是用 `Stack` 紀錄 index，最後在將不合法的 `(` 或 `)` 剔除後輸出，兩個比較直接看程式碼吧