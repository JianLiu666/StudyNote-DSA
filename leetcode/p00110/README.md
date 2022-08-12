[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/balanced-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 110. Balanced Binary Tree

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

> a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

### Example 1:

![image](./image/balance_1.jpeg)

```
Input: root = [3,9,20,null,null,15,7]
Output: true
```

### Example 2:

![image](./image/balance_2.jpeg)

```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

### Example 3:

```
Input: root = []
Output: true
```

### Constraints:

The number of nodes in the tree is in the range `[0, 5000]`.
-$10^4$ <= Node.val <= $10^4$

### Related Topics

- Tree
- Depth-First Search 
- Binary Tree
  
---

# 解題方向

### Solved using Recursion concept

按照題目提示條件來檢查逐步檢查 `node.left` 與 `node.right` 的深度差距是否超過 `1` 即可
