[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 116. Populating Next Right Pointers in Each Node

You are given a **perfect binary tree** where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to `NULL`.

Initially, all next pointers are set to `NULL`.

### Example 1:

![image](./image/116_sample.png)

```
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]

Explanation:
 - Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B.
 - The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
```

### Example 2:

```
Input: root = []
Output: []
```

### Constraints:

- The number of nodes in the tree is in the range [0, $2^{12}$ - 1].
- `-1000 <= Node.val <= 1000`
 

#### Follow-up:

- You may only use constant extra space.
- The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.

### Related Topics

- Linked List
- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using recursion concept

題目有說 `recursion` 的 implicit stack space 可以忽略不計，所以直上吧：

- `left subtree` 的 `left node` 與 `right node`
- `left subtree` 的 `right node` 與 `right subtree` 的 `left node`
- `right subtree` 的 `left node` 與 `right node`

### Solved using Breadth-First Search concept

每 `round` 都先從 `Queue` 裡面拿出第一顆 `node` 當作最左邊的 `node`，接著依序從 `Queue` 取出下一顆節點當作 `right`：

- `left.next = right`
- `left = right`

別忘了把 `left` 與 `right` 的 `child nodes` 加進 `Queue` 裡面作為下一 `round` 使用