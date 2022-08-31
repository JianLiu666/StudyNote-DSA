[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/pacific-atlantic-water-flow/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 417. Pacific Atlantic Water Flow

There is an `m x n` rectangular island that borders both the **Pacific Ocean** and **Atlantic Ocean**. The **Pacific Ocean** touches the island's left and top edges, and the **Atlantic Ocean** touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an `m x n` integer matrix `heights` where `heights[r][c]` represents the **height above sea level** of the cell at coordinate `(r, c)`.

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is **less than or equal to** the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a **2D list** of grid coordinates `result` where `result[i]` = [$r_i$, $c_i$] denotes that rain water can flow from cell ($r_i$, $c_i$) to **both** the Pacific and Atlantic oceans.

### Example 1:

![image](./image/waterflow-grid.jpeg)

```
Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
Explanation: 
 - The following cells can flow to the Pacific and Atlantic oceans, as shown below:
 - [0,4]: [0,4] -> Pacific Ocean 
          [0,4] -> Atlantic Ocean
 - [1,3]: [1,3] -> [0,3] -> Pacific Ocean 
          [1,3] -> [1,4] -> Atlantic Ocean
 - [1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean 
          [1,4] -> Atlantic Ocean
 - [2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean 
          [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
 - [3,0]: [3,0] -> Pacific Ocean 
          [3,0] -> [4,0] -> Atlantic Ocean
 - [3,1]: [3,1] -> [3,0] -> Pacific Ocean 
          [3,1] -> [4,1] -> Atlantic Ocean
 - [4,0]: [4,0] -> Pacific Ocean 
          [4,0] -> Atlantic Ocean
 - Note that there are other possible paths for these cells to flow to the Pacific and Atlantic oceans.
```

### Example 2:

```
Input: heights = [[1]]
Output: [[0,0]]

Explanation:
 - The water can flow from the only cell to the Pacific and Atlantic oceans.
```

### Constraints:

- `m == heights.length`
- `n == heights[r].length`
- `1 <= m, n <= 200`
- 0 <= heights[r][c] <= $10^5$

### Related Topics

- Array
- Depth-First Search
- Breadth-First Search
- Matrix
  
---

# 解題方向

### Solved using Depth-First Search

選擇用 **DFS** 而不是 **BFS** 的原因在於，我只求能夠最快速度找到出海口，而不再是否為最短路徑的出海口

有鑒於我們知道四周出海口的位置，所以我們可以依序從每個出海口往陸地中心追溯源頭，直到全部都走訪完畢或再也找不到更高的位置為止

剩下的細節就直接看程式碼吧