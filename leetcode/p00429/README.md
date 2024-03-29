[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/n-ary-tree-level-order-traversal/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 429. N-ary Tree Level Order Traversal

Given an n-ary tree, return the level order traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

### Example 1:

![image](./image/narytreeexample.png)

```
Input: root = [1,null,3,2,4,null,5,6]
Output: [[1],[3,2,4],[5,6]]
```

### Example 2:

![image](./image/sample_4_964.png)

```
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
```

### Constraints:

- The height of the n-ary tree is less than or equal to `1000`
- The total number of nodes is between [0, $10^4$]

### Related Topics

- Tree
- Breatdh-First Search
  
---

# 解題方向

BFS 暖身題，跟一般的 BST level-order traversal 做法一樣，差別在改成對 `children array` 處理而已

