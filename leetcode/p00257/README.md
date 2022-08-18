[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/binary-tree-paths/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 257. Binary Tree Paths

Given the `root` of a binary tree, return all root-to-leaf paths in **any order**.

A **leaf** is a node with no children.

### Example 1:

![image](./image/paths-tree.jpeg)

```
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
```

### Example 2:

```
Input: root = [1]
Output: ["1"]
```

### Constraints:

- The number of nodes in the tree is in the range `[1, 100]`.
- `-100 <= Node.val <= 100`

### Related Topics

- String
- Backtracking
- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Depth-First Search concept

透過 Recursion 實現，只要還有至少一個子節點時就繼續往下走，直到遇到 `leaf node` 後才停止，過程中把尋訪過的 `node.Val` 記錄下來即可
