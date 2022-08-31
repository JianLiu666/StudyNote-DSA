[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/partition-labels/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 763. Partition Labels

You are given a string `s`. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be `s`.

Return a list of integers representing the size of these parts.

### Example 1:

```
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]

Explanation:
 - The partition is "ababcbaca", "defegde", "hijhklij".
 - This is a partition so that each letter appears in at most one part.
 - A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
```

### Example 2:

```
Input: s = "eccbbbbdec"
Output: [10]
```

### Constraints:

- `1 <= s.length <= 500`
- `s` consists of lowercase English letters.

### Related Topics

- Hash Table
- Two Pointers
- String
- Greedy
  
---

# 解題方向

### Solved using Greedy concept

第一遍對 `s` 遍歷時用 `Hash Table` 記下每個字母的最後出現位置(`s[i]`)，接著第二遍時利用 `Greedy` 的概念搭配 `Two Pointers`，當兩個指針相遇時，就代表第一個能夠被 partitioned 的位置出現了

剩下的直接看程式碼吧