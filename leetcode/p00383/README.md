[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/ransom-note/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 383. Ransom Note

Given two strings `ransomNote` and `magazine`, return `true` if `ransomNote` can be constructed by using the letters from `magazine` and `false` otherwise.

Each letter in `magazine` can only be used once in `ransomNote`.

### Example 1:

```
Input: ransomNote = "a", magazine = "b"
Output: false
```

### Example 2:

```
Input: ransomNote = "aa", magazine = "ab"
Output: false
```

### Example 3:

```
Input: ransomNote = "aa", magazine = "aab"
Output: true
```

### Constraints:

- 1 <= ransomNote.length, magazine.length <= $10^5$
- `ransomNote` and `magazine` consist of lowercase English letters.

### Related Topics

- Hash Table
- String
- Sorting
  
---

# 解題方向

`Hash Table` 暖身題