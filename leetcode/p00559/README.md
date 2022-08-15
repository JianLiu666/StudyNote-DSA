[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-depth-of-n-ary-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 559. Maximum Depth of N-ary Tree

Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

### Example 1:

![image](./image/narytreeexample.png)

```
Input: root = [1,null,3,2,4,null,5,6]
Output: 3
```

### Example 2:

![image](./image/sample_4_964.png)

```
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: 5
```

### Constraints:

- The total number of nodes is in the range [0, $10^4$].
- The depth of the n-ary tree is less than or equal to `1000`.

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
  
---

# 解題方向

BFS 暖身題，跟 level-order traversal 的想法一樣，差別在多紀錄 `depth` 而已
