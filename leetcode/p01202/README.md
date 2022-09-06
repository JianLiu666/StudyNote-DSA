[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/smallest-string-with-swaps/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1202. Smallest String With Swaps

You are given a string `s`, and an array of pairs of indices in the string `pairs` where `pairs[i] = [a, b]` indicates 2 indices(0-indexed) of the string.

You can swap the characters at any pair of indices in the given `pairs` **any number of times**.

Return the lexicographically smallest string that `s` can be changed to after using the swaps.

### Example 1:

```
Input: s = "dcab", pairs = [[0,3],[1,2]]
Output: "bacd"

Explaination: 
 - Swap s[0] and s[3], s = "bcad"
 - Swap s[1] and s[2], s = "bacd"
```

### Example 2:

```
Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
Output: "abcd"

Explaination: 
 - Swap s[0] and s[3], s = "bcad"
 - Swap s[0] and s[2], s = "acbd"
 - Swap s[1] and s[2], s = "abcd"
```

### Example 3:

```
Input: s = "cba", pairs = [[0,1],[1,2]]
Output: "abc"

Explaination: 
 - Swap s[0] and s[1], s = "bca"
 - Swap s[1] and s[2], s = "bac"
 - Swap s[0] and s[1], s = "abc"
```

### Constraints:

- 1 <= s.length <= $10^5$
- 0 <= pairs.length <= $10^5$
- `0 <= pairs[i][0], pairs[i][1] < s.length`
- `s` only contains lower case English letters.

### Related Topics

- Hash Table
- String
- Depth-First Search
- Breadth-First Search
- Union Find
  
---

# 解題方向

### Solved using Union Find concept

先將屬於相同子集的 `pair` 歸類完畢後，剩下的就是 `pair` 內部排序後根據 `s` 的 pattern 重新輸出一個新的字串
