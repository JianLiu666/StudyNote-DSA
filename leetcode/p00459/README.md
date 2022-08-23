[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/repeated-substring-pattern/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 459. Repeated Substring Pattern

Given a string `s`, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

### Example 1:

```
Input: s = "abab"
Output: true

Explanation:
 - It is the substring "ab" twice.
```

### Example 2:

```
Input: s = "aba"
Output: false
```

### Example 3:

```
Input: s = "abcabcabcabc"
Output: true

Explanation:
 - It is the substring "abc" four times or the substring "abcabc" twice.
```

### Constraints:

- 1 <= s.length <= $10^4$
- `s` consists of lowercase English letters.

### Related Topics

- String
- String Matching
  
---

# 解題方向

我們的目標是檢查能否將 `s` 切割成由 `n` 個 `substring` 循環組合而成

換句話說只要不斷的拿 `substring` 比較即可，好像沒什麼好說的，直接看程式碼比較快
