[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/longest-univalue-path/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 687. Longest Univalue Path

Given the `root` of a binary tree, return the length of the longest path, where each node in the path has the same value. This path may or may not pass through the root.

**The length of the path** between two nodes is represented by the number of edges between them.

### Example 1:

![image](./image/ex1.jpeg)

```
Input: root = [5,4,5,1,1,null,5]
Output: 2

Explanation:
 - The shown image shows that the longest path of the same value (i.e. 5).
```

### Example 2:

![image](./image/ex2.jpeg)

```
Input: root = [1,4,5,4,4,null,5]
Output: 2

Explanation:
 - The shown image shows that the longest path of the same value (i.e. 4).
```

### Constraints:

- The number of nodes in the tree is in the range [0, $10^4$].
- `-1000 <= Node.val <= 1000`
- The depth of the tree will not exceed `1000`.

### Related Topics

- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Depth-First Search concept

一開始沒有搞懂題目，Longest Univalue Path 的意思是找出相同 value 的**最長**路徑，並非所有相同 value 的總和

換句話說就是從 `Root` 出發，左右子樹只能選中一條路走到最底，更詳細的圖文可以參考 LeetCodeCN 的[這篇圖解](https://leetcode.cn/problems/longest-univalue-path/solution/by-muse-77-rbos/)

搞懂題目後就是基本的 `Tree` 與 DFS 暖身題了，直接看程式碼吧