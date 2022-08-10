[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 297. Serialize and Deserialize Binary Tree

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

**Clarification**: The input/output format is the same as [how LeetCode serializes a binary tree](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation-). You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

### Example 1:

![image](./image/serdeser.jpeg)

```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
```

### Example 2:

```
Input: root = []
Output: []
```

### Constraints:

- The number of nodes in the tree is in the range [0, $10^4$].
- `-1000 <= Node.val <= 1000`

### Related Topics

- String
- Tree
- Depth-First Search
- Breadth-First Search
- Design
- Binary Tree
  
---

# 解題方向

### Solved using Breadth-First Search concept

**serialize:**

按照 BFS 的思路逐層掃描，有節點就輸出節點的 `value`，沒有就輸出 `null`，簡單明瞭

**deserialize:**

一樣按照 BFS 的思路逐層掃描，只要不是 `null` 就做成一顆新的 child node
