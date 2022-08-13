[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/minimum-absolute-difference-in-bst/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 530. Minimum Absolute Difference in BST

Given the `root` of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

### Example 1:

![image](./image/bst1.jpeg)

```
Input: root = [4,2,6,1,3]
Output: 1
```

### Example 2:

![image](./image/bst2.jpeg)

```
Input: root = [1,0,48,null,null,12,49]
Output: 1
```

### Constraints:

- The number of nodes in the tree is in the range [2, $10^4$].
- 0 <= Node.val <= $10^5$

**Note**: This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
- Binary Search Tree
- Binary Tree
  
---

# 解題方向

### Solved using Depth-First Search concept

先對 BST 做一遍 inorder traversal 得到 ascending order array 後，在對 array 遍歷一次比較找到最小值