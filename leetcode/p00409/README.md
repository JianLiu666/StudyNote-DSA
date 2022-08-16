[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-palindrome/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 409. Longest Palindrome

Given a string s which consists of lowercase or uppercase letters, return the length of the **longest palindrome** that can be built with those letters.

Letters are **case sensitive**, for example, `"Aa"` is not considered a palindrome here.

### Example 1:

```
Input: s = "abccccdd"
Output: 7

Explanation:
 - One longest palindrome that can be built is "dccaccd", whose length is 7.
```

### Example 2:

```
Input: s = "a"
Output: 1

Explanation:
 - The longest palindrome that can be built is "a", whose length is 1.
```

### Constraints:

- `1 <= s.length <= 2000`
- `s` consists of lowercase **and/or** uppercase English letters only.


### Related Topics

- Hash Table
- String
- Greedy
  
---

# 解題方向

最大的 Palindrome = `n` 個偶數對個字母(對稱) + `1` 個字母(擺在正中間) 

只要用 `Hash Table` 將出現過的所有字母記錄下來後，在對 `Hash Table` 遍歷過一遍即可
