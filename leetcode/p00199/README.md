[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/binary-tree-right-side-view/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 199. Binary Tree Right Side View

Given the `root` of a binary tree, imagine yourself standing on the **right side** of it, return the values of the nodes you can see ordered from top to bottom.

### Example 1:

![image](./image/tree.jpeg)

```
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
```

### Example 2:

```
Input: root = [1,null,3]
Output: [1,3]
```

### Example 3:

```
Input: root = []
Output: []
```

### Constraints:

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Breadth-First Search concept

BFS 暖身題，有趣的地方在 Right side view 看過去的時候有可能右子樹比左子樹短，就會看到下半截的左子樹，所以我們每一層要取的就是 BFS 的最後一個結點

