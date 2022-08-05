[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-depth-of-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 104. Maximum Depth of Binary Tree

Given the `root` of a binary tree, return its maximum depth.

A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

### Example 1:

![image](./image/tmp-tree.jpeg)

```
Input: root = [3,9,20,null,null,15,7]
Output: 3
```

### Example 2:

```
Input: root = [1,null,2]
Output: 2
```

### Constraints:

- The number of nodes in the tree is in the range [0, $10^4$].
- `-100 <= Node.val <= 100`

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Breadth-First Search concept

BFS 暖身題，`Queue` 解決

### Solved using Recursion concept

**Top-down solution**

核心程式碼片段

```go
func helper(node *TreeNode, depth int, maximum *int) {
    if node == nil {
        return
    }

    if depth > *maximum {
        *maximum = depth
    }

    helper(node.Left, depth+1, maximum)
    helper(node.Right, depth+1, maximum)
}
```

**Bottom-up solution**

```go
func helper(node *TreeNode, depth int) int {
    if node == nil {
        return depth
    }

    return max(helper(node.Left, depth+1), helper(node.Right, depth+1))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```