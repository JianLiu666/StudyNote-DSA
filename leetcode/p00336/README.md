[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/palindrome-pairs/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 336. Palindrome Pairs

You are given a 0-indexed array of unique strings `words`.

A palindrome pair is a pair of integers `(i, j)` such that:

- `0 <= i, j < words.length`,
- `i != j`, and
- `words[i] + words[j]` (the concatenation of the two strings) is a palindrome

Return an array of all the palindrome pairs of `words`.

### Example 1:

```
Input: words = ["abcd","dcba","lls","s","sssll"]
Output: [[0,1],[1,0],[3,2],[2,4]]
Explanation: 
 - The palindromes are ["abcddcba","dcbaabcd","slls","llssssll"]
```

### Example 2:

```
Input: words = ["bat","tab","cat"]
Output: [[0,1],[1,0]]
Explanation: 
 - The palindromes are ["battab","tabbat"]
```

### Example 3:

```
Input: words = ["a",""]
Output: [[0,1],[1,0]]
Explanation: 
 - The palindromes are ["a","a"]
```

### Constraints:

- `1 <= words.length <= 5000`
- `0 <= words[i].length <= 300`
- `words[i]` consists of lowercase English letters.

### Related Topics

- Array
- Hash Table
- String
- Trie
  
---

# 解題方向

### Hash Table 的解法

每個 `word` 都有可能本身就符合 palindrome，也有可能找得到其他的 `word` 一起組成 palindrome，要注意的是構成 palindrome 的可能只是其中某段 `subword`  
所以我們必須對每個 `word` 逐一處理，檢查 `subword` 是否屬於 palindrome, 另一半的 `subword` 是否找得到對應的 `word` 組成 palindrome
