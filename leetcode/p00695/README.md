[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/max-area-of-island/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 695. Max Area of Island

ou are given an `m x n` binary matrix `grid`. An island is a group of `1`'s (representing land) connected **4-directionally** (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The **area** of an island is the number of cells with a value `1` in the island.

Return the maximum **area** of an island in `grid`. If there is no island, return `0`.

### Example 1:

![image](./image/maxarea1-grid.jpeg)

```
Input: grid = [
    [0,0,1,0,0,0,0,1,0,0,0,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,1,1,0,1,0,0,0,0,0,0,0,0],
    [0,1,0,0,1,1,0,0,1,0,1,0,0],
    [0,1,0,0,1,1,0,0,1,1,1,0,0],
    [0,0,0,0,0,0,0,0,0,0,1,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6

Explanation:
 - The answer is not 11, because the island must be connected 4-directionally.
```

### Example 2:

```
Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0
```

### Constraints:

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 50`
- `grid[i][j]` is either `0` or `1`.

### Related Topics

- Array
- Depth-First Search
- Breadth-First Search
- Union Find
- Matrix
  
---

# 解題方向

### Solved using Depth-First Search concept

核心概念為每當踏到一塊陸地上時，先決定一個方向持續前進，直到走到邊界後再換下個方向繼續前進，直到踏遍整座島

記得當每踏遍一座島時，檢查一下這座島是不是最大的島嶼即可