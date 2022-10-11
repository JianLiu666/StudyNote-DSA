[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/letter-case-permutation/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 784. Letter Case Permutation

Given a string `s`, you can transform every letter individually to be lowercase or uppercase to create another string.

Return a list of all possible strings we could create. Return the output in **any order**.

### Example 1:

```
Input: s = "a1b2"
Output: ["a1b2","a1B2","A1b2","A1B2"]
```

### Example 2:

```
Input: s = "3z4"
Output: ["3z4","3Z4"]
```

### Constraints:

- `1 <= s.length <= 12`
- `s` consists of lowercase English letters, uppercase English letters, and digits.

### Related Topics

- String
- Backtracking
- Bit Manipulation
  
---

# 解題方向

`DFS` 暖身題，當遇到字母時處理大小寫轉換即可