[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 103. Binary Tree Zigzag Level Order Traversal

Given the `root` of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

### Example 1:

![image](./image/tree1.jpeg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
```

### Example 2:

```
Input: root = [1]
Output: [[1]]
```

### Example 3:

```
Input: root = []
Output: []
```

### Constraints:

- The number of nodes in the tree is in the range `[0, 2000]`.
- `-100 <= Node.val <= 100`

### Related Topics

- Tree
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

BFS 暖身題，只要注意在保存資料時是要從 `array` 的頭還是尾巴依序塞入資料即可