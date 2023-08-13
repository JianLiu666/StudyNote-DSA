[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/minimum-window-substring/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 76. Minimum Window Substring

Given two strings `s` and `t` of lengths m and n respectively, return the minimum window substring of `s` such that every character in `t` (including duplicates) is included in the window. If there is no such substring, return the empty string `""`.

The testcases will be generated such that the answer is unique.

### Example 1:

```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation:
 - The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
```

### Example 2:

```
Input: s = "a", t = "a"
Output: "a"
Explanation:
 - The entire string s is the minimum window.
```

### Example 3:

```
Input: s = "a", t = "aa"
Output: ""
Explanation:
 - Both 'a's from t must be included in the window.
 - Since the largest window of s only has one 'a', return empty string.
```

### Constraints:

- `m == s.length`
- `n == t.length`
- 1 <= m, n <= $10^5$
- `s` and `t` consist of uppercase and lowercase English letters.

Follow up: Could you find an algorithm that runs in `O(m + n)` time?

### Related Topics

- Hash Table
- String
- Sliding Window
  
---

# 解題方向

運用 `sliding window` + `two pointers` 的概念解題，透過 `right pointer` 一路探詢 `s` 字串直到滿足題目需求後，在不斷收斂 `left pointer` 且檢查是否仍滿足題目需求