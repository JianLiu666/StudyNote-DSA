[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 108. Convert Sorted Array to Binary Search Tree

Given an integer array `nums` where the elements are sorted in **ascending order**, convert it to a **height-balanced** binary search tree.

A **height-balanced** binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

### Example 1:

![image](./image/btree1.jpeg)

```
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]

Explanation:
 - [0,-10,5,null,-3,null,9] is also accepted:
```

![image](./image/btree2.jpeg)

### Example 2:

![image](./image/btree.jpeg)

```
Input: nums = [1,3]
Output: [3,1]

Explanation:
 - [1,null,3] and [3,1] are both height-balanced BSTs.
```

### Constraints:

- 1 <= nums.length <= $10^4$
- -$10^4$ <= nums[i] <= $10^4$
- `nums` is sorted in a **strictly increasing** order.

### Related Topics

- Array 
- Divide and Conquer 
- Tree 
- Binary Search Tree 
- Binary Tree
  
---

# 解題方向

BST 基本題，就跟 Binary Search 的思路一樣，每次都從中間切一刀，`nums[:mid]` 的部分為 left substree，`nums[mid+1:]` 的部分為 right subtree，重複持續下去直到完成整顆 BST