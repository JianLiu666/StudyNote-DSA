[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/word-search-ii/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 212. Word Search II

Given an `m x n` `board` of characters and a list of strings `words`, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where **adjacent cells** are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

### Example 1:

![image](./image/search1.jpeg)

```
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
```

### Example 2:

![image](./image/search2.jpeg)

```
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []
```

### Constraints:

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `board[i][j]` is a lowercase English letter.
- 1 <= words.length <= 3 * $10^4$
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.
- All the strings of `words` are unique.

### Related Topics

- Array
- String
- Backtracking
- Trie
- Matrix
  
---

# 解題方向

### Solved using Trie concept

最近很喜歡往 LeetCodeCN 看大神的分享，[這篇](https://leetcode.cn/problems/word-search-ii/solution/gong-shui-san-xie-yi-ti-shuang-jie-hui-s-am8f/)解釋的很詳細 ~~(肯定不是因為中文比較親切的關係)~~
