[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/validate-binary-search-tree/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 98. Validate Binary Search Tree

Given the `root` of a binary tree, determine if it is a valid binary search tree (BST).

A **valid BST** is defined as follows:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

### Example 1:

![image](./image/tree1.jpeg)

```
Input: root = [2,1,3]
Output: true
```

### Example 2:

![image](./image/tree2.jpeg)

```
Input: root = [5,1,4,null,null,3,6]
Output: false

Explanation:
 - The root node's value is 5 but its right child's value is 4.
```

### Constraints:

- The number of nodes in the tree is in the range [1, $10^4$].
- -$2^{31}$ <= Node.val <= $2^{31}$ - 1

### Related Topics

- Tree
- Depth-First Search
- Binary Search Tree
- Binary Tree
  
---

# 解題方向

### Solved using Depth-First Search concept

用 DFS 實作 inorder traversal 得到 `array`，在對這個 `array` 檢查是否符合 **ascending order**

### Solved using Recursion concept

以 inorder traversal 的思路實作 recursion function，將尋訪過程中的 `node.val` 包裝成 `array` 回傳回來後比對是否符合 **ascending order**