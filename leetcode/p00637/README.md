[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/average-of-levels-in-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 637. Average of Levels in Binary Tree

Given the `root` of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within $10^5$ of the actual answer will be accepted.

### Example 1:

![image](./image/avg1-tree.jpeg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]

Explanation:
 - The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
 - Hence return [3, 14.5, 11].
```

### Example 2:

![image](./image/avg2-tree.jpeg)

```
Input: root = [3,9,20,15,7]
Output: [3.00000,14.50000,11.00000]
```

### Constraints:

- The number of nodes in the tree is in the range [1, $10^4$].
- -$2^{31}$ <= Node.val <= $2^{31}$ - 1

### Related Topics

- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Breadth-First Search concept

BFS 暖身題，用 `Queue` 實現，只要用兩個參數 `sum` 跟 `size` 分別紀錄該層總和跟節點個數，就可以計算每層資料的平均值了

### Solved using Depth-First Search concept

DFS 暖身題，我是直接用 recursion 實現，維護兩個共用的 `array` 來紀錄對應層數的總和與節點個數，最後在對 `array` 遍歷一次取平均值即可，詳細流程直接看程式碼吧