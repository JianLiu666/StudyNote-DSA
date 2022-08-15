[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/binary-tree-postorder-traversal/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 145. Binary Tree Postorder Traversal

Given the `root` of a binary tree, return the postorder traversal of its nodes' values.

### Example 1:

![image](./image/inorder_1.jpeg)

```
Input: root = [1,null,2,3]
Output: [3,2,1]
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

- The number of the nodes in the tree is in the range `[0, 100]`.
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

DFS 暖身題，trikcy 一點的解法是對 BST 由右至左遍歷後，在把要輸出的 `array` 頭尾反轉過來之後就是 **postorder** 了
