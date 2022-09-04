[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/special-positions-in-a-binary-matrix/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 1582. Special Positions in a Binary Matrix

Given an `m x n` binary matrix `mat`, return the number of special positions in `mat`.

A position `(i, j)` is called **special** if `mat[i][j] == 1` and all other elements in row `i` and column `j` are `0` (rows and columns are **0-indexed**).

### Example 1:

![image](./image/special1.jpeg)

```
Input: mat = [[1,0,0],[0,0,1],[1,0,0]]
Output: 1

Explanation:
 - (1, 2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.
```

### Example 2:

![image](./image/special-grid.jpeg)

```
Input: mat = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3

Explanation:
 - (0, 0), (1, 1) and (2, 2) are special positions.
```

### Constraints:

- `m == mat.length`
- `n == mat[i].length`
- `1 <= m, n <= 100`
- `mat[i][j]` is either `0` or `1`.

### Related Topics

- Array
- Matrix
  
---

# 解題方向

最直覺的做法就是額外維護兩個 `Arrays` 來個別紀錄 `rows` 與 `cols` 對應位置的 `1` 的出現次數

詳細作法直接參考程式碼吧

另外如果題目繼續 follow up, 希望我們用 `O(1)` 的空間複雜度來解決這個問題時，可以怎麼做？

我們可以借用 `mat[0]` 來存放我們對每一個 `columns` 遍歷過的統計資料，詳細作法一樣直接看程式碼吧