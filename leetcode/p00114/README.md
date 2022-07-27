[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 114. Flatten Binary Tree to Linked List

Given the `root` of a binary tree, flatten the tree into a "linked list":

- The "linked list" should use the same `TreeNode` class where the `right` child pointer points to the next node in the list and the `left` child pointer is always `null`.
- The "linked list" should be in the same order as a **pre-order traversal** of the binary tree.

### Example 1:

![image](./image/flaten.jpeg)

```
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]
```

### Example 2:

```
Input: root = []
Output: []
```

### Example 3:

```
Input: root = [0]
Output: [0]
```

### Constraints:

- The number of nodes in the tree is in the range `[0, 2000]`.
- `-100 <= Node.val <= 100`

#### Follow up:

Can you flatten the tree in-place (with `O(1)` extra space)?

### Related Topics

- Linked List
- Stack
- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Stack & DFS concept

從樹的右邊開始出發，接著依序將節點由右至左，由底至上的依序加入 `Stack`，最後再將 `Stack` 內的節點依序取出重新排序即可。

但這種做法不能滿足延伸的 `O(1)` 要求，解決辦法一樣是從右邊開始出發，差別每一次尋訪到左節點後就要馬上整理一次這個局部的區域，接著在繼續尋訪即可。
