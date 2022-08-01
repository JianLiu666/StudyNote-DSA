[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/01-matrix/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 542. 01 Matrix

Given an `m x n` binary matrix `mat`, return the distance of the nearest `0` for each cell.

The distance between two adjacent cells is `1`.

### Example 1:

![image](./image/01-1-grid.jpeg)

```
Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]
```

### Example 2:

![image](./image/01-2-grid.jpeg)

```
Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]
```

### Constraints:

- `m == mat.length`
- `n == mat[i].length`
- 1 <= m, n <= $10^4$
- 1 <= m * n <= $10^4$
- `mat[i][j]` is either `0` or `1`.
- There is at least one 0 in `mat`.

### Related Topics

- Array
- Dynamic Programming
- Breadth-First Search
- Matrix
  
---

# 解題方向

### Solved using Breadth-First Search concpet

我們的目標是計算出 `matrix` 中的每個位置距離 `0` 的最短距離，因此反過來說，我們可以從所有元素為 `0` 的位置向鄰近周圍出發，直到遍歷所有非 `0` 的位置，即為答案

示意圖

```
+---+---+---+    +---+---+---+    +---+---+---+
|   |   |   |    |   | 1 |   |    | 2 | 1 | 2 |
+---+---+---+    +---+---+---+    +---+---+---+
|   | 0 |   | => | 1 | 0 | 1 | => | 1 | 0 | 1 |
+---+---+---+    +---+---+---+    +---+---+---+
|   |   |   |    |   | 1 |   |    | 2 | 1 | 2 |
+---+---+---+    +---+---+---+    +---+---+---+
```

### Solved using Dynamic Programming concept

對於每個非 `0` 的 `mat[r][c]` 而言，都是從該位置的四周(上/下/左/右)取最小值累加後得到。因此我們只要對整個 `matrix` 遍歷一遍即可得到答案，如下圖所示：

```
     ------------->
    +---+---+---+---+    +---+---+---+---+
  | | 0 |   |   |   |    | 0 | 1 | 2 | 3 |
  | +---+---+---+---+ => +---+---+---+---+
  v |   |   |   |   |    | 1 | 2 | 3 | 4 |
    +---+---+---+---+    +---+---+---+---+
```

但單純的遍歷一遍只能累加到左/上兩個方向的距離，如果 `0` 在右/下也同時出現時，最短距離就會判斷錯誤，所以我們還必須再從 `matrix` 的尾巴反向遍歷一次，如下圖所示：

```
step.1 foreach whole matrix from top-left
     ------------->
    +---+---+---+---+    +---+---+---+---+
  | | 0 |   |   |   |    | 0 | 1 | 2 | 3 |
  | +---+---+---+---+ => +---+---+---+---+
  v |   |   |   | 0 |    | 1 | 2 | 3 | 0 |
    +---+---+---+---+    +---+---+---+---+

step.2 foreach whole matrix again from bottom-right

      <-------------
    +---+---+---+---+    +---+---+---+---+
  ^ | 0 | 1 | 2 | 3 |    | 0 | 1 | 2 | 1 |
  | +---+---+---+---+ => +---+---+---+---+
  | | 1 | 2 | 3 | 0 |    | 1 | 2 | 1 | 0 |
    +---+---+---+---+    +---+---+---+---+
```