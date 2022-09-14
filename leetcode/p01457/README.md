[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 1457. Pseudo-Palindromic Paths in a Binary Tree

Given a binary tree where node values are digits from 1 to 9. A path in the binary tree is said to be **pseudo-palindromic** if at least one permutation of the node values in the path is a palindrome.

Return the number of **pseudo-palindromic** paths going from the root node to leaf nodes.

### Example 1:

![image](./image/palindromic_paths_1.png)

```
Input: root = [2,3,1,3,1,null,1]
Output: 2 

Explanation:
 - The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1]. 
 - Among these paths only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).
```

### Example 2:

![image](./image/palindromic_paths_2.png)

```
Input: root = [2,1,1,1,3,null,null,null,null,null,1]
Output: 1 

Explanation:
 - The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the green path [2,1,1], the path [2,1,3,1], and the path [2,1]. 
 - Among these paths only the green path is pseudo-palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).
```

### Example 3:

```
Input: root = [9]
Output: 1
```

### Constraints:

- The number of nodes in the tree is in the range [1, $10^5$].
- `1 <= Node.val <= 9`

### Related Topics

- Bit Manipulation
- Tree
- Depth-First Search
- Breadth-First Search
- Binary Tree
  
---

# 解題方向

### Solved using Depth-First Search concept

用 preorder traversal 的思路探索整棵樹，判斷是否構成 Palindromic 的方式有兩種，這題的 [LeetCode Solution](https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/solution/) 說明的很好

基本用 `Hash Table` 紀錄走過的路徑來檢查是否符合 Palindromic 就不多說了，有趣的是因為這題的給出的資料範圍在 `[1,9]` 的有限範圍內，所以可以用 **位運算 + 判斷2次方** 的處理技巧來加速判斷是否符合 Palindromic
