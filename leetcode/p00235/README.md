[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 235. Lowest Common Ancestor of a Binary Search Tree

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**).”

### Example 1:

![image](./image/binarysearchtree_improved.png)

```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6

Explanation:
 - The LCA of nodes 2 and 8 is 6.
```

### Example 2:

![image](./image/binarysearchtree_improved.png)

```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2

Explanation:
 - The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
```

### Example 3:

```
Input: root = [2,1], p = 2, q = 1
Output: 2
```

### Constraints:

- The number of nodes in the tree is in the range [2, $10^5$].
- -$10^9$ <= Node.val <= $10^9$
- All `Node.val` are **unique**.
- `p` != `q`
- `p` and `q` will exist in the BST.

### Related Topics

- Tree
- Depth-First Search
- Binary Search Tree
- Binary Tree
  
---

# 解題方向

### Solved using Recursion concept

基於 BST 的特性我們只要判斷 `p` 與 `q` 是否在同一側，只要不是的話，那 ancestor 一定是當前這顆節點
