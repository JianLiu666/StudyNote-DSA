[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/valid-anagram/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 242. Valid Anagram

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

### Example 1:

```
Input: s = "anagram", t = "nagaram"
Output: true
```

### Example 2:

```
Input: s = "rat", t = "car"
Output: false
```

### Constraints:

- 1 <= s.length, t.length <= 5 * $10^4$
- `s` and `t` consist of lowercase English letters.

#### Follow up:

What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

### Related Topics

- Hash Table
- String
- Sorting
  
---

# 解題方向

### Solved using data structure

Hash map : memorization