[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-palindromic-substring/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 5. Longest Palindromic Substring

Given a string `s`, return the longest palindromic substring in `s`.

### Example 1:

```
Input: s = "babad"
Output: "bab"

Explanation:
 - "aba" is also a valid answer.
```

### Example 2:

```
Input: s = "cbbd"
Output: "bb"
```

### Constraints:

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

### Related Topics

- String
- Dynamic Programming
  
---

# 解題方向

### Solved using Brute Force concept

先從最暴力也最直接的方式做起，從 `s` 的開頭開始一步步取出 `substring`，判斷是否構成回文，把最大的回文長度記下來當作後續的剪枝條件

詳細實作就直接看程式碼吧，`TC: O(n^3)`

接著來看看能夠再進一步優化的地方，上面的做法是一次一次取出 `substring` 的開頭開始往後逐步判斷回文直到 `s` 的盡頭，我們可以從這點下手

現在把 `s` 遍歷時的 `index` 當作中心點，由 `index` 開始向左與向右擴散，直到無法構成回文才停下來，如此一來就可以把時間複雜度下降到 `O(n^2)`

詳細十座一樣看程式碼吧

### Solved using Dynamic Programming concept

施工中 ...
