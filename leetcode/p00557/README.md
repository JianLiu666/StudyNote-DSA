[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reverse-words-in-a-string-iii/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 557. Reverse Words in a String III

Given a string `s`, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

### Example 1:

```
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
```

### Example 2:

```
Input: s = "God Ding"
Output: "doG gniD"
```

### Constraints:

- 1 <= s.length <= 5 * $10^4$
- `s` contains printable **ASCII** characters.
- `s` does not contain any leading or trailing spaces.
- There is **at least one** word in `s`.
- All the words in `s` are separated by a single space.

### Related Topics

- Two Pointers
- String
  
---

# 解題方向

`string` 暖身題，看是要用語言提供的 built-in function 把字串切開，或是自行處理都可以

接著再按照每個單字逐一反轉即可
