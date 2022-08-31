[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/word-pattern/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 290. Word Pattern

Given a `pattern` and a string `s`, find if `s` follows the same pattern.

Here **follow** means a full match, such that there is a bijection between a letter in `pattern` and a **non-empty** word in `s`.

### Example 1:

```
Input: pattern = "abba", s = "dog cat cat dog"
Output: true
```

### Example 2:

```
Input: pattern = "abba", s = "dog cat cat fish"
Output: false
```

### Example 3:

```
Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false
```

### Constraints:

- `1 <= pattern.length <= 300`
- `pattern` contains only lower-case English letters.
- `1 <= s.length <= 3000`
- `s` contains only lowercase English letters and spaces `' '`.
- `s` **does not contain** any leading or trailing spaces.
- All the words in `s` are separated by a **single space**.

### Related Topics

- Hash Table
- String
  
---

# 解題方向

`Hash Table` 暖身題，只要注意出現過的單字不能重複匹配到另一個 `pattern`
 - `e.g. pattern = "abbc", s = "dog cat cat dog"`
