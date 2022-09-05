[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/find-duplicate-subtrees/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 652. Find Duplicate Subtrees

Given the `root` of a binary tree, return all **duplicate subtrees**.

For each kind of duplicate subtrees, you only need to return the root node of any **one** of them.

Two trees are **duplicate** if they have the **same structure** with the **same node values**.

### Example 1:

![image](./image/e1.jpeg)

```
Input: root = [1,2,3,4,null,2,4,null,null,4]
Output: [[2,4],[4]]
```

### Example 2:

![image](./image/e2.jpeg)

```
Input: root = [2,1,1]
Output: [[1]]
```

### Example 3:

![image](./image/e33.jpeg)

```
Input: root = [2,2,2,3,null,3,null]
Output: [[2,3],[3]]
```

### Constraints:

- The number of the nodes in the tree will be in the range [1, 10^4]
- `-200 <= Node.val <= 200`

### Related Topics

- Hash Table
- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Depth-First Search concept

因為相同的兩個 subtrees 很可能是不對稱的，如圖1所示，所以我們只能把尋訪過的紀錄用一個 `Hash Table` 保存起來

- 按照 postorder 的思路遍歷整棵樹，在回到 subtree root 的時候在比較當前的 key 是否有相同的存在過
- hash key 的做法: `{root},{left subtree},{right subtree}`