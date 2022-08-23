[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/invert-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 226. Invert Binary Tree

Given the `root` of a binary tree, invert the tree, and return its root.

### Example 1:

![image](./image/invert1-tree.jpeg)

```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```

### Example 2:

![image](./image/invert2-tree.jpeg)

```
Input: root = [2,1,3]
Output: [2,3,1]
```

### Example 3:

```
Input: root = []
Output: []
```

### Constraints:

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

只要一開始把 `root.left` 跟 `root.right` 交換，之後只要每一個 `node` 一路個別整理下去就好