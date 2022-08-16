[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/valid-palindrome-ii/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 680. Valid Palindrome II

Given a string `s`, return `true` if the `s` can be palindrome after deleting **at most one** character from it.

### Example 1:

```
Input: s = "aba"
Output: true
```

### Example 2:

```
Input: s = "abca"
Output: true

Explanation:
 - You could delete the character 'c'.
```

### Example 3:

```
Input: s = "abc"
Output: false
```

### Constraints:

- 1 <= s.length <= $10^5$
- `s` consists of lowercase English letters.

### Related Topics

- Two Pointers
- String
- Greedy
  
---

# 解題方向

我們假設兩種情境：

1. 給定的 `s` 不存在不成對的字母: 直接用 Two Pointers 的概念比對解決
2. 給定的 `s` 存在一個以上不成對的字母

假設發生第二種情境時，我們無法確定不成對的字母到底是在 `head pointer` 還是 `tail pointer`，因此我們必須同時比對兩組剩餘的字串: `s[head:tail]` 跟 `s[head+1:head+1]`

倘若其中一組剩餘的字串符合 Palindrome 時代表只有一個不成對的字母，反之不成對的字母超過一個