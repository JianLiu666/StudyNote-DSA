[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 236. Lowest Common Ancestor of a Binary Tree

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**).”

### Example 1:

![image](./image/binarytree.png)

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3

Explanation:
 - The LCA of nodes 5 and 1 is 3.
```

### Example 2:

![image](./image/binarytree.png)

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5

Explanation:
 - The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
```

### Example 3:

```
Input: root = [1,2], p = 1, q = 2
Output: 1
```

### Constraints:

- The number of nodes in the tree is in the range [2, $10^5$].
- -$10^9$ <= Node.val <= $10^9$
- All `Node.val` are **unique**.
- `p != q`
- `p` and `q` will exist in the tree.

### Related Topics

- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Recursion concept

`p` 與 `q` 的可能情境只有三種，只要符合其中兩種就代表 lowest common ancestor，以 `current` 為例:

1. `current` 本身是 `p` 或 `q`
2. `current.left` 包含 `p` 或 `q`
3. `current.right` 包含 `p` 或 `q`

因此，我們要做的事情就是對每一顆節點都做一次檢查，最後將自己本身與子節點的查找結果返回回去

只有當上述三種情境有兩種符合時，才代表 LCA

### Solved using Depth-First Search concept

沒辦法像 recursion 一樣將比對結果作為回傳值向上傳遞，因此只能用一個 `Hash Table` 來額外維護對子節點的查找結果
