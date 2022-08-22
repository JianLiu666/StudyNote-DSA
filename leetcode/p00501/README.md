[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-mode-in-binary-search-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 501. Find Mode in Binary Search Tree

Given the `root` of a binary search tree (BST) with duplicates, return all the [mode(s)](https://en.wikipedia.org/wiki/Mode_(statistics)) (i.e., the most frequently occurred element) in it.

If the tree has more than one mode, return them in **any order**.

Assume a BST is defined as follows:

- The left subtree of a node contains only nodes with keys **less than or equal to** the node's key.
- The right subtree of a node contains only nodes with keys **greater than or equal to** the node's key.
- Both the left and right subtrees must also be binary search trees.

### Example 1:

![image](./image/mode-tree.jpeg)

```
Input: root = [1,null,2,2]
Output: [2]
```

### Example 2:

```
Input: root = [0]
Output: [0]
```

### Constraints:

- The number of nodes in the tree is in the range [1, $10^4$].
- -$10^5$ <= Node.val <= $10^5$

#### Follow up: 

Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search 
- Binary Tree
  
---

# 解題方向

### Solved using Depth-First Search concept

用 **in-order traversal** 的搜尋方式遍歷整棵樹，透過 `local maximum` 與 `global maximum` 兩個變數來控制要輸出的 modes

- 如果下一顆節點與上一顆節點資料相同時就對 `local maximum` 繼續累計
- 如果下一顆節點與上一顆節點資料不同時，就把 `local maximum` 歸零
- `if local maximum == global maximum`, 把現在這顆節點的資料加入 `modes array`
- `if local maximum > global maximum`, 清空舊的 `modes array` 後保存這顆節點的資料