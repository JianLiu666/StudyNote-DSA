[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/number-of-islands/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 200. Number of Islands

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

### Example 1:

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

### Example 2:

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

### Constraints:

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

### Related Topics

- Array
- Depth-First Search
- Breadth-First Search
- Union Find
- Matrix
  
---

# 解題方向

### Solved using Breadth-First Search concept

核心概念為每當踏到一塊陸地上時，就先檢查四周所有的格子是否有相連的陸地，並把他們加進 `queue` 裡，直到 `queue` 內所有的陸地都被尋訪過時，表示完整踏遍這座島了。

示意圖

```
尋訪流程: 0 為起始位置, 數字依序為迭代次數

  step.1               step.2               step.3               eventually

  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 
  | X | 1 |   |   |    | X | 1 | 2 |   |    | X | 1 | 2 | 3 |    | X | 1 | 2 | 3 | 
  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 
  | 1 |   |   |   | => | 1 | 2 |   |   | => | 1 | 2 | 3 |   | => | 1 | 2 | 3 | 4 | 
  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 
  |   |   |   |   |    | 2 |   |   |   |    | 2 | 3 |   |   |    | 2 | 3 | 4 | 5 | 
  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 
```
### Solved using Depth-First Search concept

核心概念為每當踏到一塊陸地上時，先決定一個方向持續前進，直到走到邊界後再換下個方向繼續前進，直到踏遍整座島

- e.g., 依照上下左右順序來尋訪島嶼

示意圖

```
尋訪流程: 0 為起始位置, 數字依序為迭代次數

  step.1               step.2               step.3               eventually
  
  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 
  | X | 1 |   |   |    | X | 1 | 2 |   |    | X | 1 | 2 | 3 |    | X | 1 | 2 | 3 | 
  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 
  |   |   |   |   | => |   |   |   |   | => |   |   |   |   | => | 9 | 10| 11| 4 | 
  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 
  |   |   |   |   |    |   |   |   |   |    |   |   |   |   |    | 8 | 7 | 6 | 5 | 
  +---+---+---+---+    +---+---+---+---+    +---+---+---+---+    +---+---+---+---+ 

```