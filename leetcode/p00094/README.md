[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/binary-tree-inorder-traversal/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 94. Binary Tree Inorder Traversal

Given the `root` of a binary tree, return the inorder traversal of its nodes' values.

### Example 1:

![image](./image/inorder_1.jpeg)

```
Input: root = [1,null,2,3]
Output: [1,3,2]
```

### Example 2:

```
Input: root = []
Output: []
```

### Example 3:

```
Input: root = [1]
Output: [1]
```

### Constraints:

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

#### Follow up:

Recursive solution is trivial, could you do it iteratively?

### Related Topics

- Stack
- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using DFS concept with explicit stack

直接用 DFS 處理會因為 recursion 的次數過多導致 call stack 發生 **stack overflow** 的問題，因此改用 explicit stack 的做法取代 recursion 的 implicit stack