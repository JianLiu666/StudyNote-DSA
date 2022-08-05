[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/binary-tree-preorder-traversal/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 144. Binary Tree Preorder Traversal

Given the `root` of a binary tree, return the preorder traversal of its nodes' values.

### Example 1:

![image](./image/inorder_1.jpeg)

```
Input: root = [1,null,2,3]
Output: [1,2,3]
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

基本的 pre-order traversal 問題 (root-left-right)，用 `Stack` 解決
