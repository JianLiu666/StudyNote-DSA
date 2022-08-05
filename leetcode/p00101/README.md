[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/symmetric-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 101. Symmetric Tree

Given the `root` of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

### Example 1:

![image](./image/symtree1.jpeg)

```
Input: root = [1,2,2,3,4,4,3]
Output: true
```

### Example 2:

![image](./image/symtree2.jpeg)

```
Input: root = [1,2,2,null,3,null,3]
Output: false
```

### Constraints:

- The number of nodes in the tree is in the range `[1, 1000]`.
- `-100 <= Node.val <= 100`

#### Follow up:

Could you solve it both recursively and iteratively?

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Recursion concept

每一次都對 `left`、`right` 做下列比較：

- between `left` and `right`
- between `left.left` and `right.right`
- between `left.right` and `right.left`

### Solved using Iteration concept

原理同上，只是改成用 `DFS & Stack` 或 `BFS & Queue` 實作而已