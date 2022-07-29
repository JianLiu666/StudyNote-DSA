[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-and-replace-pattern/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 890. Find and Replace Pattern

Given a list of strings `words` and a string `pattern`, return a list of `words[i]` that match `pattern`. You may return the answer in **any order**.

A word matches the pattern if there exists a permutation of letters `p` so that after replacing every letter `x` in the pattern with `p(x)`, we get the desired word.

Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.

### Example 1:

```
Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]

Explanation: 
 - "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}. 
 - "ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation, since a and b map to the same letter.
```

### Example 2:

```
Input: words = ["a","b","c"], pattern = "a"
Output: ["a","b","c"]
```

### Constraints:

- `1 <= pattern.length <= 20`
- `1 <= words.length <= 50`
- `words[i].length == pattern.length`
- `pattern` and `words[i]` are lowercase English letters.

### Related Topics

- Array
- Hash Table
- String
  
---

# 解題方向

基本的字串比對，只是資料量變成 `array` 而已

可以參考 [205. Isomorphic Strings](./../p00205/README.md)

