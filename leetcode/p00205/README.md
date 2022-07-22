[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/isomorphic-strings/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 205. Isomorphic Strings

Given two strings `s` and `t`, determine if they are isomorphic.

Two strings `s` and `t` are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

### Example 1:

```
Input: s = "egg", t = "add"
Output: true
```

### Example 2:

```
Input: s = "foo", t = "bar"
Output: false
```

### Example 3:

```
Input: s = "paper", t = "title"
Output: true
```

### Constraints:

- 1 <= s.length <= 5 * $10^4$
- `t.length == s.length`
- `s` and `t` consist of any valid ascii character.

### Related Topics

- Hash Table
- String
  
---

# 解題方向

先分析一下 Isomorphic 的規則

- 字母之間的對應順序在替換後還是要保持一致 (e.g., `abaa->djdd`)
- 不會有兩個不同的字母對應到同一個字母 (e.g., `a->b`, `c->b`)

弄清楚之後，只要用 `Hash Table` 將出現過的字母記錄下來即可。

特別記錄一下，因為題目給定的每個 `char` 限制在有限範圍內, 換句話說不會隨著給定字串的長短有變化, 因此 `Hash Table` 的空間複雜度為 `O(1)`