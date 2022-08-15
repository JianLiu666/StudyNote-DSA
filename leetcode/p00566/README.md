[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reshape-the-matrix/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 566. Reshape the Matrix

In MATLAB, there is a handy function called reshape which can `reshape` an `m x n` matrix into a new one with a different size `r x c` keeping its original data.

You are given an `m x n` matrix `mat` and two integers `r` and `c` representing the number of rows and the number of columns of the wanted reshaped matrix.

The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the `reshape` operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

### Example 1:

![image](./image/reshape1-grid.jpeg)

```
Input: mat = [[1,2],[3,4]], r = 1, c = 4
Output: [[1,2,3,4]]
```

### Example 2:

![image](./image/reshape2-grid.jpeg)

```
Input: mat = [[1,2],[3,4]], r = 2, c = 4
Output: [[1,2],[3,4]]
```

### Constraints:

- `m == mat.length`
- `n == mat[i].length`
- `1 <= m, n <= 100`
- `-1000 <= mat[i][j] <= 1000`
- `1 <= r, c <= 300`

### Related Topics

- Array
- Matrix
- Simulation
  
---

# 解題方向

`Array` 暖身題，只要注意新的 `matrix` 大小必須與原本的大小一樣即可

