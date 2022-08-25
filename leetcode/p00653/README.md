[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 653. Two Sum IV - Input is a BST

Given the `root` of a Binary Search Tree and a target number `k`, return `true` if there exist two elements in the BST such that their sum is equal to the given target.

### Example 1:

![image](./image/sum_tree_1.jpeg)

```
Input: root = [5,3,6,2,4,null,7], k = 9
Output: true
```

### Example 2:

![image](./image/sum_tree_2.jpeg)

```
Input: root = [5,3,6,2,4,null,7], k = 28
Output: false
```

### Constraints:

- The number of nodes in the tree is in the range [1, $10^4$].
- -$10^4$ <= Node.val <= $10^4$
- `root` is guaranteed to be a **valid** binary search tree.
- -$10^5$ <= k <= $10^5$

### Related Topics

- Hash Table
- Two Pointers
- Tree
- Depth-First Search
- Breadth-First Search
- Binary Search Tree
- Binary Tree
  
---

# 解題方向

跟原始的 [Two Sum](./../p00001/README.md) 一模一樣，只是換成以 `tree` 結構來表示而已，只要以 in-order traversal 還原回來後就是一顆 ascending order 的陣列

但這題不用這麼麻煩，直接用一個 `hash table` 把訪問過的節點資料記錄下來，一路比對直到 `hash table` 內存在 `memo[k-current.val]` 的資料就表示找到答案了

